package image

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/containers/image/docker/reference"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Client http.Client
	Auth   Authenticator
}

func (c *Client) GetConfig(imageName string) (*Config, error) {
	r, err := reference.ParseNormalizedNamed(imageName)
	if err != nil {
		return nil, errors.Wrap(err, "parsing name")
	}

	digest, err := c.digest(r)
	if err != nil {
		return nil, errors.Wrap(err, "getting digest")
	}

	url := fmt.Sprintf("https://%s/v2/%s/blobs/%s", domain(r), reference.Path(r), digest)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")
	err = c.Auth.Authenticate(req, r)
	if err != nil {
		return nil, errors.Wrap(err, "authenticating digest request")
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "performing http req")
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading resp bytes")
	}

	conf := configWrapper{}
	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling json")
	}

	return &conf.Config, nil
}

func domain(r reference.Named) string {
	d := reference.Domain(r)
	if d == "docker.io" {
		return "registry-1.docker.io"
	} else {
		return d
	}
}

func (c *Client) digest(r reference.Named) (string, error) {
	tag := "latest"
	if tagged, ok := r.(reference.Tagged); ok {
		tag = tagged.Tag()
	}

	url := fmt.Sprintf("https://%s/v2/%s/manifests/%s", domain(r), reference.Path(r), tag)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")
	err := c.Auth.Authenticate(req, r)
	if err != nil {
		return "", errors.Wrap(err, "authenticating digest request")
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "performing digest request")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "reading digest response")
	}

	// TODO: make some structs
	digest, err := jsonparser.GetString(body, "config", "digest")
	if err != nil {
		return "", err
	}

	return digest, nil
}

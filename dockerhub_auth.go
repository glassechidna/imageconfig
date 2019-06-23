package image

import (
	"encoding/json"
	"fmt"
	"github.com/containers/image/docker/reference"
	_ "github.com/motemen/go-loghttp/global"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"time"
)

type expiringToken struct {
	expiry time.Time
	value  string
}

type DockerHubAuthenticator struct {
	Client http.Client
	tokens map[string]*expiringToken
}

func NewDockerHubAuthenticator() *DockerHubAuthenticator {
	return &DockerHubAuthenticator{
		tokens: map[string]*expiringToken{},
	}
}

func (a *DockerHubAuthenticator) Authenticate(req *http.Request, name reference.Named) error {
	token, err := a.token(name)
	if err != nil {
		return errors.Wrap(err, "getting auth token")
	}

	req.Header.Add("Authorization", "Bearer "+token)
	return nil
}

func (a *DockerHubAuthenticator) token(name reference.Named) (string, error) {
	cacheKey := reference.Path(name)
	gracePeriod := time.Now().Add(5 * time.Second)
	if token, found := a.tokens[cacheKey]; found && token.expiry.After(gracePeriod) {
		return token.value, nil
	}

	token, err := a.retrieve(name)
	if err != nil {
		return "", err
	}

	a.tokens[cacheKey] = token
	return token.value, nil
}

func (a *DockerHubAuthenticator) retrieve(name reference.Named) (*expiringToken, error) {
	q := url.Values{}
	q.Add("scope", fmt.Sprintf("repository:%s:pull", reference.Path(name)))
	q.Add("service", "registry.docker.io")
	u, _ := url.Parse("https://auth.docker.io/token")
	u.RawQuery = q.Encode()

	resp, err := a.Client.Get(u.String())
	if err != nil {
		return nil, errors.Wrap(err, "getting auth token")
	}

	m := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		return nil, errors.Wrap(err, "decoding response")
	}

	token := m["token"].(string)
	expiresIn := m["expires_in"].(float64)
	expiry := time.Now().Add(time.Second * time.Duration(expiresIn))

	return &expiringToken{
		expiry: expiry,
		value:  token,
	}, nil
}

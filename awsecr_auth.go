package imageconfig

import (
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecr/ecriface"
	"github.com/containers/image/docker/reference"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type AwsEcrAuthenticator struct {
	Api func(region string) ecriface.ECRAPI
}

func (a *AwsEcrAuthenticator) Authenticate(req *http.Request, name reference.Named) error {
	bits := strings.Split(domain(name), ".")
	registryId := bits[0]
	region := bits[3]

	api := a.Api(region)
	resp, err := api.GetAuthorizationToken(&ecr.GetAuthorizationTokenInput{RegistryIds: []*string{&registryId}})

	if err != nil {
		return errors.Wrap(err, "getting auth token")
	}

	token := *resp.AuthorizationData[0].AuthorizationToken
	req.Header.Add("Authorization", "Basic " + token)

	return nil
}

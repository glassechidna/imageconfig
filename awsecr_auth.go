package imageconfig

import (
	"github.com/aws/aws-sdk-go/service/ecr/ecriface"
	"github.com/containers/image/docker/reference"
	"net/http"
)

type AwsEcrAuthenticator struct {
	Api ecriface.ECRAPI
}

func (a *AwsEcrAuthenticator) Authenticate(req *http.Request, name reference.Named) error {
	panic("implement me")
}

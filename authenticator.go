package imageconfig

import (
	"github.com/containers/image/docker/reference"
	"net/http"
)

type Authenticator interface {
	Authenticate(req *http.Request, name reference.Named) error
}

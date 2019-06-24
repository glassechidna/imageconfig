package imageconfig

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecr/ecriface"
	"github.com/containers/image/docker/reference"
	"net/http"
	"regexp"
)

type MultiAuth struct {
	Options []AuthOption
}

type AuthOption struct {
	Authenticator
	Regexp *regexp.Regexp
}

func (a *MultiAuth) Authenticate(req *http.Request, name reference.Named) error {
	for _, opt := range a.Options {
		if opt.Regexp.MatchString(req.Host) {
			return opt.Authenticate(req, name)
		}
	}

	return nil
}

var DefaultAuthenticator Authenticator

func init() {
	DefaultAuthenticator = &MultiAuth{
		Options: []AuthOption{
			{NewDockerHubAuthenticator(), regexp.MustCompile("^registry-1.docker.io$")},
			{&AwsEcrAuthenticator{Api: func(region string) ecriface.ECRAPI {
				config := aws.NewConfig().WithRegion(region)
				sess := session.Must(session.NewSession(config))
				return ecr.New(sess)
			}}, regexp.MustCompile("^\\.amazonaws.com(?:\\.cn)$")},
		},
	}
}

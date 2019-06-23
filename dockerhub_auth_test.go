package image

import (
	"github.com/containers/image/docker/reference"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDockerHubAuthenticator_Authenticate(t *testing.T) {
	a := NewDockerHubAuthenticator()
	c := &http.Client{}

	t.Run("with full url", func(t *testing.T) {
		req, err := http.NewRequest("GET", "https://registry-1.docker.io/v2/jetbrains/teamcity-agent/manifests/2019.1", nil)
		assert.NoError(t, err)

		name, err := reference.ParseNormalizedNamed("registry-1.index.docker.io/jetbrains/teamcity-agent:2019.1")
		assert.NoError(t, err)

		req.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")
		err = a.Authenticate(req, name)
		assert.NoError(t, err)

		resp, err := c.Do(req)
		assert.NoError(t, err)

		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("with unqualified hub image", func(t *testing.T) {
		req, err := http.NewRequest("GET", "https://registry-1.docker.io/v2/jetbrains/teamcity-agent/manifests/2019.1", nil)
		assert.NoError(t, err)

		name, err := reference.ParseNormalizedNamed("jetbrains/teamcity-agent:2019.1")
		assert.NoError(t, err)

		req.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")
		err = a.Authenticate(req, name)
		assert.NoError(t, err)

		resp, err := c.Do(req)
		assert.NoError(t, err)

		assert.Equal(t, 200, resp.StatusCode)
	})
}

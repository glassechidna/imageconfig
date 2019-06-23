package imageconfig

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_GetConfig(t *testing.T) {
	c := &Client{Auth: NewDockerHubAuthenticator()}
	conf, err := c.GetConfig("awsteele/dotnet21:1")
	assert.NoError(t, err)
	assert.Contains(t, conf.Env, "DOTNET_SDK_VERSION=2.1.604")
}

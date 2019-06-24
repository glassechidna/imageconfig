package imageconfig

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_GetConfig(t *testing.T) {
	t.Run("empty client", func(t *testing.T) {
		c := &Client{}
		conf, err := c.GetConfig("awsteele/dotnet21:1")
		assert.NoError(t, err)
		assert.Contains(t, conf.Env, "DOTNET_SDK_VERSION=2.1.604")
	})

	t.Run("zero client", func(t *testing.T) {
		var c *Client
		conf, err := c.GetConfig("awsteele/dotnet21:1")
		assert.NoError(t, err)
		assert.Contains(t, conf.Env, "DOTNET_SDK_VERSION=2.1.604")
	})
}

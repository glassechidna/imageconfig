package imageconfig

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_EnvMap(t *testing.T) {
	c := &Config{
		Env: []string{
			"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/root/.dotnet/tools",
			"DOTNET_SDK_VERSION=2.1.604",
			"ASPNETCORE_URLS=http://+:80",
			"DOTNET_RUNNING_IN_CONTAINER=true",
			"DOTNET_USE_POLLING_FILE_WATCHER=true",
			"NUGET_XMLDOC_MODE=skip",
		},
	}

	m := c.EnvMap()
	assert.Equal(t, map[string]string{
		"PATH":                            "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/root/.dotnet/tools",
		"DOTNET_SDK_VERSION":              "2.1.604",
		"ASPNETCORE_URLS":                 "http://+:80",
		"DOTNET_RUNNING_IN_CONTAINER":     "true",
		"DOTNET_USE_POLLING_FILE_WATCHER": "true",
		"NUGET_XMLDOC_MODE":               "skip",
	}, m)
}

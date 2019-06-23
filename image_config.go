package imageconfig

import "time"

type Config struct {
	Hostname     string                 `json:"Hostname"`
	Domainname   string                 `json:"Domainname"`
	User         string                 `json:"User"`
	AttachStdin  bool                   `json:"AttachStdin"`
	AttachStdout bool                   `json:"AttachStdout"`
	AttachStderr bool                   `json:"AttachStderr"`
	ExposedPorts map[string]interface{} `json:"ExposedPorts"`
	Tty          bool                   `json:"Tty"`
	OpenStdin    bool                   `json:"OpenStdin"`
	StdinOnce    bool                   `json:"StdinOnce"`
	Env          []string               `json:"Env"`
	Cmd          []string               `json:"Cmd"`
	ArgsEscaped  bool                   `json:"ArgsEscaped"`
	Image        string                 `json:"Image"`
	Volumes      map[string]interface{} `json:"Volumes"`
	WorkingDir   string                 `json:"WorkingDir"`
	Entrypoint   interface{}            `json:"Entrypoint"`
	OnBuild      interface{}            `json:"OnBuild"`
	Labels       map[string]string      `json:"Labels"`
}

type configWrapper struct {
	Architecture  string    `json:"architecture"`
	Config        Config    `json:"config"`
	Container     string    `json:"container"`
	Created       time.Time `json:"created"`
	DockerVersion string    `json:"docker_version"`
	History       []struct {
		Created    time.Time `json:"created"`
		CreatedBy  string    `json:"created_by"`
		EmptyLayer bool      `json:"empty_layer,omitempty"`
	} `json:"history"`
	Os string `json:"os"`
}

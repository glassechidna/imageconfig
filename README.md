```go
c := &imageconfig.Client{Auth: NewDockerHubAuthenticator()}
conf, _ := c.GetConfig("awsteele/dotnet21:1")

/*
(*imageconfig.Config)(0xc00068e290)({
 Hostname: (string) "",
 Domainname: (string) "",
 User: (string) "",
 AttachStdin: (bool) false,
 AttachStdout: (bool) false,
 AttachStderr: (bool) false,
 ExposedPorts: (map[string]interface {}) <nil>,
 Tty: (bool) false,
 OpenStdin: (bool) false,
 StdinOnce: (bool) false,
 Env: ([]string) (len=6 cap=6) {
  (string) (len=85) "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/root/.dotnet/tools",
  (string) (len=26) "DOTNET_SDK_VERSION=2.1.604",
  (string) (len=27) "ASPNETCORE_URLS=http://+:80",
  (string) (len=32) "DOTNET_RUNNING_IN_CONTAINER=true",
  (string) (len=36) "DOTNET_USE_POLLING_FILE_WATCHER=true",
  (string) (len=22) "NUGET_XMLDOC_MODE=skip"
 },
 Cmd: ([]string) (len=1 cap=4) {
  (string) (len=4) "bash"
 },
 ArgsEscaped: (bool) true,
 Image: (string) (len=71) "sha256:d299c5fd7a4e7a617ced296d2ca68a8bc989491b66376b6044efeb5dec0f3593",
 Volumes: (map[string]interface {}) <nil>,
 WorkingDir: (string) "",
 Entrypoint: (interface {}) <nil>,
 OnBuild: (interface {}) <nil>,
 Labels: (map[string]string) <nil>
})
 */
```

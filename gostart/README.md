# golang

Start with `nix develop`. This gives you `go` and `gopls`, which works wonders with the Go VSCode extension. The [go tutorials](https://go.dev/doc/tutorial/) are very good.

## Init project

### Make go module:

Go program packages are called modules. We can initialize one using
```bash
go mod init github.com/harvidsen/teach/gostart
```
It is normal to couple module name with source code url. For repo containing multiple modules we need to configure a [workspace](https://go.dev/doc/tutorial/workspaces).

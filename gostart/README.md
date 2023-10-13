# golang

Start with `nix develop`. This gives you `go` and `gopls`, which works wonders with the Go VSCode extension. The [go tutorials](https://go.dev/doc/tutorial/) are very good, and provide more detail on most things mentioned here.

## Make an executable program with go

**Init go module**

Go program packages are called modules. We can initialize one using
```bash
go mod init github.com/harvidsen/teach/gostart
```
It is normal to couple module name with source code url. For repo containing multiple modules we need to configure a [workspace](https://go.dev/doc/tutorial/workspaces).

**Run it**

Both `package main` and `func main()` are standard and required names that make an executable.


**Install dependencies**

```bash
go get github.com/go-resty/resty/v2
go mod tidy # every now and then
```

**Concepts**

- [x] variables and memory things
- [x] types
- [x] pointers
- [x] function
- [x] goroutine
- [x] channel
- [x] context

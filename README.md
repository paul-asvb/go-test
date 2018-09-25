# golang test repo

## Development

Run file
```bash
go run hello-world.go
```
Build to bin
```bash
go build hello-world.go
```
Format file
```bash
gofmt -w hello-world.go
```

Auto recompile on filechange
1. Install goreload
```bash
go get -u github.com/acoshift/goreload
```

2. use goreload to run script
```bash
goreload hello-world.go
```

## Todo
- [x] Hello World
- [x] install packages
- [x] manage packages
- [x] receive GET request
- [x] receive JSON
- [x] parsing in gin: Go-Obj -> JSON
- [x] parsing in gin: JSON -> Go-Obj
- [x] send JSON
- [x] send GET request
- [x] Init SQLite
- [x] use goreload for autorecompile
- [x] Init dep
- [x] Use swagger
- [ ] Document 1st Endpoint
- [ ] Use URL
- [ ] Query tests

# gRPC in Depth

"gRPC in Depth" is a [gRPC](https://grpc.io) demonstration repository containing examples
about gRPC. It is part of the [YouTube series](https://www.youtube.com/playlist?list=PLsdq-3Z1EPT3Cy2JoRJLA-gwfZ5aeLr_u) titled the same.

Please refer to the videos if you are just getting started with gRPC in Golang.

## Setting up the project

### gRPC

```
$ sudo apt install -y protobuf-compiler
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

This installs the plugins and utilities under `$GOPATH/bin` and the
default `$GOPATH` is `$HOME/go`. Ensure that your `$GOPATH/bin` is in
the `$PATH` variable. Here's what I added to my `~/.bashrc` file

```
export $GOPATH=$HOME/go
export $PATH=$PATH:$GOPATH/bin
```

## Setup Environment

Setup the `.env` file and update the settings.

```
$ mv .env.sample .env
```

Go through the `.env` file and update the secrets.

## Starting services

### Info Service

Info service is a really simple service that exposes one method
`WhatIsGitHub` that upon invocation returns what Github is.

```
$ cd infosvc
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/main.proto
$ go run main.go
```

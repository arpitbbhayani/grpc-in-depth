# gRPC in Depth

The repository contains [gRPC](https://grpc.io) demonstration and examples
to get started with gRPC in Golang. It is part of the
[YouTube series](https://www.youtube.com/playlist?list=PLsdq-3Z1EPT3Cy2JoRJLA-gwfZ5aeLr_u)
by [Arpit Bhayani](https://arpitbhayani.me).

_Please refer the videos in the playlist if you are just getting started with gRPC in Golang._

## Setting up the project

### Setting up BashRC file

Ensure that your BashRC (`~/.bashrc`, `~/.zshrc`) files sets `$GOPATH`
to something predictable. I have added the following lines to my `~/.bashrc`,
but feel free to set it to something that you prefer.

```
$ export $GOPATH=$HOME/go
$ export $PATH=$PATH:$GOPATH/bin
```

### gRPC on Ubuntu

Install the tools to compile `proto` file into Golang code.

```
$ sudo apt install -y protobuf-compiler
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

This installs the plugins and utilities under `$GOPATH/bin` and the
default `$GOPATH` is `$HOME/go`.

> Ensure that your `$GOPATH/bin` is in the `$PATH` variable.

## Postman for gRPC

Refer the video, [gRPC Requests | Postman Level Up](https://www.youtube.com/watch?v=gfYGqMb81GQ),
to setup [postman](https://www.postman.com/) to make client calls to gRPC.

## Starting services

### Info Service

Info service is a really simple service that exposes one method `WhatIsGitHub`
that upon invocation returns what Github is.

```
$ cd infosvc
$ protoc \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    proto/main.proto
$ go run main.go
```

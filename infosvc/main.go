package main

import (
	"log"
	"net"

	infosvc_proto "github.com/arpitbbhayani/grpc-in-depth/infosvc/proto"
	infosvc_server "github.com/arpitbbhayani/grpc-in-depth/infosvc/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8132")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	infosvc_proto.RegisterInfoSvcServer(server, infosvc_server.InfoSvcServerV1{})
	log.Println("listening on port: 8132")
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

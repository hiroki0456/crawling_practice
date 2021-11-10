package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"mf.crawling/pb"
)

func main() {
	listen, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("cannot listen port: %s", err)
	}

	log.Println("start crawling server")

	server := grpc.NewServer()
	serivce := initService()

	pb.RegisterMFCrawlingServiceServer(server, serivce)
	server.Serve(listen)
}

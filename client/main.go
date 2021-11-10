package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	"mf.crawling/pb"
)

func main() {
	//sampleなのでwithInsecure
	conn, err := grpc.Dial("127.0.0.1:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewMFCrawlingServiceClient(conn)

	req := &pb.UserRequest{UserId: os.Args[1], Pass: os.Args[2]}
	res, err := client.CrawlingHandler(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to crawling: %s", err)
	}
	log.Printf("result: %s", res)
}

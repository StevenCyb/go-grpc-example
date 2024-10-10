package main

import (
	"context"
	"gogrpc/pkg/client"
	"gogrpc/pkg/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serverAddr = "127.0.0.1:50051"

func main() {
	grpcClient, err := client.New(
		serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	AssertNoErr("failed to create client: %w", err)

	defer grpcClient.Close()

	resp, err := grpcClient.OneReqOneResp(context.Background(), &proto.DemoRequest{Id: 1})
	AssertNoErr("failed to call OneReqOneResp: %w", err)
	AssertEqual(uint32(2), resp.Result)

	log.Printf("OneReqOneResp: %d", resp.Result)
}

func AssertNoErr(format string, err error) {
	if err != nil {
		log.Fatalf(format, err)
	}
}

func AssertEqual[T comparable](expected, actual T) {
	if expected != actual {
		log.Fatalf("expected %v but got %v", expected, actual)
	}
}

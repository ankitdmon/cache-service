package server

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "github.com/dragno99/cache-service/proto"
	"google.golang.org/grpc"
)

type CustomUserClient struct {
	proto.UnimplementedAddUserClientServiceServer
}

func (s *CustomUserClient) GetUserByID(ctx context.Context, req *proto.Key) (*proto.Value, error) {

	val, err := rDB.Get(req.Key).Result()

	if err != nil {
		return &proto.Value{}, err
	}

	return &proto.Value{Value: []byte(val)}, nil
}

func (s *CustomUserClient) SetUser(ctx context.Context, req *proto.KeyVal) (*proto.Empty, error) {

	err := rDB.Set(req.Key, req.Value, 0).Err()

	if err != nil {
		return &proto.Empty{}, err
	}

	return &proto.Empty{}, nil
}

func StartUserClientServer() {

	listner, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()

	proto.RegisterAddUserClientServiceServer(srv, &CustomUserClient{})

	fmt.Println("Custom User-Client server started...")

	if err := srv.Serve(listner); err != nil {
		log.Fatal("failed to serve: %s", err.Error())
	}

}

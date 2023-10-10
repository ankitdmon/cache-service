package server

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "github.com/dragno99/cache-service/proto"
	"github.com/go-redis/redis"
	"google.golang.org/grpc"
)

var rDB = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

type server struct {
	proto.UnimplementedAddServiceServer
}

func (s *server) Get(ctx context.Context, req *proto.Key) (*proto.Value, error) {

	val, err := rDB.Get(req.Key).Result()

	if err != nil {
		return &proto.Value{}, err
	}

	return &proto.Value{Value: []byte(val)}, nil
}

func (s *server) Set(ctx context.Context, req *proto.KeyVal) (*proto.Empty, error) {

	err := rDB.Set(req.Key, req.Value, 0).Err()

	if err != nil {
		return &proto.Empty{}, err
	}

	return &proto.Empty{}, nil
}

func StartServer() {

	listner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()

	proto.RegisterAddServiceServer(srv, &server{})

	fmt.Println("server started...")

	if err := srv.Serve(listner); err != nil {
		log.Fatal("failed to serve: %v", err.Error())
	}

}

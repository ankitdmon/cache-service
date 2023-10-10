package client

import (
	"context"
	"fmt"
	"log"

	proto "github.com/dragno99/cache-service/proto"

	"google.golang.org/grpc"
)

type User struct {
	Name     string `json:"name"`
	Class    string `json:"class"`
	RollNum  int64  `json:"roll_num"`
	Metadata []byte `json:"metadata"`
}

func SetUser(ctx context.Context, client proto.AddUserClientServiceClient, u *User) error {
	req := &proto.KeyVal{
		Key:   fmt.Sprintf("suryansh:%s:%s:%d", u.Name, u.Class, u.RollNum),
		Value: u.Metadata,
	}
	_, err := client.SetUser(ctx, req)
	return err
}

func GetUserByID(ctx context.Context, client proto.AddUserClientServiceClient, name, class string, roll int64) (*User, error) {
	req := &proto.Key{
		Key: fmt.Sprintf("suryansh:%s:%s:%d", name, class, roll),
	}
	res, err := client.GetUserByID(ctx, req)
	if err != nil {
		return nil, err
	}
	return &User{Name: name, Class: class, RollNum: roll, Metadata: res.GetValue()}, nil
}

func TestCustomUserClient() {

	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := proto.NewAddUserClientServiceClient(conn)

	err = SetUser(context.Background(), client, &User{
		Name:     "suryansh gupta",
		Class:    "12th",
		RollNum:  64,
		Metadata: []byte("he is interested in golang"),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	res, err := GetUserByID(context.Background(), client, "suryansh gupta", "12th", 64)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Response received by Custom user client gRPC server:", res.Name, res.Class, res.Name, string(res.Metadata))

}

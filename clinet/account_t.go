package main

import (
	"context"
	"github.com/reation/micro_service_account_service/protoc"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	accountAddress = "192.168.1.104:8040"
)

func main() {

	GetAccount()
	PayOrder()
}

func GetAccount() {
	conn, err := grpc.Dial(accountAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewAccountClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetAccount(ctx, &protoc.GetAccountRequest{UserId: 1})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Printf("user_id: %d", r.GetUserId())
	log.Printf("balance: %f", r.GetBalance())
}

func PayOrder() {
	orderNum := "2023011623552601234562"
	conn, err := grpc.Dial(accountAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewAccountClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.OpAccount(ctx, &protoc.OpAccountRequest{
		UserId:      1,
		PayPassword: "19861029",
		Prices:      3792.44,
		PaymentId:   1,
		OrderNum:    orderNum,
		Type:        1,
	})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
}

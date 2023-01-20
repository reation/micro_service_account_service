// Code generated by goctl. DO NOT EDIT!
// Source: op_account.proto

package account

import (
	"context"

	"github.com/reation/micro_service_account_service/protoc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetAccountRequest  = protoc.GetAccountRequest
	GetAccountResponse = protoc.GetAccountResponse
	OpAccountRequest   = protoc.OpAccountRequest
	OpAccountResponse  = protoc.OpAccountResponse

	Account interface {
		OpAccount(ctx context.Context, in *OpAccountRequest, opts ...grpc.CallOption) (*OpAccountResponse, error)
		GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	}

	defaultAccount struct {
		cli zrpc.Client
	}
)

func NewAccount(cli zrpc.Client) Account {
	return &defaultAccount{
		cli: cli,
	}
}

func (m *defaultAccount) OpAccount(ctx context.Context, in *OpAccountRequest, opts ...grpc.CallOption) (*OpAccountResponse, error) {
	client := protoc.NewAccountClient(m.cli.Conn())
	return client.OpAccount(ctx, in, opts...)
}

func (m *defaultAccount) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	client := protoc.NewAccountClient(m.cli.Conn())
	return client.GetAccount(ctx, in, opts...)
}

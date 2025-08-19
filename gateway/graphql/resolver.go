package graphql

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductClient *grpc.ClientConn
	OrderClient   *grpc.ClientConn
}

func NewResolver() *Resolver {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ProductConn, err := grpc.DialContext(
		ctx,
		"localhost:9090",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		panic("failed to connect to product service: " + err.Error())
	}

	OrderConn, err := grpc.DialContext(
		ctx,
		"localhost:8080",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		panic("failed to connect to order service: " + err.Error())
	}

	return &Resolver{
		ProductClient: ProductConn,
		OrderClient:   OrderConn,
	}
}

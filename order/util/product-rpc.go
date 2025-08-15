package util

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"github.com/siddheshRajendraNimbalkar/bookstore/order/pb"
)

type ProductServer interface {
	ProductGrpcServer() (*grpc.ClientConn, error)
	GetProduct(productID int64) (*pb.Product, error)
}

func ProductGrpcServer() (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		"localhost:9090",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type ProductReq struct {
	id string
}

type ProductRes struct {
	product_id  string
	name        string
	image       string
	description string
	price       float64
	quantity    int64
	star        float64
	created_at  any
}

func GetProduct(conn *grpc.ClientConn, ctx context.Context, productID string) (*pb.ProductResponse, error) {
	product := &pb.GetProductRequest{
		Id: productID,
	}

	res := &pb.ProductResponse{}
	err := conn.Invoke(ctx, "pb.ProductService/GetProduct", product, res)
	if err != nil {
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

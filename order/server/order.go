package server

import (
	"context"
	"time"

	"github.com/siddheshRajendraNimbalkar/bookstore/order/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {

	return &pb.OrderResponse{
		Order: &pb.Order{
			OrderId:   "22222",
			ProductId: req.GetProductId(),
			Quantity:  req.GetQuantity(),
			Status:    "Pending",
			CreatedAt: timestamppb.New(time.Now()),
		},
	}, nil
}

func (s *Server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {

	return &pb.OrderResponse{
		Order: &pb.Order{
			OrderId:   "22222",
			ProductId: "11111",
			Quantity:  2,
			Status:    "Pending",
			CreatedAt: timestamppb.New(time.Now()),
		},
	}, nil
}

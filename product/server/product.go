package server

import (
	"context"
	"time"

	"github.com/siddheshRajendraNimbalkar/bookstore/product/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {

	return &pb.ProductResponse{
		Product: &pb.Product{
			ProductId:   "11111",
			Name:        req.GetName(),
			Image:       req.GetImage(),
			Description: req.GetDescription(),
			Price:       req.GetPrice(),
			Quantity:    req.GetQuantity(),
			Star:        req.GetStar(),
			CreatedAt:   timestamppb.New(time.Now()),
		},
	}, nil
}

func (s *Server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.ProductResponse, error) {

	return &pb.ProductResponse{
		Product: &pb.Product{
			ProductId:   "11111",
			Name:        "Book ABC11",
			Image:       "https://example.com/book.jpg",
			Description: "A great book",
			Price:       199.99000549316406,
			Quantity:    10,
			Star:        4.5,
			CreatedAt:   timestamppb.New(time.Now()),
		},
	}, nil
}

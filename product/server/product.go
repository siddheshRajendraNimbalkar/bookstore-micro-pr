package server

import (
	"context"

	"github.com/google/uuid"
	db "github.com/siddheshRajendraNimbalkar/bookstore/product/db/sqlc"
	"github.com/siddheshRajendraNimbalkar/bookstore/product/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {

	productDB := db.AddProductParams{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       float64(req.GetPrice()),
		Quantity:    int32(req.GetQuantity()),
		Image:       req.GetImage(),
	}

	product, err := s.db.AddProduct(ctx, productDB)
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Product: &pb.Product{
			ProductId:   product.ID.String(),
			Name:        product.Name,
			Image:       product.Image,
			Description: product.Description,
			Price:       float32(product.Price),
			Quantity:    int64(product.Quantity),
			CreatedAt:   timestamppb.New(product.CreatedAt),
		},
	}, nil
}

func (s *Server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.ProductResponse, error) {

	uuidID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid UUID format for Product ID")
	}
	product, err := s.db.GetProduct(ctx, uuidID)
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Product: &pb.Product{
			ProductId:   product.ID.String(),
			Name:        product.Name,
			Image:       product.Image,
			Description: product.Description,
			Price:       float32(product.Price),
			Quantity:    int64(product.Quantity),
			Star:        0,
			CreatedAt:   timestamppb.New(product.CreatedAt),
		},
	}, nil
}

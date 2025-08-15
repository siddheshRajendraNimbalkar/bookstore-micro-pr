package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	db "github.com/siddheshRajendraNimbalkar/bookstore/order/db/sqlc"
	"github.com/siddheshRajendraNimbalkar/bookstore/order/pb"
	"github.com/siddheshRajendraNimbalkar/bookstore/order/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	productUUID, err := uuid.Parse(req.GetProductId())
	if err != nil {
		return nil, err
	}

	order_params := db.AddOrderParams{
		ProductID: productUUID,
		Quantity:  req.GetQuantity(),
	}

	createdOrder, err := s.db.AddOrder(ctx, order_params)
	if err != nil {
		log.Fatal("err ", err)
		return nil, err
	}

	return &pb.OrderResponse{
		Order: &pb.Order{
			OrderId:   createdOrder.ID.String(),
			ProductId: createdOrder.ProductID.String(),
			Quantity:  createdOrder.Quantity,
			Status:    "Pending",
			CreatedAt: timestamppb.New(time.Now()),
		},
	}, nil
}

func (s *Server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	orderId, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}

	order, err := s.db.GetOrder(ctx, orderId)
	if err != nil {
		return nil, err
	}

	product, err := util.GetProduct(s.productServer, ctx, order.ProductID.String())

	fmt.Println("Product: ", product)
	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{
		Order: &pb.Order{
			OrderId:   order.ID.String(),
			ProductId: order.ProductID.String(),
			Quantity:  order.Quantity,
			Status:    order.Status,
			CreatedAt: timestamppb.New(order.CreatedAt),
		},
		Product: &pb.Product{
			ProductId:   product.Product.ProductId,
			Name:        product.Product.Name,
			Image:       product.Product.Image,
			Description: product.Product.Description,
			Price:       product.Product.Price,
			Quantity:    product.Product.Quantity,
			Star:        product.Product.Star,
			CreatedAt:   product.Product.CreatedAt,
		},
	}, nil
}

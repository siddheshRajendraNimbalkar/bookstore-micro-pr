package server

import (
	_ "github.com/lib/pq"
	db "github.com/siddheshRajendraNimbalkar/bookstore/order/db/sqlc"
	"github.com/siddheshRajendraNimbalkar/bookstore/order/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedBookstoreOrderServer
	db            *db.Queries
	productServer *grpc.ClientConn
}

func NewServer(db *db.Queries, productServer *grpc.ClientConn) *Server {
	return &Server{
		db:            db,
		productServer: productServer,
	}
}

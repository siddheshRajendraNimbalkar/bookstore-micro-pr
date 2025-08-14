package server

import (
	db "github.com/siddheshRajendraNimbalkar/bookstore/product/db/sqlc"
	"github.com/siddheshRajendraNimbalkar/bookstore/product/pb"
)

type Server struct {
	pb.UnimplementedProductServiceServer
	db *db.Queries
}

func NewServer(db *db.Queries) *Server {
	return &Server{
		db: db,
	}
}

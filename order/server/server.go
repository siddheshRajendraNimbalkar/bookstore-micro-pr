package server

import (
	_ "github.com/lib/pq"
	db "github.com/siddheshRajendraNimbalkar/bookstore/order/db/sqlc"
	"github.com/siddheshRajendraNimbalkar/bookstore/order/pb"
)

type Server struct {
	pb.UnimplementedBookstoreOrderServer
	db *db.Queries
}

func NewServer(db *db.Queries) *Server {
	return &Server{
		db: db,
	}
}

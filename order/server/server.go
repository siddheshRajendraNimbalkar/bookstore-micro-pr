package server

import (
	_ "github.com/lib/pq"
	db "github.com/siddheshRajendraNimbalkar/bookstore/order/db/sqlc"
)

type Server struct {
	db *db.Queries
}

func NewServer(db *db.Queries) *Server {
	return &Server{
		db: db,
	}
}

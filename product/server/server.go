package server

import db "github.com/siddheshRajendraNimbalkar/bookstore/product/db/sqlc"

type Server struct {
	db *db.Queries
}

func NewServer(db *db.Queries) *Server {
	return &Server{
		db: db,
	}
}

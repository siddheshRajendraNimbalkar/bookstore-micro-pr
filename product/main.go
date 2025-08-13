package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	db "github.com/siddheshRajendraNimbalkar/bookstore/product/db/sqlc"
	"github.com/siddheshRajendraNimbalkar/bookstore/product/server"
)

func main() {
	fmt.Println("product service is running")

	DB_Driver := "postgres"
	DB_Source := "postgresql://root:password@localhost:5431/bookstore-productDB?sslmode=disable"
	conn, err := sql.Open(DB_Driver, DB_Source)

	if err != nil {
		log.Panic("[ERROR_WHILE_CONN_DB]: ", err)
	}

	defer conn.Close()

	dbQueries := db.New(conn)

	server.NewServer(dbQueries)

}

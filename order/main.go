package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	db "github.com/siddheshRajendraNimbalkar/bookstore/order/db/sqlc"
	"github.com/siddheshRajendraNimbalkar/bookstore/order/pb"
	"github.com/siddheshRajendraNimbalkar/bookstore/order/server"
	"github.com/siddheshRajendraNimbalkar/bookstore/order/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("order service is running")

	DB_Driver := "postgres"
	DB_Source := "postgresql://root:password@localhost:5432/bookstore-orderDB?sslmode=disable"
	conn, err := sql.Open(DB_Driver, DB_Source)

	if err != nil {
		log.Panic("[ERROR_WHILE_CONN_DB]: ", err)
	}

	defer conn.Close()

	dbQueries := db.New(conn)

	// Connect to the product gRPC server
	productServer, err := util.ProductGrpcServer()
	if err != nil {
		log.Panic("[ERROR_WHILE_CONN_PRODUCT_SERVER]: ", err)
	}
	defer productServer.Close()

	server := server.NewServer(dbQueries, productServer)

	grpcServer := grpc.NewServer()
	pb.RegisterBookstoreOrderServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic("[ERROR_WHILE_LISTENING]: ", err)
	}

	log.Println("Order service is listening on port ", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Panic("[ERROR_WHILE_SERVING]: ", err)
	}
}

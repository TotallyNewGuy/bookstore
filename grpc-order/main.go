package main

import (
	db "bookstore/db/sqlc"
	"bookstore/gapi"
	order_gen "bookstore/pb_gen/order_gen"
	"bookstore/util"
	"context"
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	store := db.NewStore(conn)

	initialDb(store)

	runGrpcServer(store, config.GRPC_ADDRESS)
}

func runGrpcServer(store *db.Store, grpcAddress string) {
	server := gapi.NewServer(store)

	grpcServer := grpc.NewServer()
	order_gen.RegisterOrderServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}

func initialDb(store *db.Store) {
	argBook := db.CreateBookParams{
		BookName: util.RandomName(),
		Stock:    util.RandomStock(10, 100),
		Price:    util.RandomMoney(10, 50),
	}
	store.CreateBook(context.Background(), argBook)

	argUser := db.CreateUserParams{
		Name:    util.RandomName(),
		Address: util.RandomAddress(),
		Balance: util.RandomMoney(100, 500),
	}
	store.CreateUser(context.Background(), argUser)
}

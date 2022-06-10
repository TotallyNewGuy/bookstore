package gapi

import (
	db "bookstore/db/sqlc"
	order_gen "bookstore/pb_gen/order_gen"
)

type Server struct {
	order_gen.UnimplementedOrderServiceServer
	store *db.Store
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	return server
}

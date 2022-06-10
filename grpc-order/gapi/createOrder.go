package gapi

import (
	db "bookstore/db/sqlc"
	order_gen "bookstore/pb_gen/order_gen"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateOrder(ctx context.Context, req *order_gen.CreateOrderRequest) (*order_gen.CreateOrderResponse, error) {

	argDB := db.CreateOrderParams{
		UserName:   req.UserName,
		BookName:   req.BookName,
		Amount:     req.Amount,
		Address:    req.Address,
		TotalPrice: req.TotalPrice,
	}

	bookId := req.GetBookId()
	userId := req.GetUserId()
	currBalance := req.GetCurrBalance()
	currStock := req.GetCurrStock()

	// orderDB, err := server.store.CreateOrder(ctx, argDB)
	orderDB, err := server.store.PurchaseTx(ctx, argDB, &db.AuxData{
		UserID:      userId,
		BookID:      bookId,
		CurrBalance: currBalance,
		CurrStock:   currStock,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to create order: %s", err)
	}

	orderGrpc := &order_gen.CreateOrderResponse{
		Order: convertOrder(&orderDB),
	}
	return orderGrpc, nil
}

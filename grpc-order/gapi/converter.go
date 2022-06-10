package gapi

import (
	db "bookstore/db/sqlc"
	order_gen "bookstore/pb_gen/order_gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertOrder(order *db.BookOrder) *order_gen.Order {
	return &order_gen.Order{
		OrderId:    order.OrderID,
		UserName:   order.UserName,
		BookName:   order.BookName,
		Amount:     order.Amount,
		Address:    order.Address,
		TotalPrice: order.TotalPrice,
		CreatedAt:  timestamppb.New(order.CreatedAt),
	}
}

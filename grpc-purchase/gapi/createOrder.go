package gapi

import (
	"context"
	"log"
	pb_book "purchase/pb_gen/book_gen"
	pb_order "purchase/pb_gen/order_gen"
	pb_user_info "purchase/pb_gen/user_info_gen"
	"purchase/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Grpc_cretaeOrder(bookRrsp *pb_book.GetBookResponse, userInfoRsp *pb_user_info.GetUserInfoResponse, amount int32) (*pb_order.CreateOrderResponse, error) {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	orderServiceConn, err := grpc.Dial(config.OrderGrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("lalal! Failed to connect: %v\n", err)
		return nil, err
	}
	defer orderServiceConn.Close()

	orderClient := pb_order.NewOrderServiceClient(orderServiceConn)

	book_price := bookRrsp.Book.GetPrice()

	orderRsp, err := orderClient.CreateOrder(context.Background(), &pb_order.CreateOrderRequest{
		BookName:    bookRrsp.Book.BookName,
		UserName:    userInfoRsp.Userinfo.Name,
		Address:     userInfoRsp.Userinfo.Address,
		Amount:      amount,
		TotalPrice:  book_price * float64(amount),
		UserId:      userInfoRsp.GetUserinfo().GetUserId(),
		BookId:      bookRrsp.GetBook().GetBookId(),
		CurrStock:   bookRrsp.GetBook().GetStock(),
		CurrBalance: userInfoRsp.GetUserinfo().GetBalance(),
	})

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
		return &pb_order.CreateOrderResponse{}, err
	}

	return orderRsp, nil
}

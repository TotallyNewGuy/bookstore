package gapi

import (
	"errors"
	"log"

	pb_book "purchase/pb_gen/book_gen"
	pb_order "purchase/pb_gen/order_gen"
	pb_user_info "purchase/pb_gen/user_info_gen"
)

func Grpc_createPurchase(userId int64, bookId int64, amount int32) (*pb_order.Order, error) {
	bookRsp, err := Grpc_getBook(bookId)
	if err != nil {
		log.Fatalf("Error in Grpc_getBook(): %v\n", err)
		return nil, err
	}
	userInfoRsp, err := Grpc_getUserInfo(userId)
	if err != nil {
		log.Fatalf("Error in Grpc_getUserInfo(): %v\n", err)
		return nil, err
	}

	var orderRsp *pb_order.CreateOrderResponse
	if isValid(bookRsp, userInfoRsp, amount) {
		var err error
		orderRsp, err = Grpc_cretaeOrder(bookRsp, userInfoRsp, amount)
		if err != nil {
			log.Fatalf("Error in Grpc_cretaeOrder(): %v\n", err)
			return nil, err
		}
		return orderRsp.Order, nil
	}

	return nil, errors.New("Grpc_createPurchase() is fail")
}

func isValid(bookRrsp *pb_book.GetBookResponse, userInfoRsp *pb_user_info.GetUserInfoResponse, amount int32) bool {
	book_price := bookRrsp.Book.GetPrice()
	book_stock := bookRrsp.Book.GetStock()
	user_balance := userInfoRsp.Userinfo.GetBalance()

	if user_balance < book_price*float64(amount) {
		log.Panicln("You don't have enough money")
		return false
	}

	if amount > book_stock {
		log.Panicln("There are not enough books")
		return false
	}

	return true
}

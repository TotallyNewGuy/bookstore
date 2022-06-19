package gapi

import (
	"context"
	"fmt"
	"log"

	pb_book "purchase/pb_gen/book_gen"
	"purchase/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Grpc_getBook(bookid int64) (*pb_book.GetBookResponse, error) {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	bookServiceConn, err := grpc.Dial(config.BookGrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println("book_grpc_address: ", config.BookGrpcAddress)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer bookServiceConn.Close()
	bookClient := pb_book.NewBookServiceClient(bookServiceConn)

	arg := pb_book.GetBookRequest{
		BookId: bookid,
	}

	log.Println("Grpc_getBook was invoked")
	res, err := bookClient.GetBook(context.Background(), &arg)

	if err != nil {
		log.Fatalf("Could not getBook: %v\n", err)
		return nil, err
	}

	log.Printf("Book: %s\n", res.Book)
	return res, nil
}

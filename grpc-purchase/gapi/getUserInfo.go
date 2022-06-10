package gapi

import (
	"context"
	"log"

	pb_user_info "purchase/pb_gen/user_info_gen"
	"purchase/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Grpc_getUserInfo(userid int64) (*pb_user_info.GetUserInfoResponse, error) {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	userServiceConn, err := grpc.Dial(config.UserGrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer userServiceConn.Close()

	userClient := pb_user_info.NewUserInfoServiceClient(userServiceConn)

	arg := pb_user_info.GetUserInfoRequest{
		UserId: userid,
	}

	log.Println("Grpc_getUserInfo was invoked")
	res, err := userClient.GetUserInfo(context.Background(), &arg)

	if err != nil {
		log.Fatalf("Could not getUserInfo: %v\n", err)
		return nil, err
	}

	log.Printf("User: %s\n", res.Userinfo)
	return res, nil
}

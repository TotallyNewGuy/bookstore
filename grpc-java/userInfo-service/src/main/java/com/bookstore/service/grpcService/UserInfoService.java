package com.bookstore.service.grpcService;

import com.bookstore.dao.User;
import com.bookstore.pb_gen.user_info_gen.PbGetUserInfo;
import com.bookstore.pb_gen.user_info_gen.UserInfoServiceGrpc;
import io.grpc.stub.StreamObserver;

import java.util.logging.Logger;

import static com.bookstore.service.dbService.GetUserInfoService.getUserInfoById;

public class UserInfoService extends UserInfoServiceGrpc.UserInfoServiceImplBase {

    private static final Logger logger = Logger.getLogger(UserInfoService.class.getName());

    @Override
    public void getUserInfo(PbGetUserInfo.getUserInfoRequest request, StreamObserver<PbGetUserInfo.getUserInfoResponse> responseObserver) {
        long userId = request.getUserId();
        logger.info("Get userId: " + userId);

        // 从数据库得到一个user
        User userFromDB = getUserInfoById(userId);
        logger.info("Get user from database: " + userFromDB.toString());

        PbGetUserInfo.User userProtoc = PbGetUserInfo.User.newBuilder()
                .setUserId(userFromDB.getUser_id())
                .setName(userFromDB.getName())
                .setAddress(userFromDB.getAddress())
                .setBalance(userFromDB.getBalance())
                .build();

        PbGetUserInfo.getUserInfoResponse response = PbGetUserInfo.getUserInfoResponse
                .newBuilder()
                .setUserinfo(userProtoc)
                .build();

        responseObserver.onNext(response);
        responseObserver.onCompleted();

        logger.info("getUserInfo() method has been implemented successfully");
    }
}

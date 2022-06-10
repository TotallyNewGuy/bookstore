package com.bookstore;

import com.bookstore.service.grpcService.UserInfoService;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.protobuf.services.ProtoReflectionService;

import java.io.IOException;
import java.util.concurrent.TimeUnit;
import java.util.logging.Logger;

public class UserInfoServer {

    private static final Logger logger = Logger.getLogger(UserInfoServer.class.getName());

    private final int port;
    private final Server server;

    public UserInfoServer(ServerBuilder serverBuilder, int port) {
        this.port = port;
        UserInfoService userInfoService = new UserInfoService();
        this.server = serverBuilder
                .addService(userInfoService)
                .addService(ProtoReflectionService.newInstance())
                .build();
    }

    public UserInfoServer(int port) {
        this(ServerBuilder.forPort(port), port);
    }

    public void start() throws IOException {
        server.start();
        logger.info("server started on port: " + port);

        Runtime.getRuntime().addShutdownHook(new Thread() {
            @Override
            public void run() {
                System.out.println("shut down gRPC server because JVM shuts down");
                try {
                    UserInfoServer.this.stop();
                } catch (InterruptedException e) {
                    e.printStackTrace(System.err);
                }
                System.err.println("server shut down");
            }
        });
    }

    public void stop() throws InterruptedException {
        if (server != null) {
            server.shutdown().awaitTermination(30, TimeUnit.SECONDS);
        }
    }

    private void blockUntilShutdown() throws InterruptedException {
        if (server != null) {
            server.awaitTermination();
        }
    }

    public static void main(String[] args) throws IOException, InterruptedException {
        UserInfoServer userInfoServer = new UserInfoServer(9091);
        userInfoServer.start();
        userInfoServer.blockUntilShutdown();
    }
}

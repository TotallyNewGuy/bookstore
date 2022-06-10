package com.bookstore;

import com.bookstore.service.grpcService.BookService;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.protobuf.services.ProtoReflectionService;

import java.io.IOException;
import java.util.concurrent.TimeUnit;
import java.util.logging.Logger;

public class BookServer {

    private static final Logger logger = Logger.getLogger(BookServer.class.getName());

    private final int port;
    private final Server server;

    public BookServer(ServerBuilder serverBuilder, int port) {
        this.port = port;
        BookService bookService = new BookService();
        this.server = serverBuilder
                .addService(bookService)
                .addService(ProtoReflectionService.newInstance())
                .build();
    }

    public BookServer(int port) {
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
                    BookServer.this.stop();
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
        BookServer bookServer = new BookServer(9090);
        bookServer.start();
        bookServer.blockUntilShutdown();
    }
}

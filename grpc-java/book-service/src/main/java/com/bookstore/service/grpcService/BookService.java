package com.bookstore.service.grpcService;

import com.bookstore.dao.Book;

import com.bookstore.pb_gen.book_gen.BookServiceGrpc;
import com.bookstore.pb_gen.book_gen.PbGetBook;
import io.grpc.stub.StreamObserver;

import java.util.logging.Logger;

import static com.bookstore.service.dbService.GetBookService.getBookById;

public class BookService extends BookServiceGrpc.BookServiceImplBase {

    private static final Logger logger = Logger.getLogger(BookService.class.getName());


    public void getBook(PbGetBook.getBookRequest request, StreamObserver<PbGetBook.getBookResponse> responseObserver) {
        long book_id = request.getBookId();
        logger.info("Get book_Id: " + book_id);

        // 从数据库得到一个book
        Book bookFromDB = getBookById(book_id);
        logger.info("Get book from database: " + bookFromDB.toString());

        PbGetBook.Book userProtoc = PbGetBook.Book.newBuilder()
                .setBookId(bookFromDB.getBook_id())
                .setBookName(bookFromDB.getBook_name())
                .setPrice(bookFromDB.getPrice())
                .setStock(bookFromDB.getStock())
                .build();

        PbGetBook.getBookResponse response = PbGetBook
                .getBookResponse
                .newBuilder()
                .setBook(userProtoc)
                .build();

        responseObserver.onNext(response);
        responseObserver.onCompleted();

        logger.info("getBook() method has been implemented successfully");
    }
}

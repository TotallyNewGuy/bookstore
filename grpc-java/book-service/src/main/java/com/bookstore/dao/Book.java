package com.bookstore.dao;

public class Book {
    private Long book_id;
    private String book_name;
    private Integer stock;
    private float price;

    public Book(Long book_id, String book_name, Integer stock, float price) {
        this.book_id = book_id;
        this.book_name = book_name;
        this.stock = stock;
        this.price = price;
    }

    public Long getBook_id() {
        return book_id;
    }

    public void setBook_id(Long book_id) {
        this.book_id = book_id;
    }

    public String getBook_name() {
        return book_name;
    }

    public void setBook_name(String book_name) {
        this.book_name = book_name;
    }

    public Integer getStock() {
        return stock;
    }

    public void setStock(Integer stock) {
        this.stock = stock;
    }

    public float getPrice() {
        return price;
    }

    public void setPrice(float price) {
        this.price = price;
    }

    @Override
    public String toString() {
        return "Book{" +
                "book_id=" + book_id +
                ", book_name='" + book_name + '\'' +
                ", stock=" + stock +
                ", price=" + price +
                '}';
    }
}

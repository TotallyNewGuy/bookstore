package com.bookstore.dao.mapper;

import com.bookstore.dao.Book;
import org.apache.ibatis.annotations.Param;

public interface BookMapper {
    Book getBookByIdMybatis(@Param("id") Long id);
}

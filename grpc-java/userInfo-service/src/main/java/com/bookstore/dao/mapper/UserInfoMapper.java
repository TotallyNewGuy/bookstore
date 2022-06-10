package com.bookstore.dao.mapper;

import com.bookstore.dao.User;
import org.apache.ibatis.annotations.Param;

public interface UserInfoMapper {
    User getUserByIdMybatis(@Param("id") Long id);
}

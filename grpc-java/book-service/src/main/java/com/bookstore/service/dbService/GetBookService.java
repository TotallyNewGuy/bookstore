package com.bookstore.service.dbService;

import com.bookstore.dao.Book;
import com.bookstore.dao.mapper.BookMapper;
import org.apache.ibatis.session.SqlSession;
import org.apache.ibatis.session.SqlSessionFactory;
import org.apache.ibatis.session.SqlSessionFactoryBuilder;

import java.io.InputStream;
import java.math.BigDecimal;
import java.math.RoundingMode;

public class GetBookService {

    public static Book getBookById(long id) {
        //读取mybatis配置文件
        InputStream inputStream = GetBookService.class
                .getClassLoader()
                .getResourceAsStream("mybatis/config/config.xml");

        //创建builder
        SqlSessionFactoryBuilder sqlSessionFactoryBuilder = new SqlSessionFactoryBuilder();

        //创建工厂
        SqlSessionFactory sqlSessionFactory = sqlSessionFactoryBuilder.build(inputStream);
        //创建SqlSession对象
        SqlSession sqlSession = sqlSessionFactory.openSession();

        BookMapper mapper = sqlSession.getMapper(BookMapper.class);
        Book res = mapper.getBookByIdMybatis(id);

//        User res = sqlSession.selectOne("getUserByIdMybatis", id);

        roundBalance(res);

        sqlSession.commit();
        sqlSession.close();

        return res;
    }

    public static void roundBalance(Book book) {
        float dou = book.getPrice();
        BigDecimal bigDecimal = new BigDecimal(dou).setScale(2, RoundingMode.HALF_UP);
        float newDouble = bigDecimal.floatValue();

        book.setPrice(newDouble);
    }
}

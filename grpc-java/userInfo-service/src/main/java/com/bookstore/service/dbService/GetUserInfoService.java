package com.bookstore.service.dbService;

import com.bookstore.dao.User;
import com.bookstore.dao.mapper.UserInfoMapper;
import org.apache.ibatis.session.SqlSession;
import org.apache.ibatis.session.SqlSessionFactory;
import org.apache.ibatis.session.SqlSessionFactoryBuilder;

import java.io.InputStream;
import java.math.BigDecimal;
import java.math.RoundingMode;

public class GetUserInfoService {
    public static User getUserInfoById(long id){
        //读取mybatis配置文件
        InputStream inputStream = GetUserInfoService.class
                .getClassLoader()
                .getResourceAsStream("mybatis/config/config.xml");

        //创建builder
        SqlSessionFactoryBuilder sqlSessionFactoryBuilder = new SqlSessionFactoryBuilder();

        //创建工厂
        SqlSessionFactory sqlSessionFactory = sqlSessionFactoryBuilder.build(inputStream);
        //创建SqlSession对象
        SqlSession sqlSession = sqlSessionFactory.openSession();

        UserInfoMapper mapper = sqlSession.getMapper(UserInfoMapper.class);
        User res = mapper.getUserByIdMybatis(id);

//        User res = sqlSession.selectOne("getUserByIdMybatis", id);

        roundBalance(res);

        sqlSession.commit();
        sqlSession.close();

        return res;
    }

    public static void roundBalance(User user) {
        double dou = user.getBalance();
        BigDecimal bigDecimal = new BigDecimal(dou).setScale(2, RoundingMode.HALF_UP);
        double newDouble = bigDecimal.doubleValue();

        user.setBalance(newDouble);
    }
}

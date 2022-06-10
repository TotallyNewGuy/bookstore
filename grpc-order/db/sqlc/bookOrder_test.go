package db

import (
	"bookstore/util"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomBook(t *testing.T) Book {
	arg := CreateBookParams{
		BookName: util.RandomName(),
		Stock:    util.RandomStock(0, 100),
		Price:    util.RandomMoney(10, 50),
	}

	fmt.Println(arg.BookName)

	book, err := testQueries.CreateBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book)

	require.Equal(t, arg.BookName, book.BookName)
	require.Equal(t, arg.Stock, book.Stock)
	require.Equal(t, arg.Price, book.Price)

	require.NotZero(t, book.BookID)

	return book
}

func CreateRandomUser(t *testing.T) UserInfo {
	arg := CreateUserParams{
		Name:    util.RandomName(),
		Address: util.RandomAddress(),
		Balance: util.RandomMoney(100, 500),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Address, user.Address)
	require.Equal(t, arg.Balance, user.Balance)

	require.NotZero(t, user.UserID)

	return user
}

func CreateRandomOrder(t *testing.T) BookOrder {
	arg := CreateOrderParams{
		UserName:   util.RandomName(),
		BookName:   "Three Body Problem",
		Amount:     1,
		Address:    util.RandomAddress(),
		TotalPrice: util.RandomMoney(200, 500),
	}

	order, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, arg.UserName, order.UserName)
	require.Equal(t, arg.BookName, order.BookName)
	require.Equal(t, arg.Amount, order.Amount)
	require.Equal(t, arg.Address, order.Address)
	require.Equal(t, arg.TotalPrice, order.TotalPrice)

	require.NotZero(t, order.OrderID)
	require.NotZero(t, order.CreatedAt)

	return order
}

func TestCreateOrder(t *testing.T) {
	CreateRandomOrder(t)
}

func TestGetOrder(t *testing.T) {
	order1 := CreateRandomOrder(t)
	order2, err := testQueries.GetOrder(context.Background(), order1.OrderID)

	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.OrderID, order2.OrderID)
	require.Equal(t, order1.Address, order2.Address)
	require.Equal(t, order1.Amount, order2.Amount)
	require.Equal(t, order1.UserName, order2.UserName)
	require.Equal(t, order1.BookName, order2.BookName)
	require.Equal(t, order1.TotalPrice, order2.TotalPrice)
	require.WithinDuration(t, order1.CreatedAt, order2.CreatedAt, time.Second)
}

func TestUpdateOrder(t *testing.T) {
	order1 := CreateRandomOrder(t)

	arg := UpdateOrderParams{
		OrderID:    order1.OrderID,
		UserName:   util.RandomName(),
		BookName:   "Four Body Problem",
		Amount:     2,
		Address:    util.RandomAddress(),
		TotalPrice: util.RandomMoney(100, 500),
	}

	order2, err := testQueries.UpdateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.OrderID, order2.OrderID)
	require.Equal(t, arg.Address, order2.Address)
	require.Equal(t, arg.Amount, order2.Amount)
	require.Equal(t, arg.UserName, order2.UserName)
	require.Equal(t, arg.BookName, order2.BookName)
	require.Equal(t, arg.TotalPrice, order2.TotalPrice)
	require.WithinDuration(t, order1.CreatedAt, order2.CreatedAt, time.Second)
}

func TestDeleteOrder(t *testing.T) {
	order1 := CreateRandomOrder(t)
	err := testQueries.DeleteOrder(context.Background(), order1.OrderID)
	require.NoError(t, err)

	order2, err := testQueries.GetOrder(context.Background(), order1.OrderID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, order2)
}

func TestListOrders(t *testing.T) {
	arg := CreateOrderParams{
		UserName:   "List username",
		BookName:   "Three Body Problem",
		Amount:     1,
		Address:    util.RandomAddress(),
		TotalPrice: util.RandomMoney(100, 500),
	}

	for i := 0; i < 5; i++ {
		_, err := testQueries.CreateOrder(context.Background(), arg)
		require.NoError(t, err)
	}

	arg2 := ListOrderParams{
		UserName: "List username",
		Limit:    5,
	}

	orderList, err := testQueries.ListOrder(context.Background(), arg2)
	require.NoError(t, err)
	require.Len(t, orderList, 5)

	for _, order := range orderList {
		require.NotEmpty(t, order)
	}

}

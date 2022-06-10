package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type AuxData struct {
	UserID      int64   `json:"user_id"`
	BookID      int64   `json:"book_id"`
	CurrBalance float64 `json:"currBalance"`
	CurrStock   int32   `json:"currStock"`
}

func (store *Store) PurchaseTx(ctx context.Context, arg CreateOrderParams, aux *AuxData) (BookOrder, error) {

	total_price := arg.TotalPrice
	amount := arg.Amount
	userId := aux.UserID
	bookId := aux.BookID
	currBalance := aux.CurrBalance
	currStock := aux.CurrStock

	var orderDB BookOrder

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		orderDB, err = q.CreateOrder(ctx, arg)
		if err != nil {
			return err
		}

		_, err = q.UpdateUserBalance(ctx, UpdateUserBalanceParams{
			UserID:  userId,
			Balance: currBalance - total_price,
		})
		if err != nil {
			return err
		}

		_, err = q.UpdateBookStock(ctx, UpdateBookStockParams{
			BookID: bookId,
			Stock:  currStock - amount,
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		errMsg := fmt.Sprintf("transcation is fail: %v", err)
		return BookOrder{}, errors.New(errMsg)
	}

	return orderDB, err
}

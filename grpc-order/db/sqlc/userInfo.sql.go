// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: userInfo.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO user_info (
  name,
  address,
  balance
) VALUES (
  $1, $2, $3
)
RETURNING user_id, name, address, balance
`

type CreateUserParams struct {
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (UserInfo, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, arg.Name, arg.Address, arg.Balance)
	var i UserInfo
	err := row.Scan(
		&i.UserID,
		&i.Name,
		&i.Address,
		&i.Balance,
	)
	return i, err
}

const updateUserBalance = `-- name: UpdateUserBalance :one
UPDATE user_info
set 
  balance = $2
WHERE user_id = $1
RETURNING user_id, name, address, balance
`

type UpdateUserBalanceParams struct {
	UserID  int64   `json:"user_id"`
	Balance float64 `json:"balance"`
}

func (q *Queries) UpdateUserBalance(ctx context.Context, arg UpdateUserBalanceParams) (UserInfo, error) {
	row := q.queryRow(ctx, q.updateUserBalanceStmt, updateUserBalance, arg.UserID, arg.Balance)
	var i UserInfo
	err := row.Scan(
		&i.UserID,
		&i.Name,
		&i.Address,
		&i.Balance,
	)
	return i, err
}

-- name: CreateOrder :one
INSERT INTO book_order (
  user_name,
  book_name,
  amount,
  address,
  total_price
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetOrder :one
SELECT * FROM book_order
WHERE order_id = $1 LIMIT 1;

-- name: ListOrder :many
SELECT * FROM book_order
WHERE user_name = $1
ORDER BY created_at
LIMIT $2;

-- name: UpdateOrder :one
UPDATE book_order
set 
  user_name = $2,
  book_name = $3,
  amount = $4,
  address = $5,
  total_price = $6
WHERE order_id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM book_order
WHERE order_id = $1;
-- name: UpdateBookStock :one
UPDATE book
set 
  stock = $2
WHERE book_id = $1
RETURNING *;

-- name: CreateBook :one
INSERT INTO book (
  book_name,
  stock,
  price
) VALUES (
  $1, $2, $3
)
RETURNING *;
-- name: UpdateUserBalance :one
UPDATE user_info
set 
  balance = $2
WHERE user_id = $1
RETURNING *;

-- name: CreateUser :one
INSERT INTO user_info (
  name,
  address,
  balance
) VALUES (
  $1, $2, $3
)
RETURNING *;
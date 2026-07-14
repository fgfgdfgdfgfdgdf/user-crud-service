-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 AND active = true;

-- name: ListUsers :many
SELECT * FROM users WHERE active = true ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (name, active, password, email)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteUser :exec
UPDATE users SET active = false WHERE id = $1 AND active = true;

-- name: HardDeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: UpdateUserNameByID :exec
UPDATE users SET name = $1 WHERE id = $2 AND active = true;

-- name: UpdateUserPasswordByID :exec
UPDATE users SET password = $1 WHERE id = $2 AND active = true;

-- name: UpdateUserByID :exec
UPDATE users SET name = $1, email = $2, password = $3, active = $4 WHERE id = $5 AND active = true;
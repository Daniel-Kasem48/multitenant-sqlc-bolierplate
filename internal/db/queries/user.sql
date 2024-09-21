-- name: CreateUser :one
INSERT INTO users (name)
VALUES ($1)
RETURNING id, name;

-- name: GetUser :one
SELECT id, name 
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, name 
FROM users;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (id, first_name, last_name, email, created_at, updated_at)
VALUES (
    ?,  -- UUID should be generated in application code and passed as a parameter
    ?,  -- first_name
    ?,  -- last_name
    ?,  -- email
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?;

-- name: DeleteAllUsers :exec
DELETE FROM users;
-- name: CreateUser :one
INSERT INTO
    users (id, created_at, updated_at, email)
VALUES
    (
        ?, -- UUID should be generated in application code and passed as a parameter
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP,
        ?
    ) RETURNING *;

-- name: GetUserByEmail :one
SELECT
    *
FROM
    users
WHERE
    email = ?;

-- name: DeleteAllUsers :exec
DELETE FROM users;
-- name: GetUserByEmailSQLC :one
SELECT u.* FROM users AS u WHERE email = ? LIMIT 1;

-- name: UpdatePassword :exec
UPDATE users
SET password_hash = $2,
    updated_at = $3
WHERE id = $1;

-- name: CreateUser :exec
INSERT INTO users (
    username, email, password_hash, created_at, updated_at
) VALUES (
             ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
         )

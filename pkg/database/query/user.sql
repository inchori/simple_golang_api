-- name: GetUsers :many
SELECT * FROM user;

-- name: GetUser :one
SELECT * FROM user WHERE id = ?;

-- name: NewUser :execresult
INSERT INTO user(email, name, password) VALUES (?, ?, ?);

-- name: UpdateUser :execresult
UPDATE user SET email = ?, name = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM user WHERE id = ?;
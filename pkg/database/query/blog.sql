-- name: GetBlogs :many
SELECT * FROM blog;

-- name: GetBlog :one
SELECT * FROM blog WHERE id = ?;

-- name: NewBlog :execresult
INSERT INTO blog(title, content, user) VALUES (?, ?, ?);

-- name: UpdateBlog :execresult
UPDATE blog SET title = ?, content = ? WHERE id = ?;

-- name: DeleteBlog :exec
DELETE FROM blog WHERE id = ?;
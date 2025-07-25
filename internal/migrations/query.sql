-- name: CreateUser :one
INSERT INTO users(username, email, password)
VALUES($1, $2, $3)
RETURNING id, username, email, created_at, updated_at;

-- name: GetUser :one
SELECT id, username, email, created_at, updated_at
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, username, email, created_at, updated_at
FROM users
ORDER BY created_at DESC;


-- name: CreatePost :one
INSERT INTO posts(user_id, title, content)
VALUES($1, $2, $3)
RETURNING id, user_id, title, content, created_at, updated_at;

-- name: ListPosts :many
SELECT id, user_id, title, content, created_at, updated_at
FROM posts
ORDER BY created_at DESC;

-- name: GetPost :one
SELECT id, user_id, title, content, created_at, updated_at
FROM posts
WHERE id = $1;

-- name: UpdatePost :one
UPDATE posts
SET title = $2, content = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, user_id, title, content, created_at, updated_at;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;  

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET username = $2, email = $3, password = $4, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, username, email, created_at, updated_at;




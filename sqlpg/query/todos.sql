-- name: GetTodoByID :one
SELECT * FROM todos WHERE id = $1 AND user_id = $2;

-- name: ListTodosByUser :many
SELECT * FROM todos WHERE user_id = $1 ORDER BY created_at DESC;

-- name: CreateTodo :one
INSERT INTO todos (user_id, title)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos
SET title = $2, done = $3, updated_at = NOW()
WHERE id = $1 AND user_id = $4
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1 AND user_id = $2;

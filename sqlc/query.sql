-- name: GetUserByName :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (username, password, email, confirmed, created_at)
values ($1, $2, $3, $4, $5)
RETURNING *;

-- name: CreateSession :one
INSERT INTO sessions (user_id, token, created_at)
values ($1, $2, $3)
RETURNING *;

-- name: GetSession :one
SELECT users.id AS user_id, users.username, users.email, users.root_task_id FROM
sessions JOIN users ON sessions.user_id = users.id
WHERE token = $1 LIMIT 1;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE token = $1;

-- name: CreateTask :one
INSERT INTO tasks (title, description, result,
created_at, created_by, done_at) values
($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks
WHERE id = $1 LIMIT 1;

-- name: CreateTaskDependency :one
INSERT INTO tasks_dependencies 
(task_id, depends_on_id, created_at) values
($1, $2, $3)
RETURNING *;

-- name: GetTaskDependencies :many
SELECT tasks.* FROM tasks_dependencies
JOIN tasks ON tasks_dependencies.depends_on_id = tasks.id
WHERE task_id = $1 LIMIT 100;

-- name: UpdateRootTaskId :exec
UPDATE users
SET root_task_id = $1
WHERE id = $2;

-- name: UpdateTaskTitle :one
UPDATE tasks
SET title = $1
WHERE id = $2
RETURNING *;

-- name: UpdateTaskDescription :one
UPDATE tasks
SET description = $1
WHERE id = $2
RETURNING *;

-- name: GetTaskParentsTree :many
WITH RECURSIVE parents AS (
    SELECT
        t1.task_id,
        t1.depends_on_id
    FROM
        tasks_dependencies t1
    WHERE
        t1.depends_on_id = $1
    UNION
        SELECT
            t.task_id,
            t.depends_on_id
        FROM
            tasks_dependencies t
        INNER JOIN parents ON parents.task_id = t.depends_on_id
) SELECT
    *
FROM
    parents;

-- name: GetTaskDependenciesTree :many
WITH RECURSIVE children AS (
    SELECT
        t1.task_id as parent_id,
        t1.depends_on_id as task_id,
        tasks.title,
        tasks.done_at
    FROM
        tasks_dependencies t1
    INNER JOIN tasks ON tasks.id = t1.depends_on_id
    WHERE
        t1.task_id = $1
        AND tasks.done_at IS NULL
    UNION
        SELECT
            t.task_id as parent_id,
            t.depends_on_id as task_id,
            tasks.title,
            tasks.done_at
        FROM
            tasks_dependencies t
        INNER JOIN tasks ON tasks.id = t.depends_on_id
        INNER JOIN children ON children.task_id = t.task_id
        WHERE
            tasks.done_at IS NULL
) SELECT
    *
FROM
    children;

-- name: GetTaskParents :many
SELECT tasks.* FROM tasks_dependencies
JOIN tasks ON tasks_dependencies.task_id = tasks.id
WHERE depends_on_id = $1 LIMIT 100;

-- name: SearchTask :many
SELECT * FROM tasks
WHERE
    to_tsvector('english', title) @@ plainto_tsquery($1) AND
    created_by = $2
LIMIT 20;

-- name: UnlinkTask :exec
DELETE FROM tasks_dependencies
WHERE task_id = $1 AND depends_on_id = $2;

-- name: SetTaskDone :one
UPDATE tasks
SET done_at = $2
WHERE id = $1
RETURNING *;

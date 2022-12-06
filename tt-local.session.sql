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
        t1.task_id = 6
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
CREATE TABLE IF NOT EXISTS tasks (
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    title text NOT NULL,
    description text NOT NULL,
    result text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    created_by bigint NOT NULL,
    done_at timestamp without time zone,
    FOREIGN KEY (created_by) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS tasks_dependencies (
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    task_id bigint NOT NULL,
    depends_on_id bigint NOT NULL,
    created_at timestamp without time zone NOT NULL,
    FOREIGN KEY (task_id) REFERENCES tasks(id),
    FOREIGN KEY (depends_on_id) REFERENCES tasks(id)
);

CREATE INDEX tasks_dependencies_task_id_idx
ON tasks_dependencies (task_id);

CREATE INDEX tasks_dependencies_depends_on_id_idx
ON tasks_dependencies (depends_on_id);

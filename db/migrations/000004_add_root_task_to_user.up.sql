ALTER TABLE users
ADD COLUMN root_task_id bigint REFERENCES tasks(id)
ON DELETE SET NULL ON UPDATE CASCADE;

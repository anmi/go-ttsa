CREATE INDEX tasks_search_en_idx ON tasks USING GIN (to_tsvector('english', title));

CREATE TABLE IF NOT EXISTS todo_items (
    id TEXT PRIMARY KEY,
    user_id TEXT,
    name VARCHAR(20),
    description TEXT,
    status BOOLEAN
);
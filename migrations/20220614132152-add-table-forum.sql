
-- +migrate Up
CREATE TABLE IF NOT EXISTS Forums (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    tag VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME,
    updated_at DATETIME
);

CREATE INDEX IF NOT EXISTS idx_title ON Forums(title);
-- +migrate Down
DROP TABLE IF EXISTS Forums;

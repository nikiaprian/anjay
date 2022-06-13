
-- +migrate Up
CREATE TABLE IF NOT EXISTS Blogs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    photo VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME,
    updated_at DATETIME
);

CREATE INDEX IF NOT EXISTS idx_title ON Blogs(title);
-- +migrate Down
DROP TABLE IF EXISTS Blogs;

-- +goose Up
CREATE TABLE sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    expires INTEGER NOT NULL,
    checksum TEXT NOT NULL UNIQUE,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

-- +goose Down
DROP TABLE sessions;

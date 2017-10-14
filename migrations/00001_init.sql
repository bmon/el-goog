-- +goose Up
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    root_folder INTEGER REFERENCES folders(id)
);

-- +goose Down
DROP TABLE users;

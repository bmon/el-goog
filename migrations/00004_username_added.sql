-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE users ADD username TEXT;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE users RENAME TO temp_users;

CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    root_folder INTEGER REFERENCES folders(id)
);
INSERT INTO users SELECT id, email, password, root_folder FROM temp_users;
DROP TABLE temp_users;

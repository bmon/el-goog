-- +goose Up
CREATE TABLE folders (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    parent_id INTEGER,
    name TEXT NOT NULL,
    modified INTEGER NOT NULL,
    FOREIGN KEY(parent_id) REFERENCES folders(id)
);

CREATE TABLE files (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    parent_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    size INTEGER NOT NULL,
    checksum TEXT NOT NULL,
    modified INTEGER NOT NULL,
    FOREIGN KEY(parent_id) REFERENCES folders(id)
);

ALTER TABLE users ADD COLUMN root_folder INTEGER REFERENCES folders(id);



-- +goose Down
DROP TABLE folders;
DROP TABLE files;

-- no alter table drop column in sqlite :c
ALTER TABLE users RENAME TO temp_users;
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

INSERT INTO users SELECT id, email, password FROM temp_users;
DROP TABLE temp_users;

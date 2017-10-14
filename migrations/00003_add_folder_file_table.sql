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

-- +goose Down
DROP TABLE folders;
DROP TABLE files;

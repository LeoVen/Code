
CREATE TABLE folders (
    id integer PRIMARY KEY
);

CREATE TABLE items (
    id integer PRIMARY KEY,
    folder_id integer references folders(id)
);

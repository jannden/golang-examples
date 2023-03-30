CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    done BOOLEAN NOT NULL DEFAULT false
);
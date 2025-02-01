CREATE TABLE author (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    bio TEXT
);

CREATE TABLE book (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    publication_date DATE NOT NULL,
    author_id INT NOT NULL,
    FOREIGN KEY (author_id) REFERENCES author(id) ON DELETE CASCADE
);

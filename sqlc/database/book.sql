-- name: createBook :one
INSERT INTO book (title, publication_date, author_id)
VALUES ($1, $2, $3)
RETURNING id;

-- name: getBook :one
SELECT *
FROM book
WHERE id = $1;

-- name: updateBook :exec
UPDATE book
SET title = $1, publication_date = $2, author_id = $3
WHERE id = $4;

-- name: deleteBook :exec
DELETE FROM book
WHERE id = $1;

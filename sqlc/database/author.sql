-- name: createAuthor :one
INSERT INTO author (email, bio)
VALUES ($1, $2)
RETURNING id;

-- name: getAuthor :one
SELECT *
FROM author
WHERE id = $1;

-- name: updateAuthor :exec
UPDATE author
SET email = $1, bio = $2
WHERE id = $3;

-- name: deleteAuthor :exec
DELETE FROM author
WHERE id = $1;

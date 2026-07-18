-- name: AddFavorite :exec
INSERT INTO favorites(
    session_id,
    user_id,
    book_id
)
VALUES ($1, $2, $3);

-- name: GetFavoritesBySession :many
SELECT
    b.*
FROM favorites f
         JOIN books b ON b.id = f.book_id
WHERE f.session_id = $1;

-- name: GetFavoritesByUser :many
SELECT
    b.*
FROM favorites f
         JOIN books b ON b.id = f.book_id
WHERE f.user_id = $1;

-- name: IsFavorite :one
SELECT EXISTS(
    SELECT 1
    FROM favorites
    WHERE book_id = $1
      AND (
        session_id = $2
            OR user_id = $3
        )
);

-- name: RemoveFavorite :exec
DELETE FROM favorites
WHERE book_id = $1
  AND (
    session_id = $2
        OR user_id = $3
    );

-- name: MoveFavoritesToUser :exec
UPDATE favorites
SET user_id = $1,
    session_id = NULL
WHERE session_id = $2;
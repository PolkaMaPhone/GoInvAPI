-- noinspection SqlResolveForFile

-- name: GetItem :one
SELECT *
FROM items
WHERE item_id = $1;
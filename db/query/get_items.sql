-- noinspection SqlResolveForFile

-- name: GetAllItems :many
SELECT *
FROM items
ORDER BY item_id;
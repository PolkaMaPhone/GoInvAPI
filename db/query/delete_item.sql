-- noinspection SqlResolveForFile

-- name: DeleteItem :execresult
DELETE
FROM items
WHERE item_id = $1
RETURNING *;
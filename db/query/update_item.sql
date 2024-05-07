-- noinspection SqlResolveForFile

-- name: UpdateItem :execresult
UPDATE items
SET name        = $2,
    description = $3,
    category_id = $4,
    group_id    = $5,
    location_id = $6,
    is_stored   = $7,
    updatedAt   = NOW()
WHERE item_id = $1
RETURNING item_id;
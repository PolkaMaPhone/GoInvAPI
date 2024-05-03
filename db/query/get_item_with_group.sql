-- noinspection SqlResolveForFile
-- name: GetItemWithGroup :one
SELECT items.*,
       groups.name        AS group_name,
       groups.description AS group_description
FROM items
         LEFT JOIN groups ON items.group_id = groups.group_id
WHERE item_id = $1;
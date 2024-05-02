-- noinspection SqlResolveForFile
-- name: GetItemWithCategory :one
SELECT items.*,
       categories.name        AS category_name,
       categories.description AS category_description
FROM items
         LEFT JOIN categories ON items.category_id = categories.category_id
WHERE item_id = $1;
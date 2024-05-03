-- noinspection SqlResolveForFile
-- name: GetItemWithGroupAndCategory :one
SELECT items.*,
       categories.name        AS category_name,
       categories.description AS category_description,
       groups.name            AS group_name,
       groups.description     AS group_description

FROM items
         LEFT JOIN categories ON items.category_id = categories.category_id
         LEFT JOIN groups ON items.group_id = groups.group_id
WHERE items.item_id = $1;
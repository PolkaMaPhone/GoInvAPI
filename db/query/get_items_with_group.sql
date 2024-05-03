-- noinspection SqlResolveForFile
-- name: GetAllItemsWithGroups :many
SELECT items.*,
       groups.name        AS group_name,
       groups.description AS group_description
FROM items
         LEFT JOIN groups ON items.group_id = groups.group_id
ORDER BY items.item_id;
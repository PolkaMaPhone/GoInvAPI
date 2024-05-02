-- noinspection SqlResolveForFile
-- name: GetAllItemsWithCategories :many
SELECT items.*,
       categories.name        AS category_name,
       categories.description AS category_description
FROM items
         LEFT JOIN categories ON items.category_id = categories.category_id
ORDER BY items.item_id;
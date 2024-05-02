-- noinspection SqlResolveForFile

-- name: GetAllCategories :many
SELECT category_id, name, description, "createdAt", "updatedAt"
FROM categories
ORDER BY category_id;
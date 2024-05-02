-- noinspection SqlResolveForFile

-- name: GetCategory :one
SELECT category_id, name, description, "createdAt", "updatedAt"
FROM categories
WHERE category_id = $1;
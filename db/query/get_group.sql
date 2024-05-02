-- noinspection SqlResolveForFile

-- name: GetGroup :one
SELECT group_id, name, description, "createdAt", "updatedAt"
FROM groups
WHERE group_id = $1;
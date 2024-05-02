-- noinspection SqlResolveForFile

-- name: GetAllGroups :many
SELECT group_id, name, description, "createdAt", "updatedAt"
FROM groups
ORDER BY group_id;
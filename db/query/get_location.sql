-- noinspection SqlResolveForFile

-- name: GetLocation :one
SELECT location_id, tub_id, shelf_id, "createdAt", "updatedAt"
FROM locations
WHERE location_id = $1;
-- noinspection SqlResolveForFile

-- name: GetAllLocations :many
SELECT location_id, tub_id, shelf_id, "createdAt", "updatedAt"
FROM locations
ORDER BY location_id;
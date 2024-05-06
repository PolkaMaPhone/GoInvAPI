-- noinspection SqlResolveForFile

-- name: CreateItem :one
DO
$$
    BEGIN
        INSERT INTO items (name, description, category_id, group_id, location_id, is_stored, createdAt, updatedAt)
        VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
        RETURNING item_id, createdAt, updatedAt;
    EXCEPTION
        WHEN OTHERS THEN
            RAISE NOTICE 'Create operation failed for item: %', $1;
            ROLLBACK;
    END
$$;
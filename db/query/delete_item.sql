-- noinspection SqlResolveForFile

-- name: DeleteItem :exec
DO
$$
    BEGIN
        DELETE
        FROM items
        WHERE item_id = $1
        RETURNING *;
    EXCEPTION
        WHEN OTHERS THEN
            RAISE NOTICE 'Delete operation failed for item_id: %', $1;
            ROLLBACK;
    END
$$;
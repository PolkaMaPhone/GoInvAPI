// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: get_groups.sql

package db

import (
	"context"
)

const getAllGroups = `-- name: GetAllGroups :many

SELECT group_id, name, description, "createdAt", "updatedAt"
FROM groups
ORDER BY group_id
`

// noinspection SqlResolveForFile
func (q *Queries) GetAllGroups(ctx context.Context) ([]Group, error) {
	rows, err := q.db.Query(ctx, getAllGroups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Group
	for rows.Next() {
		var i Group
		if err := rows.Scan(
			&i.GroupID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
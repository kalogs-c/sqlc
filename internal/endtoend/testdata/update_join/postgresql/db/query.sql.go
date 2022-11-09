// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package db

import (
	"context"
)

const updateJoin = `-- name: UpdateJoin :exec
UPDATE  join_table
SET     is_active = $1
FROM    primary_table
WHERE   join_table.id = $2
        AND primary_table.user_id = $3
        AND join_table.primary_table_id = primary_table.id
`

type UpdateJoinParams struct {
	IsActive bool
	ID       int32
	UserID   int32
}

func (q *Queries) UpdateJoin(ctx context.Context, arg UpdateJoinParams) error {
	_, err := q.db.ExecContext(ctx, updateJoin, arg.IsActive, arg.ID, arg.UserID)
	return err
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package querytest

import (
	"context"
)

const barExists = `-- name: BarExists :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            bar
        where
            id = $1
    )
`

func (q *Queries) BarExists(ctx context.Context, id int32) (bool, error) {
	row := q.db.QueryRowContext(ctx, barExists, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

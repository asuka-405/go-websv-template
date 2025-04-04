// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: delete.sql

package libdb

import (
	"context"
)

const deleteIdentity = `-- name: DeleteIdentity :exec
UPDATE identity
SET deleted_at = CURRENT_TIMESTAMP
WHERE username = $1
`

func (q *Queries) DeleteIdentity(ctx context.Context, username string) error {
	_, err := q.db.ExecContext(ctx, deleteIdentity, username)
	return err
}

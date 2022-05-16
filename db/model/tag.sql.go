// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: tag.sql

package model

import (
	"context"
	"time"

	"cs-api/pkg/types"
)

const countListTag = `-- name: CountListTag :one
select count(*)
from tag
where IF(@name is null, 0, name) like IF(@name is null, 0, CONCAT('%', @name, '%'))
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status)
`

func (q *Queries) CountListTag(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.countListTagStmt, countListTag)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createTag = `-- name: CreateTag :exec
INSERT INTO tag (name, status, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateTagParams struct {
	Name      string       `db:"name" json:"name"`
	Status    types.Status `db:"status" json:"status"`
	CreatedBy int64        `db:"created_by" json:"created_by"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
	UpdatedBy int64        `db:"updated_by" json:"updated_by"`
	UpdatedAt time.Time    `db:"updated_at" json:"updated_at"`
}

func (q *Queries) CreateTag(ctx context.Context, arg CreateTagParams) error {
	_, err := q.exec(ctx, q.createTagStmt, createTag,
		arg.Name,
		arg.Status,
		arg.CreatedBy,
		arg.CreatedAt,
		arg.UpdatedBy,
		arg.UpdatedAt,
	)
	return err
}

const deleteTag = `-- name: DeleteTag :exec
DELETE
FROM tag
WHERE id = ?
`

func (q *Queries) DeleteTag(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteTagStmt, deleteTag, id)
	return err
}

const getAllTag = `-- name: GetAllTag :many
select id, name, status, created_by, created_at, updated_by, updated_at from tag
`

func (q *Queries) GetAllTag(ctx context.Context) ([]Tag, error) {
	rows, err := q.query(ctx, q.getAllTagStmt, getAllTag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Status,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.UpdatedBy,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTag = `-- name: GetTag :one
SELECT id, name, status, created_by, created_at, updated_by, updated_at
FROM tag
WHERE id = ? LIMIT 1
`

func (q *Queries) GetTag(ctx context.Context, id int64) (Tag, error) {
	row := q.queryRow(ctx, q.getTagStmt, getTag, id)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Status,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedBy,
		&i.UpdatedAt,
	)
	return i, err
}

const listTag = `-- name: ListTag :many
select id, name, status, created_by, created_at, updated_by, updated_at
from tag
where IF(@name is null, 0, name) like IF(@name is null, 0, CONCAT('%', @name, '%'))
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status) limit ?
offset ?
`

type ListTagParams struct {
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

func (q *Queries) ListTag(ctx context.Context, arg ListTagParams) ([]Tag, error) {
	rows, err := q.query(ctx, q.listTagStmt, listTag, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Status,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.UpdatedBy,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTag = `-- name: UpdateTag :exec
UPDATE tag
SET name       = ?,
    status     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?
`

type UpdateTagParams struct {
	Name      string       `db:"name" json:"name"`
	Status    types.Status `db:"status" json:"status"`
	UpdatedBy int64        `db:"updated_by" json:"updated_by"`
	UpdatedAt time.Time    `db:"updated_at" json:"updated_at"`
	ID        int64        `db:"id" json:"id"`
}

func (q *Queries) UpdateTag(ctx context.Context, arg UpdateTagParams) error {
	_, err := q.exec(ctx, q.updateTagStmt, updateTag,
		arg.Name,
		arg.Status,
		arg.UpdatedBy,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

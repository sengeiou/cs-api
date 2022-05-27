// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: role.sql

package model

import (
	"context"
	"encoding/json"
	"time"
)

const countListRole = `-- name: CountListRole :one
select count(*)
from role
where IF(@name is null, 0, name) like IF(@name is null, 0, CONCAT('%', @name, '%'))
  and id > 1
`

func (q *Queries) CountListRole(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.countListRoleStmt, countListRole)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createRole = `-- name: CreateRole :exec
INSERT INTO role (name, permissions, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateRoleParams struct {
	Name        string          `db:"name" json:"name"`
	Permissions json.RawMessage `db:"permissions" json:"permissions"`
	CreatedBy   int64           `db:"created_by" json:"created_by"`
	CreatedAt   time.Time       `db:"created_at" json:"created_at"`
	UpdatedBy   int64           `db:"updated_by" json:"updated_by"`
	UpdatedAt   time.Time       `db:"updated_at" json:"updated_at"`
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) error {
	_, err := q.exec(ctx, q.createRoleStmt, createRole,
		arg.Name,
		arg.Permissions,
		arg.CreatedBy,
		arg.CreatedAt,
		arg.UpdatedBy,
		arg.UpdatedAt,
	)
	return err
}

const deleteRole = `-- name: DeleteRole :exec
DELETE
FROM role
WHERE id = ?
`

func (q *Queries) DeleteRole(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteRoleStmt, deleteRole, id)
	return err
}

const getAllRoles = `-- name: GetAllRoles :many
select id, name
from role
`

type GetAllRolesRow struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (q *Queries) GetAllRoles(ctx context.Context) ([]GetAllRolesRow, error) {
	rows, err := q.query(ctx, q.getAllRolesStmt, getAllRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllRolesRow{}
	for rows.Next() {
		var i GetAllRolesRow
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const getRole = `-- name: GetRole :one
SELECT id, name, permissions, created_by, created_at, updated_by, updated_at
FROM role
WHERE id = ? LIMIT 1
`

func (q *Queries) GetRole(ctx context.Context, id int64) (Role, error) {
	row := q.queryRow(ctx, q.getRoleStmt, getRole, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Permissions,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedBy,
		&i.UpdatedAt,
	)
	return i, err
}

const listRole = `-- name: ListRole :many
select id, name, permissions, created_by, created_at, updated_by, updated_at
from role
where IF(@name is null, 0, name) like IF(@name is null, 0, CONCAT('%', @name, '%'))
  and id > 1 limit ?
offset ?
`

type ListRoleParams struct {
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

func (q *Queries) ListRole(ctx context.Context, arg ListRoleParams) ([]Role, error) {
	rows, err := q.query(ctx, q.listRoleStmt, listRole, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Role{}
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Permissions,
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

const updateRole = `-- name: UpdateRole :exec
UPDATE role
SET name        = ?,
    permissions = ?,
    updated_by  = ?,
    updated_at  = ?
WHERE id = ?
`

type UpdateRoleParams struct {
	Name        string          `db:"name" json:"name"`
	Permissions json.RawMessage `db:"permissions" json:"permissions"`
	UpdatedBy   int64           `db:"updated_by" json:"updated_by"`
	UpdatedAt   time.Time       `db:"updated_at" json:"updated_at"`
	ID          int64           `db:"id" json:"id"`
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) error {
	_, err := q.exec(ctx, q.updateRoleStmt, updateRole,
		arg.Name,
		arg.Permissions,
		arg.UpdatedBy,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

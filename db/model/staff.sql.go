// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: staff.sql

package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"cs-api/pkg/types"
)

const countListStaff = `-- name: CountListStaff :one
select count(*)
from staff
where IF(@name is null, 0, name) like IF(@name is null, 0, CONCAT('%', @name, '%'))
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status)
  and IF(@servingStatus is null, 0, serving_status) = IF(@servingStatus is null, 0, @servingStatus)
  and staff.id > 1
`

func (q *Queries) CountListStaff(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.countListStaffStmt, countListStaff)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createStaff = `-- name: CreateStaff :exec
INSERT INTO staff (role_id, name, username, password, status, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateStaffParams struct {
	RoleID    int64        `db:"role_id" json:"role_id"`
	Name      string       `db:"name" json:"name"`
	Username  string       `db:"username" json:"username"`
	Password  string       `db:"password" json:"password"`
	Status    types.Status `db:"status" json:"status"`
	CreatedBy int64        `db:"created_by" json:"created_by"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
	UpdatedBy int64        `db:"updated_by" json:"updated_by"`
	UpdatedAt time.Time    `db:"updated_at" json:"updated_at"`
}

func (q *Queries) CreateStaff(ctx context.Context, arg CreateStaffParams) error {
	_, err := q.exec(ctx, q.createStaffStmt, createStaff,
		arg.RoleID,
		arg.Name,
		arg.Username,
		arg.Password,
		arg.Status,
		arg.CreatedBy,
		arg.CreatedAt,
		arg.UpdatedBy,
		arg.UpdatedAt,
	)
	return err
}

const deleteStaff = `-- name: DeleteStaff :exec
DELETE
FROM staff
WHERE id = ?
`

func (q *Queries) DeleteStaff(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteStaffStmt, deleteStaff, id)
	return err
}

const getAllStaffs = `-- name: GetAllStaffs :many
SELECT id, name, serving_status
FROM staff
WHERE id <> 1
`

type GetAllStaffsRow struct {
	ID            int64                    `db:"id" json:"id"`
	Name          string                   `db:"name" json:"name"`
	ServingStatus types.StaffServingStatus `db:"serving_status" json:"serving_status"`
}

func (q *Queries) GetAllStaffs(ctx context.Context) ([]GetAllStaffsRow, error) {
	rows, err := q.query(ctx, q.getAllStaffsStmt, getAllStaffs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllStaffsRow{}
	for rows.Next() {
		var i GetAllStaffsRow
		if err := rows.Scan(&i.ID, &i.Name, &i.ServingStatus); err != nil {
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

const getStaff = `-- name: GetStaff :one
SELECT role_id, name, status, serving_status, avatar
FROM staff
WHERE staff.id = ? LIMIT 1
`

type GetStaffRow struct {
	RoleID        int64                    `db:"role_id" json:"role_id"`
	Name          string                   `db:"name" json:"name"`
	Status        types.Status             `db:"status" json:"status"`
	ServingStatus types.StaffServingStatus `db:"serving_status" json:"serving_status"`
	Avatar        string                   `db:"avatar" json:"avatar"`
}

func (q *Queries) GetStaff(ctx context.Context, id int64) (GetStaffRow, error) {
	row := q.queryRow(ctx, q.getStaffStmt, getStaff, id)
	var i GetStaffRow
	err := row.Scan(
		&i.RoleID,
		&i.Name,
		&i.Status,
		&i.ServingStatus,
		&i.Avatar,
	)
	return i, err
}

const getStaffCountByRoleId = `-- name: GetStaffCountByRoleId :one
SELECT COUNT(*)
FROM staff
WHERE role_id = ?
`

func (q *Queries) GetStaffCountByRoleId(ctx context.Context, roleID int64) (int64, error) {
	row := q.queryRow(ctx, q.getStaffCountByRoleIdStmt, getStaffCountByRoleId, roleID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const listAvailableStaff = `-- name: ListAvailableStaff :many
SELECT id, name
FROM staff
WHERE serving_status = 2
  and id <> ?
`

type ListAvailableStaffRow struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (q *Queries) ListAvailableStaff(ctx context.Context, id int64) ([]ListAvailableStaffRow, error) {
	rows, err := q.query(ctx, q.listAvailableStaffStmt, listAvailableStaff, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAvailableStaffRow{}
	for rows.Next() {
		var i ListAvailableStaffRow
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

const listStaff = `-- name: ListStaff :many
select staff.id, staff.name, staff.username, staff.status, staff.serving_status, role.name AS role_name
from staff
         inner join role on role.id = staff.role_id
where IF(@name is null, 0, staff.name) like IF(@name is null, 0, CONCAT('%', @name, '%'))
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status)
  and IF(@servingStatus is null, 0, serving_status) = IF(@servingStatus is null, 0, @servingStatus)
  and staff.id > 1 limit ?
offset ?
`

type ListStaffParams struct {
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

type ListStaffRow struct {
	ID            int64                    `db:"id" json:"id"`
	Name          string                   `db:"name" json:"name"`
	Username      string                   `db:"username" json:"username"`
	Status        types.Status             `db:"status" json:"status"`
	ServingStatus types.StaffServingStatus `db:"serving_status" json:"serving_status"`
	RoleName      string                   `db:"role_name" json:"role_name"`
}

func (q *Queries) ListStaff(ctx context.Context, arg ListStaffParams) ([]ListStaffRow, error) {
	rows, err := q.query(ctx, q.listStaffStmt, listStaff, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListStaffRow{}
	for rows.Next() {
		var i ListStaffRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Username,
			&i.Status,
			&i.ServingStatus,
			&i.RoleName,
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

const staffLogin = `-- name: StaffLogin :one
SELECT staff.id, staff.role_id, staff.name, staff.username, staff.serving_status, role.permissions
FROM staff
         INNER JOIN role ON role.id = staff.role_id
WHERE username = ?
  and password = ? LIMIT 1
`

type StaffLoginParams struct {
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

type StaffLoginRow struct {
	ID            int64                    `db:"id" json:"id"`
	RoleID        int64                    `db:"role_id" json:"role_id"`
	Name          string                   `db:"name" json:"name"`
	Username      string                   `db:"username" json:"username"`
	ServingStatus types.StaffServingStatus `db:"serving_status" json:"serving_status"`
	Permissions   json.RawMessage          `db:"permissions" json:"permissions"`
}

func (q *Queries) StaffLogin(ctx context.Context, arg StaffLoginParams) (StaffLoginRow, error) {
	row := q.queryRow(ctx, q.staffLoginStmt, staffLogin, arg.Username, arg.Password)
	var i StaffLoginRow
	err := row.Scan(
		&i.ID,
		&i.RoleID,
		&i.Name,
		&i.Username,
		&i.ServingStatus,
		&i.Permissions,
	)
	return i, err
}

const updateStaff = `-- name: UpdateStaff :exec
UPDATE staff
SET role_id    = ?,
    name       = ?,
    status     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?
`

type UpdateStaffParams struct {
	RoleID    int64        `db:"role_id" json:"role_id"`
	Name      string       `db:"name" json:"name"`
	Status    types.Status `db:"status" json:"status"`
	UpdatedBy int64        `db:"updated_by" json:"updated_by"`
	UpdatedAt time.Time    `db:"updated_at" json:"updated_at"`
	ID        int64        `db:"id" json:"id"`
}

func (q *Queries) UpdateStaff(ctx context.Context, arg UpdateStaffParams) error {
	_, err := q.exec(ctx, q.updateStaffStmt, updateStaff,
		arg.RoleID,
		arg.Name,
		arg.Status,
		arg.UpdatedBy,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

const updateStaffAvatar = `-- name: UpdateStaffAvatar :exec
UPDATE staff
SET avatar     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?
`

type UpdateStaffAvatarParams struct {
	Avatar    string    `db:"avatar" json:"avatar"`
	UpdatedBy int64     `db:"updated_by" json:"updated_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	ID        int64     `db:"id" json:"id"`
}

func (q *Queries) UpdateStaffAvatar(ctx context.Context, arg UpdateStaffAvatarParams) error {
	_, err := q.exec(ctx, q.updateStaffAvatarStmt, updateStaffAvatar,
		arg.Avatar,
		arg.UpdatedBy,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

const updateStaffLogin = `-- name: UpdateStaffLogin :exec
UPDATE staff
SET serving_status  = ?,
    last_login_time = ?,
    updated_at      = ?
WHERE id = ?
`

type UpdateStaffLoginParams struct {
	ServingStatus types.StaffServingStatus `db:"serving_status" json:"serving_status"`
	LastLoginTime sql.NullTime             `db:"last_login_time" json:"last_login_time"`
	UpdatedAt     time.Time                `db:"updated_at" json:"updated_at"`
	ID            int64                    `db:"id" json:"id"`
}

func (q *Queries) UpdateStaffLogin(ctx context.Context, arg UpdateStaffLoginParams) error {
	_, err := q.exec(ctx, q.updateStaffLoginStmt, updateStaffLogin,
		arg.ServingStatus,
		arg.LastLoginTime,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

const updateStaffServingStatus = `-- name: UpdateStaffServingStatus :exec
UPDATE staff
SET serving_status = ?,
    updated_by     = ?,
    updated_at     = ?
WHERE id = ?
`

type UpdateStaffServingStatusParams struct {
	ServingStatus types.StaffServingStatus `db:"serving_status" json:"serving_status"`
	UpdatedBy     int64                    `db:"updated_by" json:"updated_by"`
	UpdatedAt     time.Time                `db:"updated_at" json:"updated_at"`
	ID            int64                    `db:"id" json:"id"`
}

func (q *Queries) UpdateStaffServingStatus(ctx context.Context, arg UpdateStaffServingStatusParams) error {
	_, err := q.exec(ctx, q.updateStaffServingStatusStmt, updateStaffServingStatus,
		arg.ServingStatus,
		arg.UpdatedBy,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

const updateStaffWithPassword = `-- name: UpdateStaffWithPassword :exec
UPDATE staff
SET role_id    = ?,
    name       = ?,
    password   = ?,
    status     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?
`

type UpdateStaffWithPasswordParams struct {
	RoleID    int64        `db:"role_id" json:"role_id"`
	Name      string       `db:"name" json:"name"`
	Password  string       `db:"password" json:"password"`
	Status    types.Status `db:"status" json:"status"`
	UpdatedBy int64        `db:"updated_by" json:"updated_by"`
	UpdatedAt time.Time    `db:"updated_at" json:"updated_at"`
	ID        int64        `db:"id" json:"id"`
}

func (q *Queries) UpdateStaffWithPassword(ctx context.Context, arg UpdateStaffWithPasswordParams) error {
	_, err := q.exec(ctx, q.updateStaffWithPasswordStmt, updateStaffWithPassword,
		arg.RoleID,
		arg.Name,
		arg.Password,
		arg.Status,
		arg.UpdatedBy,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

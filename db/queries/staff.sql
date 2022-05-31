-- name: CreateStaff :exec
INSERT INTO staff (role_id, name, username, password, status, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetStaff :one
SELECT role_id, name, status, serving_status, avatar
FROM staff
WHERE staff.id = ? LIMIT 1;

-- name: UpdateStaff :exec
UPDATE staff
SET role_id    = ?,
    name       = ?,
    status     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?;

-- name: UpdateStaffWithPassword :exec
UPDATE staff
SET role_id    = ?,
    name       = ?,
    password   = ?,
    status     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?;

-- name: UpdateStaffServingStatus :exec
UPDATE staff
SET serving_status = ?,
    updated_by     = ?,
    updated_at     = ?
WHERE id = ?;

-- name: UpdateStaffAvatar :exec
UPDATE staff
SET avatar     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?;

-- name: DeleteStaff :exec
DELETE
FROM staff
WHERE id = ?;

-- name: ListStaff :many
select staff.id, staff.name, staff.username, staff.status, staff.serving_status, role.name AS role_name
from staff
         inner join role on role.id = staff.role_id
where IF(@name is null, 0, staff.name) like IF(@name is null, 0, CONCAT('%', @name, '%'))
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status)
  and IF(@servingStatus is null, 0, serving_status) = IF(@servingStatus is null, 0, @servingStatus)
  and staff.id > 1 limit ?
offset ?;

-- name: CountListStaff :one
select count(*)
from staff
where IF(@name is null, 0, name) like IF(@name is null, 0, CONCAT('%', @name, '%'))
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status)
  and IF(@servingStatus is null, 0, serving_status) = IF(@servingStatus is null, 0, @servingStatus)
  and staff.id > 1;

-- name: StaffLogin :one
SELECT staff.id, staff.role_id, staff.name, staff.username, staff.serving_status, role.permissions
FROM staff
         INNER JOIN role ON role.id = staff.role_id
WHERE username = ?
  and password = ? LIMIT 1;

-- name: UpdateStaffLogin :exec
UPDATE staff
SET serving_status  = ?,
    last_login_time = ?,
    updated_at      = ?
WHERE id = ?;

-- name: ListAvailableStaff :many
SELECT id, name
FROM staff
WHERE serving_status = 2
  and id <> ?;

-- name: GetStaffCountByRoleId :one
SELECT COUNT(*)
FROM staff
WHERE role_id = ?;

-- name: GetAllStaffs :many
SELECT id, name, serving_status
FROM staff
WHERE id <> 1;
-- name: CreateRole :exec
INSERT INTO role (name, permissions, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetRole :one
SELECT *
FROM role
WHERE id = ? LIMIT 1;

-- name: UpdateRole :exec
UPDATE role
SET name        = ?,
    permissions = ?,
    updated_by  = ?,
    updated_at  = ?
WHERE id = ?;

-- name: DeleteRole :exec
DELETE
FROM role
WHERE id = ?;

-- name: ListRole :many
select *
from role
where IF(@name is null, 0, name) like IF(@name is null, 0, CONCAT('%', @name, '%'))
  and id > 1 limit ?
offset ?;

-- name: CountListRole :one
select count(*)
from role
where IF(@name is null, 0, name) like IF(@name is null, 0, CONCAT('%', @name, '%'))
  and id > 1;

-- name: GetAllRoles :many
select id, name
from role

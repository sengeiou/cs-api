-- name: CreateTag :exec
INSERT INTO tag (name, status, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetTag :one
SELECT name, status
FROM tag
WHERE id = ? LIMIT 1;

-- name: UpdateTag :exec
UPDATE tag
SET name       = ?,
    status     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?;

-- name: DeleteTag :exec
DELETE
FROM tag
WHERE id = ?;

-- name: ListTag :many
select id, name, status
from tag
where IF(@name is null, 0, name) like IF(@name is null, 0, CONCAT('%', @name, '%'))
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status) limit ?
offset ?;

-- name: CountListTag :one
select count(*)
from tag
where IF(@name is null, 0, name) like IF(@name is null, 0, CONCAT('%', @name, '%'))
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status);

-- name: GetAllTag :many
select * from tag;

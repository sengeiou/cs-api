-- name: CreateRemind :exec
INSERT INTO remind (title, content, status, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetRemind :one
SELECT *
FROM remind
WHERE id = ? LIMIT 1;

-- name: UpdateRemind :exec
UPDATE remind
SET title      = ?,
    content    = ?,
    status     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?;

-- name: DeleteRemind :exec
DELETE
FROM remind
WHERE id = ?;

-- name: ListRemind :many
select *
from remind
where IF(@content is null, 0, content) like IF(@content is null, 0, CONCAT('%', @content, '%'))
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status)
order by updated_at desc limit ?
offset ?;

-- name: CountListRemind :one
select count(*)
from remind
where IF(@content is null, 0, content) like IF(@content is null, 0, CONCAT('%', @content, '%'))
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status);

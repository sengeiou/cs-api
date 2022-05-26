-- name: CreateFastReply :exec
INSERT INTO fast_reply (category_id, title, content, status, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateFastReply :exec
UPDATE fast_reply
SET category_id = ?,
    title       = ?,
    content     = ?,
    status      = ?,
    updated_by  = ?,
    updated_at  = ?
WHERE id = ?;

-- name: DeleteFastReply :exec
DELETE
FROM fast_reply
WHERE id = ?;

-- name: ListFastReply :many
select fast_reply.*, constant.value AS category
from fast_reply
         inner join constant on constant.id = fast_reply.category_id
where IF(@title is null, 0, title) like IF(@title is null, 0, CONCAT('%', @title, '%'))
  and IF(@content is null, 0, content) like IF(@content is null, 0, CONCAT('%', @content, '%'))
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status) limit ?
offset ?;

-- name: CountListFastReply :one
select count(*)
from fast_reply
where IF(@title is null, 0, title) like IF(@title is null, 0, CONCAT('%', @title, '%'))
  and IF(@content is null, 0, content) like IF(@content is null, 0, CONCAT('%', @content, '%'))
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status);

-- name: GetFastReply :one
SELECT *
FROM fast_reply
WHERE id = ? LIMIT 1;

-- name: GetAllAvailableFastReply :many
SELECT fast_reply.*, constant.value AS category
FROM fast_reply
         INNER JOIN constant ON constant.id = fast_reply.category_id
WHERE status = 1;
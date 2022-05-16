-- name: CreateFastMessage :exec
INSERT INTO fast_message (category_id, title, content, status, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateFastMessage :exec
UPDATE fast_message
SET category_id = ?,
    title       = ?,
    content     = ?,
    status      = ?,
    updated_by  = ?,
    updated_at  = ?
WHERE id = ?;

-- name: DeleteFastMessage :exec
DELETE
FROM fast_message
WHERE id = ?;

-- name: ListFastMessage :many
select fast_message.*, constant.value AS category
from fast_message
         inner join constant on constant.id = fast_message.category_id
where IF(@title is null, 0, title) like IF(@title is null, 0, CONCAT('%', @title, '%'))
  and IF(@content is null, 0, content) like IF(@content is null, 0, CONCAT('%', @content, '%'))
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status) limit ?
offset ?;

-- name: CountListFastMessage :one
select count(*)
from fast_message
where IF(@title is null, 0, title) like IF(@title is null, 0, CONCAT('%', @title, '%'))
  and IF(@content is null, 0, content) like IF(@content is null, 0, CONCAT('%', @content, '%'))
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status);

-- name: GetFastMessage :one
SELECT *
FROM fast_message
WHERE id = ? LIMIT 1;

-- name: GetAllAvailableFastMessage :many
SELECT fast_message.*, constant.value AS category
FROM fast_message
         INNER JOIN constant ON constant.id = fast_message.category_id
WHERE status = 1;
-- name: CreateMerchant :exec
INSERT INTO merchant (name, code, `key`, status, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetMerchant :one
SELECT name, code, status
FROM merchant
WHERE id = ? LIMIT 1;

-- name: UpdateMerchant :exec
UPDATE merchant
SET name       = ?,
    code       = ?,
    status     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?;

-- name: DeleteMerchant :exec
DELETE
FROM merchant
WHERE id = ?;

-- name: ListMerchant :many
select id, name, code, `key`, status
from merchant
where IF(@name is null, 0, name) like IF(@name is null, 0, CONCAT('%', @name, '%')) COLLATE utf8mb4_general_ci
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status) limit ?
offset ?;

-- name: CountListMerchant :one
select count(*)
from merchant
where IF(@name is null, 0, name) like IF(@name is null, 0, CONCAT('%', @name, '%')) COLLATE utf8mb4_general_ci
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status);

-- name: ListAvailableMerchant :many
select id, name
from merchant
where status = 1;

-- name: CheckMerchantKey :one
select id from merchant where `key` = ?;

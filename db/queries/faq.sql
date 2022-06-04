-- name: CreateFAQ :exec
INSERT INTO faq (question, answer, status, created_by, created_at, updated_by, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetFAQ :one
SELECT question, answer
FROM faq
WHERE id = ? LIMIT 1;

-- name: UpdateFAQ :exec
UPDATE faq
SET question   = ?,
    answer     = ?,
    status     = ?,
    updated_by = ?,
    updated_at = ?
WHERE id = ?;

-- name: DeleteFAQ :exec
DELETE
FROM faq
WHERE id = ?;

-- name: ListFAQ :many
select id, question, answer
from faq
where IF(@question is null, 0, question) like IF(@question is null, 0, CONCAT('%', @question, '%')) COLLATE utf8mb4_general_ci
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status) limit ?
offset ?;

-- name: CountListFAQ :one
select count(*)
from faq
where IF(@question is null, 0, question) like IF(@question is null, 0, CONCAT('%', @question, '%')) COLLATE utf8mb4_general_ci
  and IF(@status is null, 0, status) = IF(@status is null, 0, @status);

-- name: ListAvailableFAQ :many
select question, answer
from faq
where status = 1

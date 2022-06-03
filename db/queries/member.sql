-- name: GetGuestMember :one
select *
from member
where type = 2
  and device_id = ? LIMIT 1;

-- name: GetNormalMember :one
select *
from member
where type = 1
  and name = ? LIMIT 1;

-- name: CreateMember :execresult
INSERT INTO member (type, name, device_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);

-- name: GetMemberStatus :one
select status
from member
where id = ?;

-- name: UpdateMemberStatus :exec
update member
set status = ?
where id = ?;
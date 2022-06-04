-- name: CreateMessage :exec
INSERT INTO message (room_id, op_type, sender_type, sender_id, sender_name, content_type, content, extra, ts, created_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: ListStaffRoomMessage :many
select *
from message
where room_id = ?
order by ts desc
limit ? offset ?;

-- name: ListMemberRoomMessage :many
select *
from message
where room_id = ?
  and sender_type <> 0
order by ts desc
limit ? offset ?;

-- name: ListMessage :many
select *
from message
where IF(@roomId is null, 0, room_id) = IF(@roomId is null, 0, @roomId)
  and IF(@staffId is null, 0, sender_type) = IF(@staffId is null, 0, 2)
  and IF(@staffId is null, 0, sender_id) = IF(@staffId is null, 0, @staffId)
  and IF(@content is null, 0, content) like
      IF(@content is null, 0, CONCAT('%', @content, '%')) COLLATE utf8mb4_general_ci
    limit ?
offset ?;

-- name: CountListMessage :one
select count(*)
from message
where IF(@roomId is null, 0, room_id) = IF(@roomId is null, 0, @roomId)
  and IF(@staffId is null, 0, sender_type) = IF(@staffId is null, 0, 2)
  and IF(@staffId is null, 0, sender_id) = IF(@staffId is null, 0, @staffId)
  and IF(@content is null, 0, content) like
      IF(@content is null, 0, CONCAT('%', @content, '%')) COLLATE utf8mb4_general_ci;
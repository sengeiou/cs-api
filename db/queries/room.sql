-- name: CreateRoom :execresult
INSERT INTO room (staff_id, member_id, created_at, updated_at)
VALUES (?, ?, ?, ?);

-- 獲取會員並未關閉的房間
-- name: GetMemberAvailableRoom :one
SELECT * FROM room where member_id = ? and status <> 3 LIMIT 1;

-- name: GetRoom :one
SELECT room.*, member.name AS member_name
FROM room
INNER JOIN member ON member.id = room.member_id
WHERE room.id = ? LIMIT 1;

-- name: AcceptRoom :exec
UPDATE room SET staff_id = ?, status = 2 WHERE id = ?;

-- name: CloseRoom :exec
UPDATE room SET tag_id = ?, closed_at = ?, status = 3 WHERE id = ?;

-- name: UpdateRoomScore :exec
UPDATE room SET score = ? WHERE id = ? and status = 2;

-- name: ListRoom :many
select room.id,
       room.status,
       room.created_at,
       room.closed_at,
       COALESCE(staff.name, '') as staff_name,
       member.name              as member_name,
       COALESCE(tag.name, '')   as tag_name
from room
         left join staff on staff.id = room.staff_id
         left join tag on tag.id = room.tag_id
         inner join member on member.id = room.member_id
where IF(@roomId is null, 0, room.id) = IF(@roomId is null, 0, @roomId)
  and IF(@staffId is null, 0, staff_id) = IF(@staffId is null, 0, @staffId)
  and IF(@status is null, 0, room.status) = IF(@status is null, 0, @status) limit ?
offset ?;

-- name: CountListRoom :one
select count(*)
from room
where IF(@roomId is null, 0, room.id) = IF(@roomId is null, 0, @roomId)
  and IF(@staffId is null, 0, staff_id) = IF(@staffId is null, 0, @staffId)
  and IF(@status is null, 0, room.status) = IF(@status is null, 0, @status);

-- name: ListStaffRoom :many
select room.id, room.status, member.name as member_name
from room
         inner join member on member.id = room.member_id
where status = ?
  and IF(@staffId is null, 0, staff_id) = IF(@staffId is null, 0, @staffId) limit ?
offset ?;

-- name: CountListStaffRoom :one
select count(*)
from room
where status = ?
  and IF(@staffId is null, 0, staff_id) = IF(@staffId is null, 0, @staffId);

-- name: GetStaffRoom :many
SELECT id FROM room where staff_id = ? and status <> 3;

-- 計算每日分類諮詢數
-- name: CountClosedRoomByTag :many
select tag_id, COUNT(*) AS Count
from room
where status = 3 and closed_at between ? and ?
group by tag_id;

-- 計算每日訪客數
-- name: CountDailyRoomByMember :one
select COUNT(distinct member_id) AS GuestCount
from room
where created_at between ? and ?
group by member_id;

-- name: UpdateRoomStaff :exec
update room
set staff_id = ?
where id = ?;
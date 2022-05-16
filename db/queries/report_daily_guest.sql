-- name: DeleteReportDailyGuest :exec
delete from report_daily_guest where date = ?;

-- name: CreateReportDailyGuest :exec
INSERT INTO report_daily_guest (`date`, guest_count, created_at)
VALUES (?, ?, ?);

-- name: ListReportDailyGuest :many
select * from report_daily_guest where date between ? and ?;
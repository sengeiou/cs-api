-- name: DeleteReportDailyTag :exec
delete from report_daily_tag where date = ?;

-- name: CreateReportDailyTag :exec
INSERT INTO report_daily_tag (`date`, tag_id, count, created_at)
VALUES (?, ?, ?, ?);

-- name: ListReportDailyTag :many
select * from report_daily_tag where date between ? and ?;
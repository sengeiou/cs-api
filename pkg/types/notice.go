package types

type FilterNoticeParams struct {
	Content *string
	Status  *Status
}

type Notice struct {
	ID      int64    `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	StartAt JSONTime `json:"start_at"`
	EndAt   JSONTime `json:"end_at"`
	Status  Status   `json:"status"`
}

package types

type DiskDriver string

const (
	DiskDriverLocal DiskDriver = "local"
	DiskDriverS3    DiskDriver = "s3"
)

// Status 開啟狀態
type Status int8

const (
	StatusAll Status = iota
	StatusEnabled
	StatusDisabled
)

type Pagination struct {
	Page     int32 `json:"page" form:"page" binding:"required,gte=1"`
	PageSize int32 `json:"page_size" form:"page_size" binding:"required,gte=1"`
	Total    int64 `json:"total" form:"total"`
}

const (
	RequestCount       = "cs_request_count"
	RequestCountHelp   = "Total request count"
	RequestLatency     = "cs_request_latency"
	RequestLatencyHelp = "Total duration of request in microseconds"
)

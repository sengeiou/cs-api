package types

import (
	"fmt"
	"time"
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

type JSONTime struct {
	time.Time
}

func (t *JSONTime) UnmarshalJSON(timeStr []byte) error {
	tmp, err := time.Parse("2006-01-02 15:04:05", string(timeStr[1:len(timeStr)-1]))
	if err != nil {
		return err
	}
	t.Time = tmp.Add(-8 * time.Hour)
	return nil
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

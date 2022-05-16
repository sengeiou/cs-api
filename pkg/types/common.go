package types

type DiskDriver string

const (
	DiskDriverLocal DiskDriver = "local"
	DiskDriverS3    DiskDriver = "s3"
)

// Status 開啟狀態
type Status int8

const (
	StatusEnabled  Status = 1
	StatusDisabled Status = 2
)

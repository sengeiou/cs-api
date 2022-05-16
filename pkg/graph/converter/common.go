package converter

import (
	"cs-api/pkg/types"
	"time"
)

var StatusMapping = map[Status]types.Status{
	StatusEnabled:  types.StatusEnabled,
	StatusDisabled: types.StatusDisabled,
}

var StatusModelMapping = map[types.Status]Status{
	types.StatusEnabled:  StatusEnabled,
	types.StatusDisabled: StatusDisabled,
}

func FormatTime(t time.Time) string {
	return t.Add(8 * time.Hour).Format("2006-01-02 15:04:05")
}

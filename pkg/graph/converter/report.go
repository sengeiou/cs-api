package converter

import (
	"cs-api/pkg"
	"encoding/json"
	"sort"
)

func (resp *ListDailyTagReportResp) FromMap(columns []pkg.DailyTagReportColumn, result map[string]map[string]int32) {
	resp.Columns = make([]*DailyTagReportColumn, 0, len(columns))
	resp.Items = make([]*DailyTagReportItem, 0)
	for _, column := range columns {
		tmp := DailyTagReportColumn{
			Label: column.Label,
			Key:   column.Key,
		}
		resp.Columns = append(resp.Columns, &tmp)
	}

	var keys []string
	for date, _ := range result {
		keys = append(keys, date)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	for _, key := range keys {
		jsonData, _ := json.Marshal(result[key])
		tmp := DailyTagReportItem{
			Date:     key,
			JSONData: string(jsonData),
		}
		resp.Items = append(resp.Items, &tmp)
	}
}

func (resp *ListDailyGuestReportResp) FromMap(result map[string]int32) {
	resp.Items = make([]*DailyGuestReportItem, 0)

	var keys []string
	for date, _ := range result {
		keys = append(keys, date)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	for _, date := range keys {
		tmp := DailyGuestReportItem{
			Date:       date,
			GuestCount: int64(result[date]),
		}
		resp.Items = append(resp.Items, &tmp)
	}
}

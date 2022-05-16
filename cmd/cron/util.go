package cron

import (
	"cs-api/db/model"
	"database/sql"
	"github.com/rs/zerolog/log"
	"time"
)

func GetTimeRange(dateInput string) (*time.Time, *time.Time, *time.Time, error) {
	if dateInput != "" {
		parseResult, err := time.Parse("2006-01-02", dateFlag1)
		if err != nil {
			log.Error().Msgf("parse error: %s", err)
			return nil, nil, nil, err
		}
		date := parseResult
		start := parseResult.Add(-8 * time.Hour)
		end := start.Add(24 * time.Hour).Add(-1 * time.Second)
		return &date, &start, &end, err
	}

	date := time.Now().UTC().Add(8 * time.Hour).Add(-24 * time.Hour)
	tmp := date.Format("2006-01-02")
	parseResult, err := time.Parse("2006-01-02", tmp)
	if err != nil {
		log.Error().Msgf("parse error: %s", err)
		return nil, nil, nil, err
	}
	start := parseResult.Add(-8 * time.Hour)
	end := start.Add(24 * time.Hour).Add(-1 * time.Second)

	return &date, &start, &end, err
}

func NewDBTX(db *sql.DB) model.DBTX {
	return db
}

package cs_config

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg"
	"cs-api/pkg/types"
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

func (s *service) GetCsConfig(ctx context.Context) (config types.CsConfig, err error) {
	constant, err := s.repo.GetCsConfig(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return config, nil
	}
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(constant.Value), &config)
	if err != nil {
		return
	}

	return
}

func (s *service) UpdateCsConfig(ctx context.Context, staffId int64, config types.CsConfig) error {
	jsonStr, _ := json.Marshal(config)
	params := model.UpdateCsConfigParams{
		Value:     string(jsonStr),
		UpdatedBy: staffId,
		UpdatedAt: time.Now().UTC(),
	}

	if err := s.repo.UpdateCsConfig(ctx, params); err != nil {
		return err
	}

	event := pkg.StaffEventInfo{
		Event: pkg.StaffEventUpdateConfig,
		Payload: pkg.StaffEventPayload{
			CsConfig: &config,
		},
	}

	payload, _ := json.Marshal(event)

	if err := s.redis.Publish(ctx, "event:staff", payload); err != nil {
		return err
	}

	return nil
}

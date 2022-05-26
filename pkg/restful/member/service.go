package member

import (
	"context"
	"crypto/md5"
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"time"
)

type service struct {
	repo iface.IRepository
}

func (s *service) GetOrCreateMember(ctx context.Context, name string, deviceId string) (member model.Member, err error) {
	if name == "" {
		return s.getOrCreateMemberByDevice(ctx, deviceId)
	}
	return s.getOrCreateMemberByName(ctx, name, deviceId)
}

func (s *service) getOrCreateMemberByDevice(ctx context.Context, deviceId string) (model.Member, error) {
	member, err := s.repo.GetGuestMember(ctx, deviceId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return model.Member{}, err
	}

	now := time.Now().UTC()
	memberName := fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()+deviceId)))[16:19]

	var memberId int64
	// member 不存在
	if errors.Is(err, sql.ErrNoRows) {
		var err2 error
		result, err2 := s.repo.CreateMember(ctx, model.CreateMemberParams{
			Type:      types.MemberTypeGuest,
			Name:      "Guest-" + memberName,
			DeviceID:  deviceId,
			CreatedAt: now,
			UpdatedAt: now,
		})
		if err2 != nil {
			log.Error().Msgf("create member error: %s", err2)
			return model.Member{}, err2
		}
		memberId, err2 = result.LastInsertId()
		if err2 != nil {
			log.Error().Msgf("get last insert id error: %s", err2)
			return model.Member{}, err2
		}
	} else {
		memberId = member.ID
	}

	result := model.Member{
		ID:        memberId,
		Type:      types.MemberTypeGuest,
		Name:      "Guest-" + memberName,
		DeviceID:  deviceId,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return result, nil
}

func (s *service) getOrCreateMemberByName(ctx context.Context, name string, deviceId string) (model.Member, error) {
	member, err := s.repo.GetNormalMember(ctx, name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return model.Member{}, err
	}

	now := time.Now().UTC()

	var memberId int64
	// member 不存在
	if errors.Is(err, sql.ErrNoRows) {
		result, err := s.repo.CreateMember(ctx, model.CreateMemberParams{
			Type:      types.MemberTypeNormal,
			Name:      name,
			DeviceID:  deviceId,
			CreatedAt: now,
			UpdatedAt: now,
		})
		if err != nil {
			log.Error().Msgf("create member error: %s", err)
			return model.Member{}, err
		}
		memberId, err = result.LastInsertId()
		if err != nil {
			log.Error().Msgf("get last insert id error: %s", err)
			return model.Member{}, err
		}
	} else {
		memberId = member.ID
	}

	result := model.Member{
		ID:        memberId,
		Type:      types.MemberTypeNormal,
		Name:      name,
		DeviceID:  deviceId,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return result, nil
}

func NewService(Repo iface.IRepository) iface.IMemberService {
	return &service{
		repo: Repo,
	}
}

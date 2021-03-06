package room

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"database/sql"
	_ "embed"
	"encoding/json"
	"errors"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"time"
)

type service struct {
	redis     ifaceTool.IRedis
	lua       iface.ILusScript
	memberSvc iface.IMemberService
	repo      iface.IRepository
}

// CreateRoom 如果 name 有傳入則尋找該會員名是否已經存在，不存在則創建。
// 如果 name 沒有傳入，則尋找是否有相同 deviceId 的訪客，不存在則創建。
// 如果該會員的諮詢房尚未結束就延續，以結束就創建新的
func (s *service) CreateRoom(ctx context.Context, deviceId string, name string) (room model.Room, member model.Member, err error) {
	member, err = s.memberSvc.GetOrCreateMember(ctx, name, deviceId)
	if err != nil {
		return
	}

	room, err = s.repo.GetMemberAvailableRoom(ctx, member.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}

	if room.ID != 0 {
		return room, member, nil
	}

	now := time.Now().UTC()
	result, err := s.repo.CreateRoom(ctx, model.CreateRoomParams{
		StaffID:   0,
		MemberID:  member.ID,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	tmpRoom := model.Room{
		ID:        id,
		StaffID:   0,
		MemberID:  member.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return tmpRoom, member, nil
}

func (s *service) AcceptRoom(ctx context.Context, staffId int64, roomId int64) error {
	var (
		err  error
		room model.GetRoomRow
	)

	room, err = s.repo.GetRoom(ctx, roomId)
	if err != nil {
		return err
	}

	if room.Status == types.RoomStatusClosed {
		return errors.New("wrong room status")
	}

	event := pkg.StaffEventInfo{
		Event: pkg.StaffEventAcceptRoom,
		Payload: pkg.StaffEventPayload{
			StaffID: &staffId,
			RoomID:  &roomId,
		},
	}

	payload, _ := json.Marshal(event)

	if err = s.redis.Publish(ctx, "event:staff", payload); err != nil {
		return err
	}

	return s.repo.AcceptRoom(ctx, model.AcceptRoomParams{
		StaffID: staffId,
		ID:      roomId,
	})
}

func (s *service) CloseRoom(ctx context.Context, staffId int64, roomId int64, tagId int64) error {
	var (
		err  error
		room model.GetRoomRow
	)

	room, err = s.repo.GetRoom(ctx, roomId)
	if err != nil {
		return err
	}

	if room.StaffID != staffId {
		return errors.New("operation error")
	}

	if room.Status != types.RoomStatusServing {
		return errors.New("wrong room status")
	}

	event := pkg.StaffEventInfo{
		Event: pkg.StaffEventCloseRoom,
		Payload: pkg.StaffEventPayload{
			StaffID: &room.StaffID,
			RoomID:  &roomId,
		},
	}

	payload, _ := json.Marshal(event)
	if err = s.redis.Publish(ctx, "event:staff", payload); err != nil {
		log.Error().Msgf("Publish to redis error: %s", err)
		return err
	}

	// remove member token
	if err = s.lua.RemoveToken(ctx, "member", room.MemberName); err != nil {
		log.Error().Msgf("clear token error: %s\n", err)
		return err
	}

	return s.repo.CloseRoom(ctx, model.CloseRoomParams{
		TagID:    tagId,
		ClosedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		ID:       roomId,
	})
}

func (s *service) UpdateRoomScore(ctx context.Context, roomId int64, score int32) error {
	return s.repo.UpdateRoomScore(ctx, model.UpdateRoomScoreParams{
		Score: score,
		ID:    roomId,
	})
}

func (s *service) ListRoom(ctx context.Context, params model.ListRoomParams, filterParams types.FilterRoomParams) (result []types.Room, count int64, err error) {
	rooms := make([]model.ListRoomRow, 0)
	err = s.repo.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		var err2 error

		_, err2 = tx.Exec("SET @roomId = ?", filterParams.RoomID)
		if err2 != nil {
			return err2
		}

		_, err2 = tx.Exec("SET @staffId = ?", filterParams.StaffID)
		if err2 != nil {
			return err2
		}

		_, err2 = tx.Exec("SET @status = ?", filterParams.Status)
		if err2 != nil {
			return err2
		}

		rooms, err2 = s.repo.WithTx(tx).ListRoom(ctx, params)
		if err2 != nil {
			return err2
		}

		count, err2 = s.repo.WithTx(tx).CountListRoom(ctx)
		if err2 != nil {
			return err2
		}

		return nil
	})

	result = make([]types.Room, 0, len(rooms))
	for _, room := range rooms {
		tmp := types.Room{
			ID:         room.ID,
			Status:     room.Status,
			CreatedAt:  types.JSONTime{Time: room.CreatedAt},
			ClosedAt:   nil,
			StaffName:  room.StaffName,
			MemberName: room.MemberName,
			TagName:    room.TagName,
		}
		if room.ClosedAt.Valid {
			tmp.ClosedAt = &types.JSONTime{Time: room.ClosedAt.Time}
		}
		result = append(result, tmp)
	}

	return
}

func (s *service) ListStaffRoom(ctx context.Context, params model.ListStaffRoomParams, filterParams types.FilterStaffRoomParams) (rooms []model.ListStaffRoomRow, count int64, err error) {
	rooms = make([]model.ListStaffRoomRow, 0)
	err = s.repo.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		var err2 error

		_, err2 = tx.Exec("SET @staffId = ?", filterParams.StaffID)
		if err2 != nil {
			return err2
		}

		rooms, err2 = s.repo.WithTx(tx).ListStaffRoom(ctx, params)
		if err2 != nil {
			return err2
		}

		count, err2 = s.repo.WithTx(tx).CountListStaffRoom(ctx, params.Status)
		if err2 != nil {
			return err2
		}

		return nil
	})

	return
}

// GetStaffRooms 獲取客服尚未關閉服務的 房間ID 列表
func (s *service) GetStaffRooms(ctx context.Context, staffId int64) ([]int64, error) {
	return s.repo.GetStaffRoom(ctx, staffId)
}

func (s *service) TransferRoom(ctx context.Context, staffId, roomId, toStaffId int64) error {
	room, err := s.repo.GetRoom(ctx, roomId)
	if err != nil {
		return err
	}

	if room.StaffID != staffId {
		return errors.New("operation error")
	}

	if room.StaffID == toStaffId {
		return errors.New("no need to transfer")
	}

	staff, err := s.repo.GetStaff(ctx, toStaffId)
	if err != nil {
		return err
	}

	if staff.ServingStatus != types.StaffServingStatusServing {
		return errors.New("staff not available")
	}

	if err = s.repo.UpdateRoomStaff(ctx, model.UpdateRoomStaffParams{
		StaffID: toStaffId,
		ID:      roomId,
	}); err != nil {
		return err
	}

	event := pkg.StaffEventInfo{
		Event: pkg.StaffEventTransferRoom,
		Payload: pkg.StaffEventPayload{
			StaffID: &toStaffId,
			RoomID:  &roomId,
		},
	}

	payload, _ := json.Marshal(event)

	return s.redis.Publish(ctx, "event:staff", payload)
}

type ServiceParams struct {
	fx.In

	Redis     ifaceTool.IRedis
	Lua       iface.ILusScript
	MemberSvc iface.IMemberService
	Repo      iface.IRepository
}

func NewService(p ServiceParams) iface.IRoomService {
	return &service{
		redis:     p.Redis,
		lua:       p.Lua,
		memberSvc: p.MemberSvc,
		repo:      p.Repo,
	}
}

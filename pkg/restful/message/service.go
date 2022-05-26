package message

import (
	"context"
	"cs-api/pkg"
	iface2 "cs-api/pkg/interface"
	"cs-api/pkg/model"
	"cs-api/pkg/types"
	iface "github.com/AndySu1021/go-util/interface"
	"go.mongodb.org/mongo-driver/bson"
)

const Collection = "messages"

type service struct {
	repo iface.IMongoRepository
}

func (s *service) CreateMessage(ctx context.Context, message model.Message) error {
	return s.repo.InsertOne(ctx, Collection, message)
}

func (s *service) ListRoomMessage(ctx context.Context, roomId int64, clientType pkg.ClientType) (messages []model.Message, err error) {
	messages = make([]model.Message, 0)

	filter := bson.M{"room_id": roomId}
	if clientType == pkg.ClientTypeMember {
		filter["type"] = bson.M{"$ne": 1}
	}

	if err = s.repo.ListAll(ctx, Collection, &messages, filter); err != nil {
		return
	}

	return
}

func (s *service) ListMessage(ctx context.Context, params types.ListMessageParams) (messages []model.Message, count int64, err error) {
	messages = make([]model.Message, 0)

	filter := bson.M{}
	if params.RoomID != 0 {
		filter["room_id"] = params.RoomID
	}
	if params.StaffID != 0 {
		filter["type"] = model.MessageTypeStaff
		filter["sender_id"] = params.StaffID
	}
	if params.Content != "" {
		filter["content"] = bson.M{"$regex": params.Content}
	}

	count, err = s.repo.List(ctx, Collection, &messages, filter, params.Page, params.PageSize)
	if err != nil {
		return
	}

	return
}

func NewService(Repo iface.IMongoRepository) iface2.IMessageService {
	return &service{
		repo: Repo,
	}
}

package converter

import (
	"cs-api/pkg/model"
	"encoding/json"
)

var MessageTypeDtoMapping = map[model.MessageType]MessageType{
	model.MessageTypeSystem: MessageTypeSystem,
	model.MessageTypeMember: MessageTypeMember,
	model.MessageTypeStaff:  MessageTypeStaff,
}

var MessageContentTypeDtoMapping = map[model.MessageContentType]MessageContentType{
	model.MessageContentTypeTyping:       MessageContentTypeTyping,
	model.MessageContentTypeText:         MessageContentTypeText,
	model.MessageContentTypeImage:        MessageContentTypeImage,
	model.MessageContentTypeScore:        MessageContentTypeScore,
	model.MessageContentTypeJoin:         MessageContentTypeJoin,
	model.MessageContentTypeLeave:        MessageContentTypeLeave,
	model.MessageContentTypeNoStaff:      MessageContentTypeNoStaff,
	model.MessageContentTypeRoomClosed:   MessageContentTypeRoomClosed,
	model.MessageContentTypeRoomAccepted: MessageContentTypeRoomAccepted,
}

func (resp *ListRoomMessageResp) FromDto(messages []model.Message) {
	resp.Messages = make([]*Message, 0)

	for _, message := range messages {
		tmp := Message{
			ID:          message.ID,
			MessageType: MessageTypeDtoMapping[message.Type],
			RoomID:      message.RoomID,
			SenderName:  message.SenderName,
			ContentType: MessageContentTypeDtoMapping[message.ContentType],
			Content:     message.Content,
			Timestamp:   message.Timestamp,
		}

		if message.ExtraInfo != nil {
			tmp.ExtraInfo = &MessageExtraInfo{
				ClientName: message.ExtraInfo.ClientName,
			}
		}

		resp.Messages = append(resp.Messages, &tmp)
	}
}

func (resp *ListMessageResp) FromDto(messages []model.Message, pagination PaginationInput, total int64) {
	resp.Messages = make([]*Message, 0)

	for _, message := range messages {
		tmp := Message{
			ID:          message.ID,
			MessageType: MessageTypeDtoMapping[message.Type],
			RoomID:      message.RoomID,
			SenderName:  message.SenderName,
			ContentType: MessageContentTypeDtoMapping[message.ContentType],
			Content:     message.Content,
			Timestamp:   message.Timestamp,
		}

		if message.ExtraInfo != nil {
			tmp.ExtraInfo = &MessageExtraInfo{
				ClientName: message.ExtraInfo.ClientName,
			}
		}

		resp.Messages = append(resp.Messages, &tmp)
	}

	resp.Pagination = &Pagination{
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
		Total:    total,
	}
}

func ConvertDtoMsgToGql(message model.Message) []byte {
	tmp := Message{
		ID:          message.ID,
		MessageType: MessageTypeDtoMapping[message.Type],
		RoomID:      message.RoomID,
		SenderName:  message.SenderName,
		ContentType: MessageContentTypeDtoMapping[message.ContentType],
		Content:     message.Content,
		Timestamp:   message.Timestamp,
	}

	if message.ExtraInfo != nil {
		tmp.ExtraInfo = &MessageExtraInfo{
			ClientName: message.ExtraInfo.ClientName,
		}
	}

	jsonStr, _ := json.Marshal(tmp)
	return jsonStr
}

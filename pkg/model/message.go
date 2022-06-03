package model

import (
	"cs-api/pkg/types"
	"time"
)

type Message struct {
	ID        string          `bson:"_id,omitempty" json:"id,omitempty"`
	OpType    types.OpType    `bson:"op_type" json:"op_type,omitempty"`
	Payload   types.JSONField `bson:"payload" json:"payload,omitempty"`
	Timestamp int64           `bson:"timestamp" json:"timestamp,omitempty"`
	CreatedAt time.Time       `bson:"created_at" json:"created_at,omitempty"`
}

package models

import (
	"time"

	"github.com/Krystian19/cactus-core/proto"
)

const tableName = "EpisodesSeen"

// EpisodeSeen : EpisodeSeen model
type EpisodeSeen struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations
	EpisodeID int64
}

// ToProto : Translates EpisodeSeen struct into a proto EpisodeSeen struct
func (e *EpisodeSeen) ToProto() *proto.EpisodeSeen {
	return &proto.EpisodeSeen{
		Id:        e.ID,
		CreatedAt: e.CreatedAt.String(),
		UpdatedAt: e.UpdatedAt.String(),

		// Relations
		EpisodeId: e.EpisodeID,
	}
}

// TableName : Sets the model's tablename
func (EpisodeSeen) TableName() string {
	return tableName
}

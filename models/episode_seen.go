package models

import (
	"time"

	"github.com/Krystian19/cactus-core/proto"
)

const tableName = "EpisodesSeen"

// EpisodeSeen : EpisodeSeen model
type EpisodeSeen struct {
	*proto.EpisodeSeen
}

// EpisodeSeenFields : Use it to create EpisodeSeen record
type EpisodeSeenFields struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations
	EpisodeID int64
}

// TableName : Sets the model's tablename
func (EpisodeSeen) TableName() string {
	return tableName
}

// TableName : Sets the model's tablename
func (EpisodeSeenFields) TableName() string {
	return tableName
}

package models

import (
	"github.com/Krystian19/cactus-core/proto"
)

// EpisodeSeen : EpisodeSeen model
type EpisodeSeen struct {
	*proto.EpisodeSeen
}

// TableName : Sets the model's tablename
func (EpisodeSeen) TableName() string {
	return "EpisodesSeen"
}

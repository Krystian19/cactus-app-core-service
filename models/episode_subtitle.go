package models

import (
	"github.com/Krystian19/cactus-core/proto"
)

// EpisodeSubtitle : EpisodeSubtitle model
type EpisodeSubtitle struct {
	*proto.EpisodeSubtitle
}

// TableName : Sets the model's tablename
func (EpisodeSubtitle) TableName() string {
	return "EpisodeSubtitles"
}

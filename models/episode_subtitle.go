package models

import (
	"github.com/Krystian19/cactus-core/proto"
)

// EpisodeSubtitle : EpisodeSubtitle model
type EpisodeSubtitle struct {
	ID           int64
	SubtitleCode string

	// Relations
	EpisodeID  int64
	LanguageID int64
}

// ToProto : Translates EpisodeSubtitle struct into a proto EpisodeSubtitle struct
func (e *EpisodeSubtitle) ToProto() *proto.EpisodeSubtitle {
	return &proto.EpisodeSubtitle{
		Id:           e.ID,
		SubtitleCode: e.SubtitleCode,

		// Relations
		EpisodeId:  e.EpisodeID,
		LanguageId: e.LanguageID,
	}
}

// TableName : Sets the model's tablename
func (EpisodeSubtitle) TableName() string {
	return "EpisodeSubtitles"
}

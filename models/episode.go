package models

import (
	"time"

	"github.com/Krystian19/cactus-core/proto"
)

// Episode : Episode model
type Episode struct {
	ID           int64
	Thumbnail    string
	EpisodeOrder int64
	EpisodeCode  string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Relations
	ReleaseID int64
}

// ToProto : Translates Episode struct into a proto Episode struct
func (e *Episode) ToProto() *proto.Episode {
	return &proto.Episode{
		Id:           e.ID,
		Thumbnail:    e.Thumbnail,
		EpisodeOrder: e.EpisodeOrder,
		EpisodeCode:  e.EpisodeCode,
		CreatedAt:    e.CreatedAt.String(),
		UpdatedAt:    e.UpdatedAt.String(),

		// Relations
		ReleaseId: e.ReleaseID,
	}
}

// TableName : Sets the model's tablename
func (Episode) TableName() string {
	return "Episodes"
}

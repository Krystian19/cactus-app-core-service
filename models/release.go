package models

import (
	"time"

	"github.com/Krystian19/cactus-core/proto"
)

// Release : Release model
type Release struct {
	ID            int64
	Title         string
	Poster        string
	Background    string
	ReleaseOrder  int64
	StartedAiring string
	StoppedAiring string
	CreatedAt     time.Time
	UpdatedAt     time.Time

	// Relations
	AnimeID       int64
	ReleaseTypeID int64
}

// ToProto : Translates Release struct into a proto Release struct
func (r *Release) ToProto() *proto.Release {
	return &proto.Release{
		Id:            r.ID,
		Title:         r.Title,
		Poster:        r.Poster,
		Background:    r.Background,
		ReleaseOrder:  r.ReleaseOrder,
		StartedAiring: r.StartedAiring,
		StoppedAiring: r.StoppedAiring,
		CreatedAt:     r.CreatedAt.String(),
		UpdatedAt:     r.UpdatedAt.String(),
	}
}

// TableName : Sets the model's tablename
func (Release) TableName() string {
	return "Releases"
}

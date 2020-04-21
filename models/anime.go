package models

import (
	"time"

	"github.com/Krystian19/cactus-core/proto"
)

// Anime : Anime model
type Anime struct {
	ID        int64
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ToProto : Translates Anime struct into a proto Anime struct
func (a *Anime) ToProto() *proto.Anime {
	return &proto.Anime{
		Id:        a.ID,
		Title:     a.Title,
		CreatedAt: a.CreatedAt.String(),
		UpdatedAt: a.UpdatedAt.String(),
	}
}

// TableName : Sets the model's tablename
func (Anime) TableName() string {
	return "Animes"
}

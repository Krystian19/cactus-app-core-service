package models

import (
	"github.com/Krystian19/cactus-core/proto"
)

// Anime : Anime model
type Anime struct {
	ID        int64
	Title     string
	CreatedAt string
	UpdatedAt string
}

// ToProto : Translates Anime struct into a proto Anime struct
func (a *Anime) ToProto() *proto.Anime {
	return &proto.Anime{
		Id:        a.ID,
		Title:     a.Title,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

// TableName : Sets the model's tablename
func (Anime) TableName() string {
	return "Animes"
}

package models

import (
	"github.com/Krystian19/cactus-core/proto"
)

// Genre : Genre model
type Genre struct {
	ID        int64
	Title     string
	Thumbnail string
}

// ToProto : Translates Genre struct into a proto Genre struct
func (a *Genre) ToProto() *proto.Genre {
	return &proto.Genre{
		Id:        a.ID,
		Title:     a.Title,
		Thumbnail: a.Thumbnail,
	}
}

// TableName : Sets the model's tablename
func (Genre) TableName() string {
	return "Genres"
}

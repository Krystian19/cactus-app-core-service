package models

import (
	"github.com/Krystian19/cactus-core/proto"
)

// ReleaseType : ReleaseType model
type ReleaseType struct {
	ID    int64
	Title string
}

// ToProto : Translates ReleaseType struct into a proto ReleaseType struct
func (r *ReleaseType) ToProto() *proto.ReleaseType {
	return &proto.ReleaseType{
		Id:    r.ID,
		Title: r.Title,
	}
}

// TableName : Sets the model's tablename
func (ReleaseType) TableName() string {
	return "ReleaseTypes"
}

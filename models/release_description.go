package models

import (
	"github.com/Krystian19/cactus-core/proto"
)

// ReleaseDescription : ReleaseDescription model
type ReleaseDescription struct {
	ID          int64
	Description string

	// Relations
	ReleaseID  int64
	LanguageID int64
}

// ToProto : Translates ReleaseDescription struct into a proto ReleaseDescription struct
func (l *ReleaseDescription) ToProto() *proto.ReleaseDescription {
	return &proto.ReleaseDescription{
		Id:          l.ID,
		Description: l.Description,

		// Relations
		ReleaseId:  l.ReleaseID,
		LanguageId: l.LanguageID,
	}
}

// TableName : Sets the model's tablename
func (ReleaseDescription) TableName() string {
	return "ReleaseDescriptions"
}

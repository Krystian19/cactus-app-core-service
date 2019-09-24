package models

import (
	"github.com/Krystian19/cactus-core/proto"
)

// ReleaseType : ReleaseType model
type ReleaseType struct {
	*proto.ReleaseType
}

// TableName : Sets the model's tablename
func (ReleaseType) TableName() string {
	return "ReleaseTypes"
}

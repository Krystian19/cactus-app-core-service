package models

import (
	"github.com/Krystian19/cactus-core/proto"
)

// ReleaseDescription : ReleaseDescription model
type ReleaseDescription struct {
	*proto.ReleaseDescription
}

// TableName : Sets the model's tablename
func (ReleaseDescription) TableName() string {
	return "ReleaseDescriptions"
}

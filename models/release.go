package models

import (
	"github.com/Krystian19/cactus-core/proto"
)

// Release : Release model
type Release struct {
	*proto.Release
}

// TableName : Sets the model's tablename
func (Release) TableName() string {
	return "Releases"
}

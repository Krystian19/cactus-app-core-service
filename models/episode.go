package models

import (
	"github.com/Krystian19/cactus-core/proto"
)

// Episode : Episode model
type Episode struct {
	*proto.Episode
}

// TableName : Sets the model's tablename
func (Episode) TableName() string {
	return "Episodes"
}

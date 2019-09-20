package models

import (
	"github.com/Krystian19/cactus-core/proto"
)

// Genre : Genre model
type Genre struct {
	*proto.Genre
}

// TableName : Sets the model's tablename
func (Genre) TableName() string {
	return "Genres"
}

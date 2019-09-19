package models

import (
	"github.com/Krystian19/cactus-core/proto"
)

// Anime : Anime model
type Anime struct {
	*proto.Anime
}

// TableName : Sets the model's tablename
func (Anime) TableName() string {
	return "Animes"
}

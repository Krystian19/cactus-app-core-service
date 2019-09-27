package models

import (
	"github.com/Krystian19/cactus-core/proto"
)

// Language : Language model
type Language struct {
	*proto.Language
}

// TableName : Sets the model's tablename
func (Language) TableName() string {
	return "Languages"
}

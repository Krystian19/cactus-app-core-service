package models

import "github.com/Krystian19/cactus-core/proto"

// Language : Language model
type Language struct {
	ID      int64
	Name    string
	IsoCode string
}

// ToProto : Translates Language struct into a proto Language struct
func (l *Language) ToProto() *proto.Language {
	return &proto.Language{
		Id:      l.ID,
		Name:    l.Name,
		IsoCode: l.IsoCode,
	}
}

// TableName : Sets the model's tablename
func (Language) TableName() string {
	return "Languages"
}

package models

import "time"

// EpisodeSeen : EpisodeSeen model
type EpisodeSeen struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations
	EpisodeID int64
}

// TableName : Sets the model's tablename
func (EpisodeSeen) TableName() string {
	return "EpisodesSeen"
}

package models

import "gorm.io/gorm"

type Victim struct {
	gorm.Model
	FullName       string `json:"full_name"`
	CauseOfDeath   string `json:"cause_of_death"`
	Details        string `json:"details"`
	ImageURL       string `json:"image_url"`
	IsDead         bool   `json:"is_dead"`
	DeathTimestamp int64  `json:"death_timestamp"`
}

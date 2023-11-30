package repository

import "time"

type GainProjection struct {
	Id          string
	CreatedAt   time.Time
	PayIn       time.Time
	Description string
	Value       float64
	IsPassive   bool
	IsDone      bool
	UserId      string
	Category    GainCategory
}

type GainCategory struct {
	Id       uint
	Category string
}

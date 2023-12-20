package repository

import "time"

type GainProjection struct {
	Id            string
	CreatedAt     time.Time
	PayIn         time.Time
	Description   string
	Value         float64
	IsPassive     bool
	IsAlreadyDone bool
	UserId        string
	Category      GainCategory
}

type Gain struct {
	Id               string
	CreatedAt        time.Time
	PayIn            time.Time
	Description      string
	Value            float64
	IsPassive        bool
	GainProjectionId string
	UserId           string
	Category         GainCategory
}

type GainCategory struct {
	Id       uint
	Category string
}

type QueryParams struct {
	month  uint
	year   uint
	limit  uint
	offset uint
}

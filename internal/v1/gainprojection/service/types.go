package service

import "time"

type CreateRequest struct {
	PayIn       time.Time `json:"pay_in"`
	Description string    `json:"description"`
	Value       float64   `json:"value"`
	IsPassive   bool      `json:"is_passive"`
	Recurrence  uint      `json:"recurrence"`
	CategoryId  uint      `json:"category_id"`
}

type UpdateRequest struct {
	PayIn       time.Time `json:"pay_in"`
	Description string    `json:"description"`
	Value       float64   `json:"value"`
	IsPassive   bool      `json:"is_passive"`
	CategoryId  uint      `json:"category_id"`
}

type GainProjectionResponse struct {
	Id          string           `json:"id"`
	PayIn       time.Time        `json:"pay_in"`
	Description string           `json:"description"`
	Value       float64          `json:"value"`
	IsPassive   bool             `json:"is_passive"`
	Recurrence  uint             `json:"recurrence,omitempty"`
	Category    CategoryResponse `json:"category"`
}

type CategoryResponse struct {
	Id       uint   `json:"id"`
	Category string `json:"category"`
}

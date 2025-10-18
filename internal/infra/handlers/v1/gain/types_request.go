package gain

import "time"

type CreateRequest struct {
	PayIn       time.Time `json:"pay_in"`
	Description string    `json:"description"`
	Value       float64   `json:"value"`
	IsPassive   bool      `json:"is_passive"`
	CategoryId  uint      `json:"category_id"`
}

type EditRequest struct {
	PayIn       time.Time `json:"pay_in"`
	Description string    `json:"description"`
	Value       float64   `json:"value"`
	IsPassive   bool      `json:"is_passive"`
	CategoryId  uint      `json:"category_id"`
}

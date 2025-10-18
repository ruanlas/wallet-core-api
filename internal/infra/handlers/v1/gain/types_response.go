package gain

import "time"

type GainResponse struct {
	Id               string           `json:"id"`
	GainProjectionId string           `json:"gain_projection_id,omitempty"`
	PayIn            time.Time        `json:"pay_in"`
	Description      string           `json:"description"`
	Value            float64          `json:"value"`
	IsPassive        bool             `json:"is_passive"`
	Category         CategoryResponse `json:"category"`
}

type CategoryResponse struct {
	Id       uint   `json:"id"`
	Category string `json:"category"`
}

type GainPaginatedResponse struct {
	CurrentPage  uint           `json:"current_page"`
	TotalPages   uint           `json:"total_pages"`
	TotalRecords uint           `json:"total_records"`
	PageLimit    uint           `json:"page_limit"`
	Records      []GainResponse `json:"records"`
}

type DeleteResponse struct{}

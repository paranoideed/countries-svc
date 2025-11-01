package models

import "time"

type Policy struct {
	CountryID string

	CitiesAllowedStatuses []string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p Policy) IsNil() bool {
	return p.CountryID == ""
}

type PoliciesCollection struct {
	Data  []Policy `json:"data"`
	Page  uint64   `json:"page"`
	Size  uint64   `json:"size"`
	Total uint64   `json:"total"`
}

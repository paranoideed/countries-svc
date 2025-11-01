package models

import (
	"time"
)

type Country struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c Country) IsNil() bool {
	return c.ID == ""
}

type CountriesCollection struct {
	Data  []Country `json:"data"`
	Page  uint64    `json:"page"`
	Size  uint64    `json:"size"`
	Total uint64    `json:"total"`
}

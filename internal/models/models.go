package models

import (
	"time"
)

type City struct {
	ID        string    `json:"id" gorm:"type:varchar(255);primaryKey;uniqueIndex:idx_city_unique"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`

	Name        string  `json:"name" gorm:"type:varchar(255);uniqueIndex:idx_city_unique"`
	RegionCode  string  `json:"region_code" gorm:"type:varchar(10);uniqueIndex:idx_city_unique"`
	CountryCode string  `json:"country_code" gorm:"type:varchar(3);uniqueIndex:idx_city_unique;index:idx_country_code"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Population  int     `json:"population" example:"3000000"`
}

type PaginationMeta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

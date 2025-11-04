package models

import (
	"time"
)

// @name City
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

// @name Country
type Country struct {
	ID        string    `json:"id" gorm:"type:varchar(255);primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`

	Code       string  `json:"code" example:"CM" gorm:"type:varchar(3);uniqueIndex:idx_code;column:country_code"`
	Name       string  `json:"name" example:"Republic of Cameroon" gorm:"type:varchar(255)"`
	ISO3Code   string  `json:"iso3_code" example:"CMR" gorm:"type:varchar(3)"`
	Capital    string  `json:"capital" example:"Yaounde" gorm:"type:varchar(255)"`
	Area       float64 `json:"area" example:"475000.7"`
	Population int     `json:"population" example:"3000000"`
	Continent  string  `json:"continent" example:"AFRICA" gorm:"type:varchar(50)"`
	TLD        string  `json:"tld" example:"cm" gorm:"type:varchar(10)"`
	Currency   string  `json:"currency" example:"Franc" gorm:"type:varchar(50)"`
	Phone      string  `json:"phone" example:"237" gorm:"type:varchar(20)"`
	Languages  string  `json:"languages" example:"en-CM,fr-CM" gorm:"type:varchar(255)"`
	Neighbours string  `json:"neighbours" example:"TD,CF,GA,GQ,CG,NG" gorm:"type:varchar(255)"`
}

// @name Region
type Region struct {
	ID        string    `json:"id" gorm:"type:varchar(255);primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`

	Name               string `json:"name" example:"Centre" gorm:"type:varchar(255)"`
	CountryCode        string `json:"country_code" example:"CM" gorm:"type:varchar(3)"`
	TopLevelRegionCode string `json:"top_level_region_code" example:"02" gorm:"type:varchar(10)"`
	Level              string `json:"level" example:"ADM1" gorm:"type:varchar(10)"`
}

// @name PaginationMeta
type PaginationMeta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

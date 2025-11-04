package city

import (
	"errors"
	"fmt"
	"time"

	"github.com/melkishengue/gotemplate/internal/models"
	"github.com/melkishengue/gotemplate/pkg/utils"
	"gorm.io/gorm"
)

type CityFilters struct {
	ID          string // Filter by city ID
	CountryCode string // Filter by country code
	Name        string // Fuzzy search by name
	Limit       int    // Pagination limit
	Offset      int    // Pagination offset
}

type City struct {
	ID        string    `json:"id" gorm:"type:varchar(255);primaryKey;uniqueIndex:idx_city_unique"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`

	Name        string `json:"name" gorm:"type:varchar(255);uniqueIndex:idx_city_unique"`
	RegionCode  string `json:"region_code" gorm:"type:varchar(10);uniqueIndex:idx_city_unique"`
	CountryCode string `json:"country_code" gorm:"type:varchar(3);uniqueIndex:idx_city_unique"`
	Population  int    `json:"population" example:"3000000"`
}

type CitiesResponse struct {
	Data []City                `json:"data"`
	Meta models.PaginationMeta `json:"meta"`
}

func GetCities(db *gorm.DB, filters CityFilters) ([]City, int64, error) {
	var cities []City
	var total int64

	query := db.Model(&City{})

	if filters.ID != "" {
		query = query.Where("id = ?", filters.ID)
	}

	if filters.CountryCode != "" {
		query = query.Where("country_code = ?", filters.CountryCode)
	}

	if filters.Name != "" {
		similarityThreshold := utils.GetEnvOrDie("SIMILARITY_PERCENTAGE")
		query = query.Where(
			fmt.Sprintf("similarity(name, ?) >= %s", similarityThreshold),
			filters.Name,
		).Order(gorm.Expr("similarity(name, ?) DESC", filters.Name))
	}

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count cities: %w", err)
	}

	query = query.Limit(filters.Limit).Offset(filters.Offset)

	if err := query.Find(&cities).Error; err != nil {
		return nil, 0, fmt.Errorf("gorm query failed: %w", err)
	}

	return cities, total, nil
}

func GetCityByID(db *gorm.DB, id string) (*City, error) {
	var city City

	if err := db.Model(&City{}).Where("id = ?", id).First(&city).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("city not found")
		}
		return nil, fmt.Errorf("gorm query failed: %w", err)
	}

	return &city, nil
}

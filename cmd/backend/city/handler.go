package city

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	db "github.com/melkishengue/gotemplate/internal/database"
	"github.com/melkishengue/gotemplate/internal/models"
)

// @Summary Get cities
// @Description Get a list of cities with optional filtering including radius search
// @Accept json
// @Produce json
// @Param country_code path string true "Country code - Example: AD, CM, ES, GA, US, CN, RU, FR, DE, etc..."
// @Param name query string false "City name (fuzzy search)"
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} CitiesResponse
// @Router /countries/{country_code}/cities [get]
func CitiesHandler(c *gin.Context) {
	const (
		defaultLimit = 10
		maxLimit     = 100
	)

	cityID := c.Param("id")

	countryCode := c.Param("country_code")
	name := c.Query("name")

	limit := defaultLimit
	if limitStr := c.Query("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			if parsedLimit > maxLimit {
				limit = maxLimit
			} else {
				limit = parsedLimit
			}
		}
	}

	offset := 0
	if offsetStr := c.Query("offset"); offsetStr != "" {
		if parsedOffset, err := strconv.Atoi(offsetStr); err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	filters := CityFilters{
		ID:          cityID,
		CountryCode: countryCode,
		Name:        name,
		Limit:       limit,
		Offset:      offset,
	}

	// Query using GORM
	cities, total, err := GetCities(db.Connection, filters)
	if err != nil {
		slog.Error("Error querying cities", "err", slog.AnyValue(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	response := CitiesResponse{
		Data: cities,
		Meta: models.PaginationMeta{
			Limit:  limit,
			Offset: offset,
			Total:  int(total),
		},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get a city by ID
// @Description Get a city by its ID
// @Accept json
// @Produce json
// @Param id path string true "City ID"
// @Param country_code path string true "Country code - Example: AD, CM, ES, GA, US, CN, RU, FR, DE, etc..."
// @Success 200 {object} City
// @Router /countries/{country_code}/cities/{id} [get]
func CityHandler(c *gin.Context) {
	cityID := c.Param("id")

	city, err := GetCityByID(db.Connection, cityID)
	if err != nil {
		if err.Error() == "city not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
			return
		}
		slog.Error("Error querying city", "id", cityID, "err", slog.AnyValue(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, city)
}

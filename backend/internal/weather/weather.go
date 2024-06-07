package weather

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	findWeather(hometown string) (bson.M, error)
	updateWeather(hometown string, weather HistoryWeatherResponse) error
	addWeather(hometown string, weather HistoryWeatherResponse) error
}

type Service interface {
	findWeather(ctx context.Context, hometown, startDate string) (HistoryWeatherResponse, error)
	addWeather(ctx context.Context, hometown string) error
	updateWeather(ctx context.Context, hometown string) error
}

type Handler interface {
	FindWeather(c *gin.Context)
	UpdateWeather(c *gin.Context)
}

type GeocodingResponse struct {
	Results []struct {
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat float32 `json:"lat"`
				Lng float32 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
	Status string `json:"status"`
}

type HistoryWeatherResponse struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Hourly    struct {
		Time []string  `json:"time"`
		Temp []float32 `json:"temperature_2m"`
	} `json:"hourly"`
}

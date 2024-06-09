package weather

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	findWeather(hometown string) (RepositoryBson, error)
	updateWeather(hometown string, weather HistoryWeatherResponse) error
	addWeather(hometown string, weather HistoryWeatherResponse) error
}

type Service interface {
	findWeather(ctx context.Context, hometown, startDate string) (HistoryWeatherResponse, error)
	addWeather(ctx context.Context, hometown, startDate string) error
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
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
	Status string `json:"status"`
}
type HourlyTemps struct {
	Time []string  `json:"time" bson:"time"`
	Temp []float64 `json:"temp" bson:"temp"`
}
type HistoryWeatherResponse struct {
	Latitude  float64     `json:"latitude" bson:"latitude"`
	Longitude float64     `json:"longitude" bson:"longitude"`
	Hourly    HourlyTemps `json:"hourly" bson:"hourly"`
}
type WeatherData struct {
	Hourly struct {
		Time          []string  `json:"time"`
		Temperature2m []float64 `json:"temperature_2m"`
	} `json:"hourly"`
}

type RepositoryBson struct {
	Town    string `json:"town" bson:"town"`
	Weather struct {
		Latitude  float64 `json:"latitude" bson:"latitude"`
		Longitude float64 `json:"longitude" bson:"longitude"`
		Hourly    struct {
			Time []string  `json:"time" bson:"time"`
			Temp []float64 `json:"temp" bson:"temp"`
		} `json:"hourly" bson:"hourly"`
	} `json:"weather" bson:"weather"`
}

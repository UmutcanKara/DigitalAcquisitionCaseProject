package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(r Repository) Service { return &service{r, time.Duration(5) * time.Second} }

func (s *service) findWeather(ctx context.Context, hometown, startDate string) (HistoryWeatherResponse, error) {
	res := HistoryWeatherResponse{}
	_, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	weatherBson, err := s.Repository.findWeather(hometown)
	if err != nil {
		return HistoryWeatherResponse{}, err
	}
	weatherBytes, err := json.Marshal(weatherBson)
	if err != nil {
		return HistoryWeatherResponse{}, err
	}
	err = bson.Unmarshal(weatherBytes, res)
	if err != nil {
		return HistoryWeatherResponse{}, err
	}
	res, err = sliceFromStartDate(res, startDate)
	if err != nil {
		return HistoryWeatherResponse{}, err
	}
	return res, nil
}

func (s *service) updateWeather(ctx context.Context, hometown string) error {
	_, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	res := HistoryWeatherResponse{}
	googleApiKey := os.Getenv("")
	weatherApiUrl := os.Getenv("")

	geocodeResult, err := getLongLat(hometown, googleApiKey)
	if err != nil {
		return err
	}
	location := geocodeResult.Results[0].Geometry.Location
	res, err = getWeatherHistory(location.Lat, location.Lng, weatherApiUrl)

	err = s.Repository.updateWeather(hometown, res)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) addWeather(ctx context.Context, hometown string) error {
	_, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	res := HistoryWeatherResponse{}
	googleApiKey := os.Getenv("")
	weatherApiUrl := os.Getenv("")

	geocodeResult, err := getLongLat(hometown, googleApiKey)
	if err != nil {
		return err
	}
	location := geocodeResult.Results[0].Geometry.Location
	res, err = getWeatherHistory(location.Lat, location.Lng, weatherApiUrl)

	err = s.Repository.updateWeather(hometown, res)
	if err != nil {
		return err
	}
	return nil
}

func getWeatherHistory(lat, long float32, apiUrl string) (HistoryWeatherResponse, error) {
	weatherData := HistoryWeatherResponse{}
	now := time.Now()
	nowString := now.Format(time.DateOnly)
	pastString := now.AddDate(-5, 0, 0).Format(time.DateOnly)
	url := fmt.Sprintf("%s?longitude=%f&latitude=%f&start_date=%s&end_date=%s", apiUrl, lat, long, nowString, pastString)
	resp, err := http.Get(url)
	if err != nil {
		return HistoryWeatherResponse{}, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return HistoryWeatherResponse{}, err
	}
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return HistoryWeatherResponse{}, err
	}
	return weatherData, nil
}

func getLongLat(city, mapsKey string) (GeocodingResponse, error) {
	geolocData := GeocodingResponse{}
	mapsUrl := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", strings.Replace(city, " ", "+", -1), mapsKey)

	resp, err := http.Get(mapsUrl)
	if err != nil {
		return GeocodingResponse{}, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GeocodingResponse{}, err
	}
	err = json.Unmarshal(body, &geolocData)
	if err != nil {
		return GeocodingResponse{}, err
	}
	return geolocData, nil
}

func sliceFromStartDate(weather HistoryWeatherResponse, startDate string) (HistoryWeatherResponse, error) {
	// Parse the start date
	start, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		return HistoryWeatherResponse{}, fmt.Errorf("invalid start date format: %w", err)
	}

	var filteredTimes []string
	var filteredTemps []float32

	// Iterate through the Time slice and filter based on the start date
	for i, t := range weather.Hourly.Time {
		parsedTime, err := time.Parse(time.RFC3339, t)
		if err != nil {
			return HistoryWeatherResponse{}, fmt.Errorf("invalid time format in data: %w", err)
		}

		// Check if the parsed time is after the start date
		if parsedTime.After(start) {
			filteredTimes = append(filteredTimes, t)
			filteredTemps = append(filteredTemps, weather.Hourly.Temp[i])
		}
	}

	// Create a new HistoryWeatherResponse with the filtered data
	filteredData := HistoryWeatherResponse{
		Latitude:  weather.Latitude,
		Longitude: weather.Longitude,
		Hourly: struct {
			Time []string  `json:"time"`
			Temp []float32 `json:"temperature_2m"`
		}{
			Time: filteredTimes,
			Temp: filteredTemps,
		},
	}

	return filteredData, nil
}

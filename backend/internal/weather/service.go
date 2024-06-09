package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"googlemaps.github.io/maps"
	"io"
	"net/http"
	"os"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(r Repository) Service { return &service{r, time.Duration(5) * time.Second} }

func (s *service) findWeather(ctx context.Context, hometown, startDate string) (HistoryWeatherResponse, error) {
	var res HistoryWeatherResponse
	_, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	weatherData, err := s.Repository.findWeather(hometown)
	if err != nil {
		return HistoryWeatherResponse{}, err
	}
	err = sliceFromStartDates(&weatherData, startDate)
	if err != nil {
		return HistoryWeatherResponse{}, err
	}
	res = HistoryWeatherResponse{
		Latitude:  weatherData.Weather.Latitude,
		Longitude: weatherData.Weather.Longitude,
		Hourly: HourlyTemps{
			Time: weatherData.Weather.Hourly.Time,
			Temp: weatherData.Weather.Hourly.Temp,
		},
	}
	return res, nil
}

func (s *service) updateWeather(ctx context.Context, hometown string) error {
	_, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	res := HistoryWeatherResponse{}
	googleApiKey := os.Getenv("GOOGLE_API_KEY")
	weatherApiUrl := os.Getenv("WEATHER_API_URL")

	geocodeResult, err := getLongLat(hometown, googleApiKey)
	if err != nil {
		return err
	}
	location := geocodeResult.Geometry.Location
	res, err = getWeatherHistory(location.Lat, location.Lng, weatherApiUrl)

	err = s.Repository.updateWeather(hometown, res)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) addWeather(ctx context.Context, hometown, startDate string) error {
	_, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	res := HistoryWeatherResponse{}
	googleApiKey := os.Getenv("GOOGLE_API_KEY")
	weatherApiUrl := os.Getenv("WEATHER_API_URL")

	geocodeResult, err := getLongLat(hometown, googleApiKey)
	if err != nil {
		return err
	}
	location := geocodeResult.Geometry.Location
	res, err = getWeatherHistory(location.Lat, location.Lng, weatherApiUrl)
	if err != nil {
		return err
	}
	err = s.Repository.addWeather(hometown, res)
	if err != nil {
		return err
	}
	return nil
}

func getWeatherHistory(lat, long float64, apiUrl string) (HistoryWeatherResponse, error) {
	now := time.Now()
	nowString := now.AddDate(0, 0, -5).Format(time.DateOnly)
	pastString := now.AddDate(-5, 0, 0).Format(time.DateOnly)
	url := fmt.Sprintf("%s?longitude=%f&latitude=%f&start_date=%s&end_date=%s&hourly=temperature_2m", apiUrl, lat, long, pastString, nowString)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return HistoryWeatherResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Received non-200 response code: %d\n", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Response body:", string(body))
		return HistoryWeatherResponse{}, err
	}

	var data WeatherData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return HistoryWeatherResponse{}, err
	}
	return HistoryWeatherResponse{
		Latitude:  lat,
		Longitude: long,
		Hourly: HourlyTemps{
			Time: data.Hourly.Time,
			Temp: data.Hourly.Temperature2m,
		},
	}, nil
}

func getLongLat(city, mapsKey string) (maps.GeocodingResult, error) {
	c, err := maps.NewClient(maps.WithAPIKey(mapsKey))
	if err != nil {
		return maps.GeocodingResult{}, err
	}
	resp, err := c.Geocode(context.TODO(), &maps.GeocodingRequest{
		Address:      city,
		Components:   nil,
		Bounds:       nil,
		Region:       "",
		LatLng:       nil,
		ResultType:   nil,
		LocationType: nil,
		PlaceID:      "",
		Language:     "",
		Custom:       nil,
	})
	if err != nil {
		return maps.GeocodingResult{}, err
	}
	return resp[0], nil
}

func sliceFromStartDates(weather *RepositoryBson, startDate string) error {
	// Parse the start date
	start, err := time.Parse(time.DateOnly, startDate)
	if err != nil {
		return fmt.Errorf("invalid start date format: %w", err)
	}

	var filteredTimes []string
	var filteredTemps []float64

	// Iterate through the Time slice and filter based on the start date
	for i, t := range weather.Weather.Hourly.Time {
		layout := "2006-01-02T15:04"
		parsedTime, err := time.Parse(layout, t)
		if err != nil {
			return fmt.Errorf("invalid time format in data: %w", err)
		}
		// Check if the parsed time is after the start date
		if parsedTime.After(start) && i < len(weather.Weather.Hourly.Time)-20 {
			filteredTimes = append(filteredTimes, t)
			filteredTemps = append(filteredTemps, weather.Weather.Hourly.Temp[i])
		}
	}

	// Create a new HistoryWeatherResponse with the filtered data
	//filteredData := HistoryWeatherResponse{
	//	Latitude:  weather.Weather.Latitude,
	//	Longitude: weather.Weather.Longitude,
	//	Hourly: HourlyTemps{
	//		Time: filteredTimes,
	//		Temp: filteredTemps,
	//	},
	//}
	weather.Weather.Hourly.Time = filteredTimes
	weather.Weather.Hourly.Temp = filteredTemps

	return nil
}

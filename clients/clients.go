package clients

import (
	"encoding/json"
	"fmt"
)

const BaseURL = "https://api.openweathermap.org"

type Weather struct {
	Sky string
	// Temperature float64
	// City        string
}

type OWMResponse struct {
	Weather []struct {
		Main string
	}
}

func ParseResponse(data []byte) (Weather, error) {
	var tmp OWMResponse
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return Weather{}, err
	}
	weatherTmp := tmp.Weather
	if len(weatherTmp) == 0 {
		return Weather{}, fmt.Errorf("error parsing json: not enough weather elements (%v).", data)
	}

	weather := Weather{Sky: weatherTmp[0].Main}
	return weather, nil
}

func FormatURL(baseURL, location, key string) string {
	apiURL := "%s/data/2.5/weather?q=%s&appid=%s"

	return fmt.Sprintf(apiURL, baseURL, location, key)
}

package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const BaseURL = "https://api.openweathermap.org"
const Usage = `Usage: weather LOCATION

Example: weather London,UK`

type Weather struct {
	Sky         string
	Temperature Temperature
	// City        string
}

type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

type Temperature float64

func NewClient(apikey string) *Client {
	return &Client{
		BaseURL: BaseURL,
		APIKey:  apikey,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type OWMResponse struct {
	Weather []struct {
		Main string
	}
	Main struct {
		Temp float64
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

	weather := Weather{
		Sky:         weatherTmp[0].Main,
		Temperature: Temperature(tmp.Main.Temp),
	}
	return weather, nil
}

func (c *Client) FormatURL(location string) string {
	apiURL := "%s/data/2.5/weather?q=%s&appid=%s"

	return fmt.Sprintf(apiURL, c.BaseURL, location, c.APIKey)
}

func (c *Client) GetWeather(location string) (Weather, error) {
	URL := c.FormatURL(location)
	r, err := c.HTTPClient.Get(URL)
	if err != nil {
		return Weather{}, nil
	}

	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return Weather{}, fmt.Errorf("unexpected response status %q", r.StatusCode)
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return Weather{}, err
	}
	weather, err := ParseResponse(data)
	if err != nil {
		return Weather{}, err
	}

	return weather, nil
}

func Get(location, key string) (Weather, error) {
	c := NewClient(key)
	weather, err := c.GetWeather(location)
	if err != nil {
		return Weather{}, err
	}
	return weather, nil
}

func Main() int {
	if len(os.Args) < 2 {
		fmt.Println(Usage)
		return 0
	}

	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	if key == "" {
		fmt.Fprintln(os.Stderr, "Please set the environment variable OPENWEATHERMAP_API_KEY.")
		return 1
	}

	location := os.Args[1]
	c := NewClient(key)

	weather, err := c.GetWeather(location)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("%s %.1fºC\n", weather.Sky, weather.Temperature.Celsius())
	return 0
}

func (t Temperature) Celsius() float64 {
	return float64(t) - 273.15
}

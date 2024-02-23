package clients_test

import (
	"clients"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseResponse_CorrectlyParsesResponseFromJSON(t *testing.T) {
	t.Parallel()
	want := clients.Weather{
		Sky:         "Clouds",
		Temperature: 277.29,
		// City:        "London",
	}
	data, err := os.ReadFile("testdata/weatherJSON.json")
	if err != nil {
		t.Fatal(err)
	}
	got, err := clients.ParseResponse(data)

	if !cmp.Equal(got, want) {
		t.Error(cmp.Diff(got, want))
	}
}

func TestParseResponse_ReturnsErrorGivenEmptyData(t *testing.T) {
	t.Parallel()
	_, err := clients.ParseResponse([]byte{})
	if err == nil {
		t.Fatal("want error parsing empty resopnse, got nil")
	}
}

func TestParseResponse_ReturnsErrorWhenWeatherSliceIsEmpty(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/weatherJSONIncorrect.json")
	if err != nil {
		t.Fatal(err)
	}
	_, err = clients.ParseResponse(data)

	if err == nil {
		t.Fatal("want error parsing empty resopnse, got nil")
	}
}

func TestFormatURL_FormatsCorrectURLProvidedLocationAndKey(t *testing.T) {
	t.Parallel()
	key := "woienvkdfjlsns"
	location := "London,UK"
	want := "https://api.openweathermap.org/data/2.5/weather?q=London,UK&appid=woienvkdfjlsns"

	c := clients.NewClient(key)
	got := c.FormatURL(location)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHTTPGet_SuccessfullyGetsFromLocalServer(t *testing.T) {
	t.Parallel()
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/weatherJSON.json")
	}))
	defer ts.Close()

	client := ts.Client()

	r, err := client.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		t.Error("Status code is not ok.")
	}
	want, err := os.ReadFile("testdata/weatherJSON.json")
	if err != nil {
		t.Fatal(err)
	}

	got, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(got, want) {
		t.Error(cmp.Diff(got, want))
	}
}

func TestGetWeather_RetrievesCorrectData(t *testing.T) {
	t.Parallel()
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/weatherJSON.json")
	}))
	defer ts.Close()

	client := ts.Client()
	location := "London,UK"

	c := clients.NewClient("dummyKey")
	c.BaseURL = ts.URL
	c.HTTPClient = client

	got, err := c.GetWeather(location)
	if err != nil {
		t.Fatal(err)
	}

	want := clients.Weather{Sky: "Clouds", Temperature: 277.29}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

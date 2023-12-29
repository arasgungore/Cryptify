package chart

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// RealProvider is an implementation of the chart data provider that fetches real data from an external API
type RealProvider struct {
	APIURL string
}

// GetChartData returns real chart data from the external API
func (r *RealProvider) GetChartData(currency string, startTime, endTime time.Time) (string, error) {
	// Construct API URL with parameters
	apiURL := fmt.Sprintf("%s?currency=%s&start=%s&end=%s", r.APIURL, currency, startTime.Format(time.RFC3339), endTime.Format(time.RFC3339))

	// Make HTTP request to the external API
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch real chart data: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(body), nil
}

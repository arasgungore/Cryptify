package chart

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// BinanceClient is a client for interacting with the Binance API
type BinanceClient struct {
	APIKey    string
	APISecret string
}

// BinanceResponse represents the response from the Binance API
type BinanceResponse struct {
	Chart []struct {
		Timestamp        int64   `json:"timestamp"`
		Open             float64 `json:"open"`
		Close            float64 `json:"close"`
		High             float64 `json:"high"`
		Low              float64 `json:"low"`
		Volume           float64 `json:"volume"`
		CloseTime        int64   `json:"closeTime"`
		QuoteAssetVolume float64 `json:"quoteAssetVolume"`
		NumberOfTrades   int     `json:"numberOfTrades"`
		TakerBuyBaseVol  float64 `json:"takerBuyBaseVol"`
		TakerBuyQuoteVol float64 `json:"takerBuyQuoteVol"`
	} `json:"k"`
}

// GetChartData returns real chart data from the Binance API
func (b *BinanceClient) GetChartData(symbol string, startTime, endTime time.Time) (string, error) {
	// Convert time to milliseconds
	startTimeMillis := startTime.UnixNano() / int64(time.Millisecond)
	endTimeMillis := endTime.UnixNano() / int64(time.Millisecond)

	// Binance API endpoint for klines (candlestick chart data)
	apiURL := fmt.Sprintf("https://api.binance.com/api/v3/klines?symbol=%s&interval=1h&startTime=%d&endTime=%d", symbol, startTimeMillis, endTimeMillis)

	// Create HTTP request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create Binance API request: %v", err)
	}

	// Set API key headers
	req.Header.Set("X-MBX-APIKEY", b.APIKey)

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute Binance API request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read Binance API response: %v", err)
	}

	// Parse the Binance API response
	var binanceResponse BinanceResponse
	if err := json.Unmarshal(body, &binanceResponse); err != nil {
		return "", fmt.Errorf("failed to parse Binance API response: %v", err)
	}

	// Format the chart data
	chartData := ""
	for _, kline := range binanceResponse.Chart {
		chartData += fmt.Sprintf("[%s, %.2f], ", time.Unix(0, kline.Timestamp*int64(time.Millisecond)).Format("2006-01-02T15:04:05"), kline.Close)
	}

	// Remove trailing comma and space
	if len(chartData) > 0 {
		chartData = chartData[:len(chartData)-2]
	}

	return fmt.Sprintf("[%s]", chartData), nil
}

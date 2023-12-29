package chart

import (
	"fmt"
	"time"
)

// Provider defines the interface for a chart data provider
type Provider interface {
	GetChartData(currency string, startTime, endTime time.Time) (string, error)
}

// BinanceProvider is an implementation of the chart data provider using the Binance API
type BinanceProvider struct {
	BinanceClient *BinanceClient
}

// GetChartData returns real chart data from the Binance API
func (b *BinanceProvider) GetChartData(currency string, startTime, endTime time.Time) (string, error) {
	// Specify the trading pair symbol on Binance (e.g., "BTCUSDT")
	symbol := fmt.Sprintf("%sUSDT", currency)

	// Fetch real chart data using the Binance API client
	chartData, err := b.BinanceClient.GetChartData(symbol, startTime, endTime)
	if err != nil {
		return "", fmt.Errorf("failed to fetch real chart data: %v", err)
	}

	return chartData, nil
}

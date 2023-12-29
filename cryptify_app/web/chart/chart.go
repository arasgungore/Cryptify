package chart

import "time"

// Provider defines the interface for a chart data provider
type Provider interface {
	GetChartData(currency string, startTime, endTime time.Time) string
}

// MockProvider is a mock implementation of the chart data provider
type MockProvider struct{}

// GetChartData returns mock chart data
func (m *MockProvider) GetChartData(currency string, startTime, endTime time.Time) string {
	// Implement actual chart data retrieval here
	return "Mock chart data"
}

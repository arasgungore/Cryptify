package chart

import (
	"fmt"
	"math/rand"
	"time"
)

// MockProvider is a mock implementation of the chart data provider
type MockProvider struct{}

// GetChartData returns mock chart data
func (m *MockProvider) GetChartData(currency string, startTime, endTime time.Time) string {
	// Mock data generation
	dataPoints := generateMockDataPoints(startTime, endTime)

	// Format the chart data
	chartData := ""
	for _, point := range dataPoints {
		chartData += fmt.Sprintf("[%s, %.2f], ", point.Timestamp.Format("2006-01-02T15:04:05"), point.Price)
	}

	// Remove trailing comma and space
	if len(chartData) > 0 {
		chartData = chartData[:len(chartData)-2]
	}

	return fmt.Sprintf("[%s]", chartData)
}

// mockDataPoint represents a data point in the chart
type mockDataPoint struct {
	Timestamp time.Time
	Price     float64
}

// generateMockDataPoints generates mock data points within the specified time range
func generateMockDataPoints(startTime, endTime time.Time) []mockDataPoint {
	// Seed for reproducibility
	rand.Seed(time.Now().UnixNano())

	// Generate random data points
	var dataPoints []mockDataPoint
	currentTime := startTime
	for currentTime.Before(endTime) {
		price := rand.Float64() * 1000 // Random price between 0 and 1000
		dataPoints = append(dataPoints, mockDataPoint{Timestamp: currentTime, Price: price})

		// Increment time by a random duration
		currentTime = currentTime.Add(time.Duration(rand.Intn(3600)) * time.Second) // Random duration between 0 and 3600 seconds
	}

	return dataPoints
}

package exchange

import (
	"fmt"
	"net/http"
	"time"

	"github.com/arasgungore/Cryptify/cryptify_app/internal/trading"
	"github.com/arasgungore/Cryptify/cryptify_app/internal/wallet"
	"github.com/arasgungore/Cryptify/cryptify_app/web/chart"
)

// ExchangeHandler handles HTTP requests related to the cryptocurrency exchange
type ExchangeHandler struct {
	tradingEngine *trading.TradingEngine
	userWallet    *wallet.Wallet
	chartProvider chart.Provider
}

// NewExchangeHandler creates a new instance of ExchangeHandler
func NewExchangeHandler(tradingEngine *trading.TradingEngine, userWallet *wallet.Wallet) *ExchangeHandler {
	return &ExchangeHandler{
		tradingEngine: tradingEngine,
		userWallet:    userWallet,
		chartProvider: &chart.MockProvider{}, // Use a chart provider implementation
	}
}

// StartServer starts the web server
func (e *ExchangeHandler) StartServer() {
	http.HandleFunc("/buy", e.buyHandler)
	http.HandleFunc("/sell", e.sellHandler)
	http.HandleFunc("/chart", e.chartHandler)

	go func() {
		http.ListenAndServe(":8080", nil)
	}()
}

// buyHandler handles the buy request
func (e *ExchangeHandler) buyHandler(w http.ResponseWriter, r *http.Request) {
	// Handle buy request
}

// sellHandler handles the sell request
func (e *ExchangeHandler) sellHandler(w http.ResponseWriter, r *http.Request) {
	// Handle sell request
}

// chartHandler handles the chart request
func (e *ExchangeHandler) chartHandler(w http.ResponseWriter, r *http.Request) {
	currency := r.URL.Query().Get("currency")
	if currency == "" {
		http.Error(w, "Currency parameter is required", http.StatusBadRequest)
		return
	}

	data := e.chartProvider.GetChartData(currency, time.Now().AddDate(0, 0, -7), time.Now())
	fmt.Fprintf(w, "Chart data for %s:\n%s", currency, data)
}

// DisplayChart displays the chart for a specific cryptocurrency
func (e *ExchangeHandler) DisplayChart(currency string) {
	url := fmt.Sprintf("http://localhost:8080/chart?currency=%s", currency)
	fmt.Printf("Opening chart for %s: %s\n", currency, url)
}

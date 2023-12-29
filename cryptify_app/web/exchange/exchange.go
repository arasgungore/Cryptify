package exchange

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/arasgungore/Cryptify/cryptify_app/internal/order"
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
	userID := 1 // Replace with your authentication logic
	quantityStr := r.URL.Query().Get("quantity")
	priceStr := r.URL.Query().Get("price")

	if quantityStr == "" || priceStr == "" {
		http.Error(w, "Quantity and price parameters are required", http.StatusBadRequest)
		return
	}

	quantity, err := strconv.ParseFloat(quantityStr, 64)
	if err != nil {
		http.Error(w, "Invalid quantity parameter", http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Invalid price parameter", http.StatusBadRequest)
		return
	}

	baseCurrency := "BTC"  // Replace with your logic to determine the base currency
	quoteCurrency := "USD" // Replace with your logic to determine the quote currency
	orderType := order.Buy // This is a buy order
	orderID := e.tradingEngine.ProcessOrder(userID, baseCurrency, quoteCurrency, orderType, quantity, price)

	// Update user's wallet
	totalCost := quantity * price
	e.userWallet.SubtractFromBalance(quoteCurrency, totalCost)
	e.userWallet.AddToBalance(baseCurrency, quantity)

	// Display order details
	fmt.Fprintf(w, "Buy order placed successfully! Order ID: %d\n", orderID)
}

// sellHandler handles the sell request
func (e *ExchangeHandler) sellHandler(w http.ResponseWriter, r *http.Request) {
	userID := 1 // Replace with your authentication logic
	quantityStr := r.URL.Query().Get("quantity")
	priceStr := r.URL.Query().Get("price")

	if quantityStr == "" || priceStr == "" {
		http.Error(w, "Quantity and price parameters are required", http.StatusBadRequest)
		return
	}

	quantity, err := strconv.ParseFloat(quantityStr, 64)
	if err != nil {
		http.Error(w, "Invalid quantity parameter", http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Invalid price parameter", http.StatusBadRequest)
		return
	}

	baseCurrency := "BTC"   // Replace with your logic to determine the base currency
	quoteCurrency := "USD"  // Replace with your logic to determine the quote currency
	orderType := order.Sell // This is a sell order
	orderID := e.tradingEngine.ProcessOrder(userID, baseCurrency, quoteCurrency, orderType, quantity, price)

	// Update user's wallet
	e.userWallet.SubtractFromBalance(baseCurrency, quantity)
	e.userWallet.AddToBalance(quoteCurrency, quantity*price)

	// Display order details
	fmt.Fprintf(w, "Sell order placed successfully! Order ID: %d\n", orderID)
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

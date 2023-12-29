package main

import (
	"fmt"
	"log"

	"github.com/arasgungore/Cryptify/cryptify_app/internal/authentication"
	"github.com/arasgungore/Cryptify/cryptify_app/internal/order"
	"github.com/arasgungore/Cryptify/cryptify_app/internal/trading"
	"github.com/arasgungore/Cryptify/cryptify_app/internal/wallet"
	"github.com/arasgungore/Cryptify/cryptify_app/web/exchange"
)

func main() {
	fmt.Println("Crypto Exchange App")

	// Initialize components
	authService := authentication.NewAuthService()
	orderService := order.NewOrderService()
	tradingEngine := trading.NewTradingEngine(orderService)
	userWallet := wallet.NewWallet()

	// Initialize web components
	exchangeHandler := exchange.NewExchangeHandler(tradingEngine, userWallet)
	exchangeHandler.StartServer()

	// Example usage
	userID := 1
	if authService.AuthenticateUser(userID) {
		orderID := orderService.PlaceOrder(userID, "BTC", "ETH", order.Buy, 1, 0.5)
		tradingEngine.ProcessOrder(orderID)
		exchangeHandler.DisplayChart("BTC")
	} else {
		log.Fatal("Authentication failed")
	}
}

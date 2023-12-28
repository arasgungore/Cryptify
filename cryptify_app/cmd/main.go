package main

import (
	"fmt"
	"log"

	"github.com/arasgungore/cryptify_app/internal/authentication"
	"github.com/arasgungore/cryptify_app/internal/order"
	"github.com/arasgungore/cryptify_app/internal/trading"
)

func main() {
	fmt.Println("Cryptify App")

	// Initialize components
	authService := authentication.NewAuthService()
	orderService := order.NewOrderService()
	tradingEngine := trading.NewTradingEngine(orderService)

	// Example usage
	userID := 1
	if authService.AuthenticateUser(userID) {
		orderID := orderService.PlaceOrder(userID, "BTC", "ETH", order.Buy, 1, 0.5)
		tradingEngine.ProcessOrder(orderID)
	} else {
		log.Fatal("Authentication failed")
	}
}

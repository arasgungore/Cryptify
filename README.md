# Cryptify

A cryptocurrency exchange app written in Go.



## Project Structure

- `cmd/`: Main application entry point
  - `main.go`: Main application file

- `internal/`: Internal packages
  - `authentication/`: User authentication package
    - `authentication.go`: User authentication logic
  - `order/`: Order package
    - `order.go`: Order and order service logic
  - `trading/`: Trading package
    - `trading.go`: Trading engine logic
  - `wallet/`: Wallet package
    - `wallet.go`: User wallet logic

- `pkg/`: External packages
  - `database/`: Simulated in-memory database
    - `database.go`: Database logic

- `web/`: Web-related packages
  - `exchange/`: Exchange-related web handlers
    - `exchange.go`: HTTP handlers for buy, sell, and chart requests
  - `chart/`: Chart-related functionality
    - `chart.go`: Chart data provider interface
    - `mock_provider.go`: Mock implementation of the chart data provider



## Usage

1. Navigate to the project directory: `cd cryptify_app`
2. Run the application: `go run cmd/main.go`
3. Open your web browser and visit:
   - Buy: [http://localhost:8080/buy](http://localhost:8080/buy)
   - Sell: [http://localhost:8080/sell](http://localhost:8080/sell)
   - Chart for BTC: [http://localhost:8080/chart?currency=BTC](http://localhost:8080/chart?currency=BTC)



## Notes

This example includes additional features such as a wallet, chart tracking, and basic web functionality. However, it is still a simplified version and may not cover all aspects of a real-world cryptocurrency exchange.



## Contributing

Feel free to contribute to the project by opening issues or pull requests.



## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.



## Author

👤 **Aras Güngöre**

- LinkedIn: [@arasgungore](https://www.linkedin.com/in/arasgungore)
- GitHub: [@arasgungore](https://github.com/arasgungore)

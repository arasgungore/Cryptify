# Cryptify

The Cryptify App is a cryptocurrency exchange application that allows users to register, authenticate, trade cryptocurrencies, and visualize real-time chart data. The application is built with Go and incorporates basic functionalities to provide a starting point for further development.



## Features

- **Authentication:**
  - User registration with email verification.
  - User login with email and password.
  - Secure password hashing.

- **Trading:**
  - Basic trading engine for executing buy and sell orders.
  - Integration with a cryptocurrency exchange (Binance API).

- **Wallet:**
  - User wallet for managing cryptocurrency balances.
  - Transaction history tracking.

- **Chart Visualization:**
  - Real-time chart data using the Binance API.
  - Mock chart data for testing purposes.



## Getting Started

1. **Navigate to the Project Directory:**
   ```bash
   cd cryptify_app
   ```

2. **Create a `.env` File:**
   Create a `.env` file at the root of the project and add your Binance API key and secret (if using real chart data).
   ```env
   BINANCE_API_KEY=your_binance_api_key
   BINANCE_API_SECRET=your_binance_api_secret
   ```

3. **Install Dependencies:**
   ```bash
   go mod tidy
   ```

4. **Run the Application:**
   ```bash
   go run main.go
   ```

5. **Access the App:**
   Open your web browser and go to [http://localhost:8080](http://localhost:8080) to access the application.



## Contributing

If you find any issues or have suggestions for improvement, feel free to open an issue or submit a pull request. Contributions are welcome!



## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.



## Author

👤 **Aras Güngöre**

- LinkedIn: [@arasgungore](https://www.linkedin.com/in/arasgungore)
- GitHub: [@arasgungore](https://github.com/arasgungore)

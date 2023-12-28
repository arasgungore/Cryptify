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

- `pkg/`: External packages
  - `database/`: Simulated in-memory database
    - `database.go`: Database logic



## Usage

1. Navigate to the project directory: `cd cryptify_app`
2. Run the application: `go run cmd/main.go`



## Notes

This is a simplified example for educational purposes only. Building a production-ready cryptocurrency exchange involves various complexities and considerations, including security, compliance, and performance.



## Contributing

Feel free to contribute to the project by opening issues or pull requests.



## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.



## Author

👤 **Aras Güngöre**

- LinkedIn: [@arasgungore](https://www.linkedin.com/in/arasgungore)
- GitHub: [@arasgungore](https://github.com/arasgungore)

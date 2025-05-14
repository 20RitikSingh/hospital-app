# Hospital App

A modern hospital management system built with Go. This application provides APIs for managing patients, users, authentication, and medical records, following best practices for maintainability and scalability.

## Features
- User authentication and role-based access control
- Patient management (CRUD)
- Medical record management
- RESTful API design
- OpenAPI (Swagger) documentation
- Modular project structure

## Project Structure
```
├── bin/                  # Compiled binaries
├── cmd/hospital-app/     # Main application entrypoint
├── internal/             # Application core logic
│   ├── api/              # OpenAPI specs
│   ├── config/           # Configuration
│   ├── db/               # Database migration and seeding
│   ├── handlers/         # HTTP handlers
│   ├── mapper/           # Data mappers
│   ├── middleware/       # Middleware (auth, roles)
│   ├── openapi/          # Generated OpenAPI code
│   ├── repository/       # Data repositories
│   ├── routes/           # Route definitions
│   ├── service/          # Business logic
│   └── utils/            # Utility functions
├── models/               # Data models
├── makefile              # Build, test, and clean commands
├── go.mod, go.sum        # Go modules
```

## Getting Started

### Prerequisites
- Go 1.22 or later
- Make

### Build
```sh
make build
```

### Run
```sh
make run
```

### Test
```sh
make test
```

### Clean
```sh
make clean
```

## API Documentation

The API is documented using OpenAPI (Swagger). You can view and interact with the API documentation here:

[Hospital App API on SwaggerHub](https://app.swaggerhub.com/apis/ddd-518/hospital-app_api/1.0.0)

## GitHub Actions CI

A standard GitHub Actions workflow is included at `.github/workflows/go-ci.yml` to automatically build and test the application on every push and pull request to the `main` branch.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/YourFeature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/YourFeature`)
5. Open a pull request

## License

This project is licensed under the MIT License.

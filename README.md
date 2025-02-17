# Go Web Application Template

This is a template for a Go web application following the MVC (Model-View-Controller) pattern.

## Project Structure

```
.
├── controllers/         # HTTP request handlers
├── models/             # Data models and database interactions
├── views/              # HTML templates
├── static/             # Static files (CSS, JS, images)
├── main.go             # Application entry point
├── Dockerfile          # Docker configuration
├── docker-compose.yml  # Docker Compose configuration
└── README.md          # Project documentation
```

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- PostgreSQL (if running locally)

## Getting Started

### Using Docker (Recommended)

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <project-directory>
   ```

2. Build and run the application using Docker Compose:
   ```bash
   docker-compose up --build
   ```

3. Access the application at `http://localhost:8080`

### Local Development

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Set up the PostgreSQL database:
   - Create a database named `webapp`
   - Update the database connection string in `models/db.go` if needed

3. Run the application:
   ```bash
   go run main.go
   ```

4. Access the application at `http://localhost:8080`

## Database Migrations

The application uses PostgreSQL as its database. You'll need to create the following tables:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
```

## Adding New Features

### Controllers
Add new controllers in the `controllers` directory. Each controller should handle a specific resource or feature.

### Models
Add new models in the `models` directory. Models should represent your data structures and contain business logic.

### Views
Add new templates in the `views` directory. Templates use Go's html/template package.

## Testing

Run tests using:
```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

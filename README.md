# Snippetbox

A secure web application built in Go for creating, storing, and sharing text snippets with expiration dates.

## Overview

Snippetbox is a full-featured web application that allows users to paste and share text snippets (similar to Pastebin or GitHub Gist). The application demonstrates professional Go web development practices including authentication, database integration, secure session management, and HTTPS encryption.

## Features

### Snippet Management
- Create text snippets with title, content, and expiration time
- Set snippet expiration: 1 day, 7 days, or 365 days
- View individual snippets by ID
- Browse the latest snippets on the home page
- Automatic snippet expiration handling

### User Authentication
- User registration with email validation
- Secure login/logout functionality
- Password hashing with bcrypt
- Session-based authentication
- Protected routes requiring authentication

### Security
- HTTPS/TLS encryption enabled by default
- CSRF protection on all forms
- Secure session cookies with 12-hour expiration
- Modern TLS cipher suite configuration
- SQL injection protection via prepared statements
- XSS protection through template escaping

### User Experience
- Flash messages for user feedback
- Form validation with detailed error messages
- Clean HTML templates with partials and layouts
- Responsive static asset serving

## Technology Stack

- **Language**: Go 1.25.5
- **Database**: MySQL
- **Router**: Pat (github.com/bmizerany/pat)
- **Session Management**: golangcollege/sessions
- **Middleware**: Alice (github.com/justinas/alice)
- **CSRF Protection**: nosurf (github.com/justinas/nosurf)
- **Password Hashing**: golang.org/x/crypto/bcrypt

## Project Structure

```
snippetbox/
├── cmd/
│   └── web/              # Application entry point and HTTP handlers
│       ├── main.go       # Server configuration and startup
│       ├── handlers.go   # HTTP request handlers
│       ├── routes.go     # Route definitions
│       ├── middleware.go # Custom middleware
│       ├── helpers.go    # Helper functions
│       └── templates.go  # Template rendering logic
├── pkg/
│   ├── forms/           # Form validation package
│   └── models/          # Data models and database logic
│       └── mysql/       # MySQL-specific implementations
├── ui/
│   ├── html/            # HTML templates
│   └── static/          # Static assets (CSS, JS, images)
├── tls/                 # TLS certificates
│   ├── cert.pem
│   └── key.pem
├── go.mod
└── go.sum
```

## Prerequisites

- Go 1.25 or higher
- MySQL 5.7 or higher
- OpenSSL (for generating TLS certificates)

## Database Setup

1. Create a MySQL database:

```sql
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. Create the snippets table:

```sql
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);
```

3. Create the users table:

```sql
CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
```

4. Create a database user:

```sql
CREATE USER 'web'@'localhost' IDENTIFIED BY 'pass';
GRANT SELECT, INSERT, UPDATE ON snippetbox.* TO 'web'@'localhost';
```

## Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd snippetbox
```

2. Install dependencies:

```bash
go mod download
```

3. Generate TLS certificates (for development):

```bash
cd tls
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

## Running the Application

### Basic Usage

```bash
go run ./cmd/web
```

The server will start on `https://localhost:4000`

### Configuration Options

```bash
go run ./cmd/web -addr=":443" -dsn="web:pass@/snippetbox?parseTime=true" -secret="your-secret-key"
```

#### Command-line Flags

- `-addr`: HTTP network address (default: ":4000")
- `-dsn`: MySQL data source name (default: "web:pass@/snippetbox?parseTime=true")
- `-secret`: Secret key for session encryption (default: "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge")

## Usage

### Creating Snippets

1. Navigate to `https://localhost:4000/`
2. Sign up for an account or log in
3. Click "Create Snippet"
4. Fill in the title, content, and select an expiration time
5. Submit the form

### Viewing Snippets

- Home page displays the 10 most recent snippets
- Click any snippet title to view its full content
- Expired snippets are automatically excluded from listings

### User Management

- Sign up: `/user/signup`
- Log in: `/user/login`
- Log out: `/user/logout`

## Security Considerations

### Production Deployment

Before deploying to production:

1. Generate strong TLS certificates from a trusted CA
2. Change the default session secret to a random 32-byte string
3. Use environment variables for sensitive configuration
4. Update database credentials
5. Enable appropriate firewall rules
6. Configure proper logging and monitoring
7. Review and adjust timeout settings
8. Consider using a reverse proxy (nginx, Caddy)

### Session Secret

Generate a secure random secret:

```bash
openssl rand -base64 32
```

## Development

### Running Tests

```bash
go test ./...
```

### Code Organization

- **cmd/web**: Application entry point and HTTP layer
- **pkg/models**: Data models and database operations
- **pkg/forms**: Form validation logic
- **ui/html**: HTML templates (pages, layouts, partials)
- **ui/static**: CSS, JavaScript, and images

## Error Handling

The application includes comprehensive error handling:

- Client errors (4xx): Handled with appropriate HTTP status codes
- Server errors (5xx): Logged with stack traces and shown as generic errors
- Database errors: Properly propagated and handled
- Form validation errors: Displayed inline with field-specific messages

## Logging

- Info logs: Sent to stdout
- Error logs: Sent to stderr with timestamps and file locations
- Request logging: All HTTP requests are logged via middleware

## License

This project is for educational purposes.

## Contributing

Contributions are welcome. Please ensure:

1. Code follows Go best practices and idioms
2. All tests pass
3. New features include appropriate tests
4. Documentation is updated accordingly

## Acknowledgments

Built following best practices for Go web development, including patterns from the Go community and standard library documentation.
Let's GO! Learn to Build Web Applications with Go By Alex Edwards
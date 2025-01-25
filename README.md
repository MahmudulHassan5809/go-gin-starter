# Go Gin Starter: Dockerized Go Application with Redis and PostgreSQL

Welcome to your Dockerized Go application! This project is built with a robust architecture, leveraging Docker, Redis, PostgreSQL, and clean code practices to create a scalable and maintainable application.

## Table of Contents
1. [Overview](#overview)
2. [Features](#features)
3. [Architecture](#architecture)
4. [Prerequisites](#prerequisites)
5. [Setup](#setup)
6. [Usage](#usage)
7. [Project Structure](#project-structure)
8. [Commands](#commands)
9. [Contributing](#contributing)
10. [License](#license)

---

## Overview
This project serves as a boilerplate for building web applications in Go, complete with user authentication, caching, and database interactions. It is fully containerized using Docker and supports seamless integration with PostgreSQL and Redis.

## Features
- **Authentication**: JWT-based user authentication.
- **Rate Limiting**: Prevent abuse with customizable rate limiting middleware.
- **Caching**: Redis-based caching with a modular cache manager.
- **Database Integration**: PostgreSQL for structured data storage.
- **Migration Management**: Database migrations using SQL scripts.
- **Clean Architecture**: Modular and maintainable codebase.
- **Error Handling**: Centralized and customizable error handling.
- **Pagination**: Helper functions for managing paginated responses.
- **Dockerized**: Fully containerized with Docker Compose.

## Architecture
This application follows clean architecture principles, dividing the codebase into logical layers for better maintainability and testability:
- **Core**: Contains foundational modules like caching, error handling, database connectors, and middlewares.
- **Modules**: Encapsulates application features (e.g., authentication, health checks, user management).
- **Routes**: Handles HTTP routing.

## Prerequisites
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Setup
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <repository-name>
   ```

2. Copy the `.env.example` file to `.env`:
   ```bash
   cp .env.example .env
   ```

3. Build and start the application using Docker Compose:
   ```bash
   docker-compose up --build
   ```

4. Access the application:
   - App: [http://localhost:8080](http://localhost:8080)
   - PostgreSQL: `localhost:5432`
   - Redis: `localhost:6379`

## Usage
### Database Migrations
Run migrations using the `docker.Makefile`:

- **Migrate up**:
  ```bash
  make -f docker.Makefile migrate-up
  ```

- **Migrate down**:
  ```bash
  make -f docker.Makefile migrate-down
  ```

### API Endpoints
- **Authentication**:
  - `POST /auth/login`: Login a user.
  - `POST /auth/register`: Register a new user.
- **Health Check**:
  - `GET /health`: Check the application's health.
- **User Management**:
  - `GET /users`: Fetch a list of users.

## Project Structure
```plaintext
├── docker-compose.yml
├── Dockerfile
├── docker.Makefile       # Docker-specific commands
├── Makefile              # General project commands
├── migrations            # Database migration scripts
├── src
│   ├── core              # Core modules (e.g., caching, db, errors)
│   ├── main.go           # Entry point of the application
│   ├── modules           # Features (e.g., auth, health, users)
│   └── routes            # HTTP route definitions
```

## Commands
- **Create Database**:
  ```bash
  make -f docker.Makefile create-db
  ```

- **Drop Database**:
  ```bash
  make -f docker.Makefile drop-db
  ```

- **Run Tests**:
  ```bash
  make test
  ```

## Contributing
1. Fork the repository.
2. Create a new branch for your feature or bugfix:
   ```bash
   git checkout -b feature/awesome-feature
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add awesome feature"
   ```
4. Push to your branch:
   ```bash
   git push origin feature/awesome-feature
   ```
5. Create a pull request.

## License
This project is licensed under the [MIT License](LICENSE).


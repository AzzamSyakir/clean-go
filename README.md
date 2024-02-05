
# clean-go Project 

Welcome to the **clean-go Project !** This project provides a structured and comprehensive template for testing **APIs**, covering basic **CRUD operations** for user management, including features like login and register.

## Table of Contents

-[Architecture](#Architecture)
- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
- [Running Tests](#running-tests)
- [Project Structure](#project-structure)
  - [Explanation of Project Structure](#explanation-of-project-structure)
- [Dependencies](#dependencies)
- [Contributing](#contributing)

## Clean Architecture
![Clean Architecture](https://github.com/AzzamSyakir/clean-go/blob/main/architecture.png)

### Workflow:
1. External system initiates a request (HTTP, gRPC, Messaging, etc).
2. The Delivery layer creates various Models from request data.
3. The Delivery layer calls the Use Case, executing it using Model data.
4. The Use Case creates Entity data for business logic.
5. The Use Case calls the Repository, executing it using Entity data.
6. The Repository uses Entity data to perform database operations.
7. The Repository performs database operations on the database.
8. The Use Case creates various Models for the Gateway or from Entity data.
9. The Use Case calls the Gateway, executing it using Model data.
10. The Gateway uses Model data to construct requests to an external system.
11. The Gateway performs requests to external systems (HTTP, gRPC, Messaging, etc).

This architecture promotes separation of concerns and enhances maintainability, scalability, and testability. Each layer has a specific responsibility, contributing to a well-organized and comprehensible codebase.



## Introduction

This Golang project template is designed to showcase best practices for testing in Go applications. It focuses on creating a simple API for user management, including registration, login, and basic CRUD operations. The goal is to provide a clean and well-organized foundation for building scalable and maintainable applications.


# Features

- **User Management**: Implement user registration, login, and basic CRUD operations.
- **API Testing**: Comprehensive testing suite covering API endpoints.
- **Environment Configuration**: Utilize environment variables for configuration.
- **Database Interaction**: Interact with a MySQL database for user data storage.
- **Structured Logging**: Employ structured logging for better traceability.
- **Dependency Management**: Use Go modules for efficient dependency management.
- **Consistent Coding Style**: Follow a consistent coding style for better code readability.
- **Documentation**: Well-documented code and a README for easy understanding.

# Getting Started

To get started with this project, follow these steps:

1. Clone the repository:
``` git
git clone this repo link
cd your-repo
```
2. Set up your environment variables:

Create a **.env** file based on **.env.example** and fill in the required configuration.

3. Install dependencies:

```go
go mod download
```
4. change your database configuration in **db.go** and setup tabels in package **migrate**

5. run migrate to migrate your tabel to database 
 ```make
 make migrate
 ```
 6. run projects
 ```make
 make or make run
 ```
# Running Tests

To run tests, run the following command

```bash
  make test
```

# Project Structure

The project structure is designed for clarity and maintainability, following a clean architecture approach:

```plaintext
/
├── go.mod
├── makefile
├── .env
├── cmd/
│   └── main.go
├── api/
│   ├── controller/
│   │   └── user_handlers.go (handles HTTP requests)
│   ├── service/
│   │   └── user_service.go (handles business logic)
│   ├── repositories/
│   │   └── user_repository.go (manages database interactions)
│   ├── entity/
│   │   └── user.go (declares the user entity)
│   ├── middleware/
│   │   └── middleware.go (houses middleware code)
│   └── routes/
│       └── routes.go (initializes routes and runs the server)
├── migration/
│   ├── User.go (initializes user table migration)
│   ├── migrate.go (sets up migration file)
│   └── token.go (initializes token table migration)
└── config/
    └── db.go (initializes database connection)
```



## explanation of project structure:

- **httpLayer:** Handles HTTP requests and serves as the entry point for external communication.

- **usecaseLayer:** Contains business logic and orchestrates data flow between the httpand repository layers.

- **Repository Layer:** Manages interactions with the database, providing an abstraction layer to the service.

- **Entity Layer:** Declares data structures representing entity in the application, such as database tables.

- **Middleware Layer:** Houses middleware code for common functionalities.

- **Routes Layer:** Initializes routes for the server, connecting controllers to specific HTTP endpoints.

- **Migration Layer:** Manages database migration files for initializing tables.

- **Config Layer:** Initializes and manages the database connection.

This architecture enhances maintainability by separating concerns and facilitating easier testing and scalability. Each layer has a distinct responsibility, contributing to a well-organized and comprehensible codebase.
# dependencies
The project uses Go modules for dependency management. To install dependencies, run:

```bash
go mod download
```

# Contributing

Contributions are welcome! To contribute to the project, follow these steps:

1. Fork the repository.
2. Create a new branch: `git checkout -b feature-name`.
3. Make your changes and commit them: `git commit -m 'Add new feature'`.
4. Push to the branch: `git push origin feature-name`.
5. Submit a pull request.

Please ensure your code follows the project's coding style and includes relevant tests. Your contributions will be reviewed, and once approved, they will be merged into the main branch.

Thank you for contributing to the Clean Golang Template!

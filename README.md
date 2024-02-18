
# clean-go Project 

Welcome to the **clean-go Project !** This project provides a structured and comprehensive template for testing **APIs**, covering basic **CRUD operations** for user management, including features like login and register.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
- [Running Tests](#running-tests)
- [Project Structure](#project-structure)
  - [Explanation of Project Structure](#Project-Structure-Explanation)
- [Contributing](#contributing)


# Clean Architecture
![Clean Architecture](https://github.com/AzzamSyakir/clean-go/blob/main/architecture.png)


## Workflow:
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


# Introduction

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

-   
# Getting Started

To get started with this project, follow these steps:

## 1.  Clone the repository:
``` git
git clone this repo link
cd your-repo
```
##  2. run docker compose :
```bash
docker compose up --build -d
```
## 3.  Set up your environment variables:

   
Create a **.env** file based on **.env.example** and fill in the required configuration.



## 4. Install dependencies:

```go
go mod download
```
## 5.  change your database configuration

in **db.go** and setup tabels in package **migrate**



## 8. run migrate to migrate your tabel to database 
 ```make
 make migrate
 ```
## 7.  run projects
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
clean-go/
├── cmd/
│   └── main.go
├── api/
│   └── api-spec.json (Postman API collection)
├── internal/
│   ├── delivery/
│   │   ├── http/
│   │   │   ├── middleware/
│   │   │   │   └── middleware.go (code for middleware)
│   │   │   ├── route/
│   │   │   │   └── route.go (initialize routes for server and run server)
│   │   │   └── user_controller.go (layer httphandling HTTP requests)
│   │   └── messaging/
│   ├── entity/
│   │   └── user_entity.go (declaration of user entity struct)
│   ├── usecase/
│   │   └── user_usecase.go (layer usecase handling business logic)
│   ├── repository/
│   │   └── user_repository.go (layer repositories handling HTTP interactions to the database)
│   └── config/
│       └── db.go (initialize database connection)
|   └── gateway/messaging/
│       └── db.go (initialize database connection)
├── migration/
│   ├── User.go (initialize user table migration)
│   ├── migrate.go (setup migrate file)
│   └── token.go (initialize token table migration)
├── test/
│   └── user_test.go (unit testing here)
├── go.mod
├── makefile
├── .env
└── vendor/
```



# Project Structure Explanation

## Overview

The project follows a clean architecture approach, emphasizing separation of concerns and maintainability. It is organized into distinct layers to facilitate scalability and modularity.

### 1. **cmd/**

- **main.go:** Entry point of the application, initiating the core functionalities.

### 2. **api/**

- **api-spec.json:** Postman API collection providing documentation and examples for API endpoints.

### 3. **internal/**

The `internal` directory houses the core components of the application.

#### a. **delivery/**

- **http/:**
  - **middleware/:** Contains middleware.go, housing code for handling middleware operations.
  - **route/:** Manages route.go, which initializes routes for the server and runs the server.
  - **user_controller.go:** Responsible for handling HTTP requests related to users.

- **messaging/:** Manages messaging components related to delivery.

#### b. **entity/**

- **user_entity.go:** Declares the structure for the user entity, capturing essential attributes.

#### c. **usecase/**

- **user_usecase.go:** Manages the use case layer, handling the business logic related to users.

#### d. **repository/**

- **user_repository.go:** Handles interactions with the database and manages HTTP interactions related to users.

#### e. **config/**

- **db.go:** Initializes the database connection.

#### f. **gateway/**

- **messaging/:** Manages messaging components related to the gateway.

### 4. **migration/**

- **User.go:** Initializes user table migration.
- **migrate.go:** Sets up migration file.
- **token.go:** Initializes token table migration.

### 5. **test/**

- **user_test.go:** Unit tests for user-related functionalities.

### 6. **go.mod, makefile, .env**

- **go.mod:** Specifies the Go modules and their dependencies.
- **makefile:** Automates common tasks and build processes.
- **.env:** Configuration file for environment variables.

### 7. **vendor/**

- Houses external dependencies.

This structure promotes a clear separation of concerns, making the codebase modular and easily maintainable. Each layer focuses on specific responsibilities, contributing to the overall cleanliness and scalability of the architecture.

# Contributing

Contributions are welcome! To contribute to the project, follow these steps:

1. Fork the repository.
2. Create a new branch: `git checkout -b feature-name`.
3. Make your changes and commit them: `git commit -m 'Add new feature'`.
4. Push to the branch: `git push origin feature-name`.
5. Submit a pull request.

Please ensure your code follows the project's coding style and includes relevant tests. Your contributions will be reviewed, and once approved, they will be merged into the main branch.

Thank you for contributing to the Clean Golang Template!

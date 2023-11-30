# Clean Golang
This is a clean architecture template for a Golang project, designed to enhance maintainability, scalability, and separation of concerns. It follows the principles of Clean Architecture, with three main layers: Controller (HTTP handling), Service (business logic), and Repository (database interaction).

## Project Structure
```clean-golang
|-- cmd
|   `-- main.go
|-- application
|   |-- router
|   |   `-- routes.go
|   |-- controller
|   |   `-- user_controller.go
|   |-- service
|   |   `-- user_service.go
|   |-- repository
|   |   `-- user_repository.go
|   |-- responses
|   |   `-- responses.go
|   |-- models
|   |   `-- user.go
|-- config
|   `-- database.go
|-- .env.example
|-- go.mod
|-- go.sum
|-- README.md 
```


## Usage
Clone the repository: ``git clone https://github.com/yourusername/clean-golang.git``
Navigate to the project directory: ``cd clean-golang``
Copy the example environment file: ``cp .env.example .env``
Modify the .env file with your database configuration.
Run the project: ``go run cmd/main.go``
## Overview
cmd/main.go: Entry point of the application. It initializes necessary components and starts the server.
### Application Layer
- application/router/routes.go: Defines the application routes and connects them to the corresponding controllers.

- application/controller/user_controller.go: Handles HTTP requests, validates input, and delegates business logic to the service layer.

- application/service/user_service.go: Contains business logic, interacts with the repository, and orchestrates operations.

- application/repository/user_repository.go: Communicates with the database, performs CRUD operations, and returns data to the service layer.

- application/responses/responses.go: Defines standard response formats for the application.

- application/models/user.go: Defines the data structure for the user entity.

### Config
- config/database.go: Manages the database connection and configuration.
### Environment
- .env.example: Example environment file with placeholders for configuration variables. Copy this file as .env and replace the placeholders with actual values.
- Dependencies
- github.com/gorilla/mux: A powerful HTTP router for building web applications.
- golang.org/x/crypto/bcrypt: Library for handling password hashing.
- github.com/joho/godotenv: Loads environment variables from a file.
- github.com/dgrijalva/jwt-go: JSON Web Token (JWT) library for Go.
## Getting Started
Follow the steps under [Usage](#usage) to set up and run the project. You can then start building your application logic within the controller, service, and repository layers.

Contributing
Feel free to contribute by opening issues or pull requests. Your feedback and improvements are highly appreciated!
## License
This project is licensed under the [MIT License](./LICENSE).


package router

import (
	"clean-go/config"
	"clean-go/internal/controller"
	"clean-go/internal/middleware"
	"clean-go/internal/repositories"
	"clean-go/internal/service"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router(db *sql.DB) *mux.Router {
	// Initialize repositories
	userRepository := repositories.NewUserRepository(db)

	// Initialize services
	userService := service.NewUserService(*userRepository)

	// Initialize controllers
	userController := controller.NewUserController(*userService)

	// Create a new router
	router := mux.NewRouter()

	// Protected routes
	protectedRoutes := router.PathPrefix("/").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)

	// Authentication routes
	router.HandleFunc("/users", userController.CreateUserController).Methods("POST")
	router.HandleFunc("/users/login", userController.LoginUser).Methods("POST")
	router.HandleFunc("/users/logout", userController.LogoutUser).Methods("POST")

	// User routes
	protectedRoutes.HandleFunc("/users", userController.FetchUserController).Methods("GET")
	protectedRoutes.HandleFunc("/users/{id}", userController.GetUserController).Methods("GET")
	protectedRoutes.HandleFunc("/users/{id}", userController.UpdateUserController).Methods("PUT")
	protectedRoutes.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	return router
}

func RunServer() {
	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(fmt.Errorf("error connecting to db: %w", err))
	}
	defer db.Close() // Pastikan koneksi database ditutup

	router := Router(db)

	// Mulai server HTTP dengan router yang telah dikonfigurasi
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

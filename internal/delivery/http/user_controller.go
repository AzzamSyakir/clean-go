package http

import (
	"clean-go/internal/usecase"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	models "clean-go/internal/entity"
	"clean-go/internal/gateway/responses"

	"github.com/gorilla/mux"
)

type UserController struct {
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase) *UserController {
	return &UserController{UseCase: useCase}
}
func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var user struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responses.ErrorResponse(w, "Failed to read user data from the request", http.StatusBadRequest)
		return
	}

	err := c.UseCase.Register(user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to create user: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}

	currentTime := time.Now()

	// Create a user data object to be sent in the response
	userData := struct {
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}{
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	responses.SuccessResponse(w, "Success", userData, http.StatusCreated)
}

func (c *UserController) Fetch(w http.ResponseWriter, r *http.Request) {
	usersData, err := c.UseCase.Fetch()
	if err != nil {
		fmt.Println("tes")
		errorMessage := fmt.Sprintf("Failed to get users: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}
	// Check if users slice is empty
	if len(usersData) == 0 {
		errorMessage := "No user data found"
		responses.ErrorResponse(w, errorMessage, http.StatusNotFound)
		return
	}

	// Create a user data object to be sent in the response
	var responseData []struct {
		ID        string `json:"id"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	for _, user := range usersData {
		userData := struct {
			ID        string `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
		}{
			ID:        user.ID,
			Username:  user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(), // Using String() to get string representation of time.Time
			UpdatedAt: user.UpdatedAt.String(), // Using String() to get string representation of time.Time
		}

		responseData = append(responseData, userData)
	}

	// Return user data as JSON
	w.Header().Set("Content-Type", "internal/json")
	responses.SuccessResponse(w, "Success", responseData, http.StatusOK)
}

func (c *UserController) Get(w http.ResponseWriter, r *http.Request) {
	// Get the id parameter
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		responses.ErrorResponse(w, "ID must be provided", http.StatusBadRequest)
		return
	}

	userData, err := c.UseCase.Get(id)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get user: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}

	// Create a user data object to be sent in the response
	responseData := struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		Username:  userData.Name,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt.String(), // Using String() to get string representation of time.Time
		UpdatedAt: userData.UpdatedAt.String(), // Using String() to get string representation of time.Time
	}

	// Return user data as JSON
	w.Header().Set("Content-Type", "internal/json")
	responses.SuccessResponse(w, "Success", responseData, http.StatusOK)
}
func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	// Get the id parameter
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		responses.ErrorResponse(w, "ID must be provided", http.StatusBadRequest)
		return
	}

	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		responses.ErrorResponse(w, "Failed to read user data from the request", http.StatusBadRequest)
		return
	}

	// Update user in the usecase layer
	_, err := c.UseCase.Update(id, updatedUser)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to update user: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}

	currentTime := time.Now()

	// Create a user data object to be sent in the response
	updatedData := models.User{
		Name:      updatedUser.Name,
		Email:     updatedUser.Email,
		UpdatedAt: currentTime,
	}

	// Return response with only updated data
	w.Header().Set("Content-Type", "internal/json")
	responses.SuccessResponse(w, "Success", updatedData, http.StatusOK)
}
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	// Get the id parameter
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		responses.ErrorResponse(w, "ID must be provided", http.StatusBadRequest)
		return
	}
	// Delete user in the usecase layer
	err := c.UseCase.Delete(id)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to delete user: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}

	responses.OtherResponses(w, "Success delete user", http.StatusOK)
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var requestUser map[string]string

	// Read JSON from request body
	if err := json.NewDecoder(r.Body).Decode(&requestUser); err != nil {
		responses.ErrorResponse(w, "Failed to read user data from request", http.StatusBadRequest)
		return
	}

	// Get email and password from the user data
	email, ok := requestUser["email"]
	if !ok {
		responses.ErrorResponse(w, "Email must be provided", http.StatusBadRequest)
		return
	}

	password, ok := requestUser["password"]
	if !ok {
		responses.ErrorResponse(w, "Password must be provided", http.StatusBadRequest)
		return
	}

	// Call the usecase to perform login
	token, err := c.UseCase.Login(email, password)
	if err != nil {
		errorMessage := fmt.Sprintf("Login failed: %v", err)

		responses.ErrorResponse(w, errorMessage, http.StatusUnauthorized)
		return
	}

	// Add token to cookie
	expiration := time.Now().Add(7 * 24 * time.Hour) // Expiration 7 days
	cookie := http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)

	// Respond with the token
	response := map[string]interface{}{
		"token": token,
	}
	json.NewEncoder(w).Encode(response)
}

func (c *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	// Get the token from the "Authorization" header
	tokenString := r.Header.Get("Authorization")

	// Remove "Bearer " from the token string
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Call the LogoutUser function from the service
	err := c.UseCase.Logout(tokenString) //delete token in redis
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to logout user: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}
	// Delete RefToken from cookie
	expiration := time.Now().AddDate(0, 0, -1) // Set the expiration time to the past
	cookie := http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)

	// Return a success message after logout
	responses.OtherResponses(w, "Logout successful", http.StatusOK)
}

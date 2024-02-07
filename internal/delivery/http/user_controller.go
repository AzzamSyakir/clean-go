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

	// Membuat objek data pengguna untuk dikirim dalam respons
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

	// Membuat objek data pengguna untuk dikirim dalam respons
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
			CreatedAt: user.CreatedAt.String(), // Menggunakan String() untuk mendapatkan representasi string dari time.Time
			UpdatedAt: user.UpdatedAt.String(), // Menggunakan String() untuk mendapatkan representasi string dari time.Time
		}

		responseData = append(responseData, userData)
	}

	// Mengembalikan data pengguna sebagai JSON
	w.Header().Set("Content-Type", "internal/json")
	responses.SuccessResponse(w, "Success", responseData, http.StatusOK)
}

func (c *UserController) Get(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan parameter id
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		responses.ErrorResponse(w, "id harus disertakan", http.StatusBadRequest)
		return
	}

	userData, err := c.UseCase.Get(id)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get user: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}

	// Membuat objek data pengguna untuk dikirim dalam respons
	responseData := struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		Username:  userData.Name,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt.String(), // Menggunakan String() untuk mendapatkan representasi string dari time.Time
		UpdatedAt: userData.UpdatedAt.String(), // Menggunakan String() untuk mendapatkan representasi string dari time.Time
	}

	// Mengembalikan data pengguna sebagai JSON
	w.Header().Set("Content-Type", "internal/json")
	responses.SuccessResponse(w, "Success", responseData, http.StatusOK)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan parameter id
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		responses.ErrorResponse(w, "id harus disertakan", http.StatusBadRequest)
		return
	}

	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		responses.ErrorResponse(w, "Failed to read user data from the request", http.StatusBadRequest)
		return
	}

	// Update user in the usecaselayer
	_, err := c.UseCase.Update(id, updatedUser)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to update user: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}

	currentTime := time.Now()

	// Membuat objek data pengguna untuk dikirim dalam respons
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
	// Mendapatkan parameter id
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		responses.ErrorResponse(w, "id harus disertakan", http.StatusBadRequest)
		return
	}
	// delete user in the usecaselayer
	err := c.UseCase.Delete(id)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to Delete user: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}

	responses.OtherResponses(w, "Success delete user", http.StatusOK)
}
func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var requestUser map[string]string

	// Membaca data JSON dari body permintaan
	if err := json.NewDecoder(r.Body).Decode(&requestUser); err != nil {
		responses.ErrorResponse(w, "Gagal membaca data pengguna dari permintaan", http.StatusBadRequest)
		return
	}

	// Mendapatkan email dan password dari data pengguna
	email, ok := requestUser["email"]
	if !ok {
		responses.ErrorResponse(w, "Email harus diisi", http.StatusBadRequest)
		return
	}

	password, ok := requestUser["password"]
	if !ok {
		responses.ErrorResponse(w, "Password harus diisi", http.StatusBadRequest)
		return
	}

	// Memanggil usecaseuntuk melakukan login
	token, err := c.UseCase.Login(email, password)
	if err != nil {
		errorMessage := fmt.Sprintf("login failed: %v", err)

		responses.ErrorResponse(w, errorMessage, http.StatusUnauthorized)
		return
	}

	// Mengembalikan token dan pesan sukses
	response := map[string]string{"token": token}
	responses.SuccessResponse(w, "Login berhasil", response, http.StatusOK)
}

func (c *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan token dari header "Authorization"
	tokenString := r.Header.Get("Authorization")

	// Membersihkan token dari string "Bearer "
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Memanggil fungsi LogoutUser dari service
	err := c.UseCase.Logout(tokenString)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to logout user: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}

	// Mengembalikan token yang baru setelah logout
	responses.OtherResponses(w, "logout berhasil", http.StatusOK)
}

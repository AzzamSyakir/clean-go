package http

type AuthController struct {
	UseCase *usecase.AuthUseCase
}

func NewAuthController(useCase *usecase.AuthUseCase) *AuthController {
	return &AuthController{UseCase: useCase}
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
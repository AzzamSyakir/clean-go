package usecase

func (c *UserUseCase) Register(id string, name, email, password string) error {
	// Hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Generate a UUID if ID is not provided
	var userID string
	if id == "" {
		uuid := uuid.New()
		userID = uuid.String()
	} else {
		userID = id
	}
	// Current time
	currentTime := time.Now()

	// Save the user to the database using the validated data
	err = c.UserRepository.Register(userID, name, email, string(hashedPassword), currentTime, currentTime)
	if err != nil {
		return err
	}

	return err
}
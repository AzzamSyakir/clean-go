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
func (c *UserUseCase) Login(email string, password string) (string, error) {
	// Mendapatkan informasi pengguna dari repository
	user, err := c.UserRepository.LoginUser(email)
	if err != nil {
		return "", err
	}

	// Membandingkan password yang dimasukkan dengan password yang ada di database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("password salah")
	}

	// Menghasilkan string acak sebagai bagian dari token
	// randomString, err := utils.GenerateRandomString(12)
	if err != nil {
		return "", err
	}

	// Membuat token dari UUID user ditambah string acak
	token := fmt.Sprintf("%s:%s", user.ID, utils.GenerateRandomString(16))

	// Menentukan waktu kadaluarsa acces token
	expirationTime := time.Now().Add(15 * time.Minute) // Token berlaku selama 15 menit

	// Menyiapkan data untuk disimpan di Redis
	redisKey := fmt.Sprintf("tokens:%s", token) // Gunakan token sebagai kunci Redis

	// Struct untuk menyimpan token dan user ID
	type TokenAndUserID struct {
		Token  string `json:"token"`
		UserID string `json:"userId"`
	}

	// Simpan token dan ID pengguna ke dalam cache
	tokenAndUserID := TokenAndUserID{
		Token:  token,
		UserID: user.ID,
	}

	// Mengonversi struct ke dalam bentuk JSON
	data, err := json.Marshal(tokenAndUserID)
	if err != nil {
		return "", err
	}

	// Simpan data ke dalam cache dengan waktu kadaluarsa
	err = cache.SetCached(redisKey, data, expirationTime)
	if err != nil {
		return "", err
	}

	return token, nil
}
func (c *UserUseCase) Logout(tokenString string) error {
	redisKey := fmt.Sprintf("tokens:%s", tokenString) // Gunakan token sebagai kunci Redis
	c.UserRepository.LogoutUser(redisKey)             //delete AccesTokentoken in redis
	return nil
}

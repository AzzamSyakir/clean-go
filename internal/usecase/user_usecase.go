package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing-golang/cache"
	"testing-golang/internal/entity"
	"testing-golang/internal/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

// authentication
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

	// Jika login berhasil, buat token JWT
	token := jwt.New(jwt.SigningMethodHS256)

	// Menentukan klaim (claims) token
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["username"] = user.Name
	// buat exp time
	expirationTime := time.Now().Add(time.Hour)
	claims["exp"] = expirationTime.Unix() // set token dengan exp time yng ditentukan

	// Menandatangani token dengan secret key
	secretKeyString := os.Getenv("SECRET_KEY")
	secretKey := []byte(secretKeyString)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	redisKey := fmt.Sprintf("tokens:%s", user.ID) // Gunakan id token sebagai RedisKey

	//struct token dan user id
	type TokenAndUserID struct {
		Token  string
		UserID string
	}

	// Simpan token dan ID pengguna ke dalam cache
	var tokenAndUserID TokenAndUserID
	tokenAndUserID.Token = tokenString
	tokenAndUserID.UserID = user.ID

	// Mengonversi instance ke dalam bentuk byte menggunakan encoding JSON
	data, err := json.Marshal(tokenAndUserID)
	if err != nil {
		return "", err
	}

	// Simpan data ke dalam cache
	err = cache.SetCached(redisKey, data, expirationTime)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (c *UserUseCase) Logout(tokenString string) error {
	type DataUsers struct {
		Username string `json:"username"`
		UserId   string `json:"user_id"`
		jwt.StandardClaims
	}
	// Parse token dan dapatkan claims
	token, err := jwt.ParseWithClaims(tokenString, &DataUsers{}, func(token *jwt.Token) (interface{}, error) {
		// Verifikasi bahwa metode tanda tangan sesuai
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metode tanda tangan tidak valid: %v", token.Header["alg"])
		}
		// Kembalikan secret key untuk memverifikasi tanda tangan
		secretKeyString := os.Getenv("SECRET_KEY")
		return []byte(secretKeyString), nil
	})
	if err != nil {
		return fmt.Errorf("gagal mengurai token: %v", err)
	}

	// Mengekstrak klaim dari token
	claims, ok := token.Claims.(*DataUsers)
	if !ok {
		return err
	}
	userID := claims.UserId
	currentTime := time.Now()

	c.UserRepository.LogoutUser(userID, currentTime)
	return err
}

// basic user's operations
func (c *UserUseCase) Fetch() ([]entity.User, error) {
	return c.UserRepository.FetchUsers()
}
func (c *UserUseCase) Update(id string, updatedUser entity.User) (entity.User, error) {
	// Business logic/validation goes here

	// Validate if the user exists
	user, err := c.UserRepository.GetUser(id)
	if err != nil {
		return entity.User{}, err
	}
	if user == nil {
		return entity.User{}, errors.New("user not found")
	}

	// Update only non-empty fields
	if updatedUser.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return entity.User{}, err
		}
		user.Password = string(hashedPassword)
	}

	if updatedUser.Name != "" {
		user.Name = updatedUser.Name
	}

	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}

	// Set the updated time
	user.UpdatedAt = time.Now()

	// Call repository to update user in the database
	err = c.UserRepository.UpdateUser(*user)
	if err != nil {
		return entity.User{}, err
	}

	// Return only the updated data
	updatedData := entity.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
	}

	return updatedData, nil
}
func (usecase *UserUseCase) Delete(id string) error {
	// delete token cache
	RedisKey := fmt.Sprintf("tokens:%s", id) // Gunakan token + id user  sebagai RedisKey
	cache.DeleteCached(RedisKey)
	// Call repository to delete user in the database
	err := usecase.UserRepository.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
func (c *UserUseCase) Get(id string) (*entity.User, error) {
	return c.UserRepository.GetUser(id)
}

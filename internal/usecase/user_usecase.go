package usecase

import (
	"clean-go/cache"
	"clean-go/internal/entity"
	"clean-go/internal/repository"
	"clean-go/utils"
	"encoding/json"
	"errors"
	"fmt"
	"time"

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

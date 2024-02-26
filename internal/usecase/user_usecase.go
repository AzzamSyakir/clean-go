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

package storage

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"user-api/models"
)

var (
	mu        sync.Mutex
	users     = make(map[int]models.User)
	currentID int
)

func GetAllUsers() []models.User {
	mu.Lock()
	defer mu.Unlock()
	allUsers := make([]models.User, 0, len(users))
	for _, user := range users {
		allUsers = append(allUsers, user)
	}
	return allUsers
}
func CreateUser(user models.User) (models.User, error) {
	mu.Lock()
	defer mu.Unlock()

	if user.Username == "" {
		return models.User{}, errors.New("Это поле не может быть пустым")
	}
	if user.Email == "" {
		return models.User{}, errors.New("И это поле тоже не может быть пустым")
	}

	currentID++
	user.ID = currentID
	users[user.ID] = user
	user.CreatedAt = time.Now()

	fmt.Printf("Создан новый пользователь %+v\n", user)
	return user, nil
}
func GetUserById(id int) (models.User, error) {
	mu.Lock()
	defer mu.Unlock()
	user, exist := users[id]

	if !exist {
		return models.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

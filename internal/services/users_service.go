package services

import "github.com/Matheus002/go-microservice-rest/internal/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}

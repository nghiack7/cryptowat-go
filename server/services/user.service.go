package services

import "github.com/cryptowat-go/server/models"

type UserService interface {
	FindUserById(string) (*models.User, error)
	FindUserByEmail(string) (*models.User, error)
	UpsertUser(string, *models.UpdateDBUser) (*models.User, error)
}

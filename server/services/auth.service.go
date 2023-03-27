package services

import "github.com/cryptowat-go/server/models"

type AuthService interface {
	SignUpUser(*models.SignUpInput) (*models.User, error)
	SignInUser(*models.SignInInput) (*models.User, error)
}

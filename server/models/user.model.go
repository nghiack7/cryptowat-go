package models

import (
	"github.com/cryptowat-go/server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type SignUpInput struct {
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	PasswordConfirm string    `json:"passwordConfirm"`
	Role            string    `json:"role"`
	Provider        string    `json:"provider,omitempty"`
	Photo           string    `json:"photo,omitempty"`
	Verified        bool      `json:"verified"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID              string    `json:"id", gorm:"primary_key"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	PasswordConfirm string    `json:"passwordConfirm,omitempty"`
	Provider        string    `json:"provider"`
	Photo           string    `json:"photo,omitempty"`
	Role            string    `json:"role"`
	Verified        bool      `json:"verified"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty" `
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateDBUser struct {
	ID              string    `json:"id,omitempty"`
	Name            string    `json:"name,omitempty"`
	Email           string    `json:"email,omitempty"`
	Password        string    `json:"password,omitempty"`
	PasswordConfirm string    `json:"passwordConfirm,omitempty"`
	Role            string    `json:"role,omitempty"`
	Provider        string    `json:"provider"`
	Photo           string    `json:"photo,omitempty"`
	Verified        bool      `json:"verified,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}
type Position struct {
	ID        uint      `gorm: primery_key`
	UserID    string    `json:"userId"`
	Amount    float64   `json:"amount"`
	Currency  string    `json:"currency"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Price struct {
	ID       uint    `json:"id"`
	Currency string  `json:"currency"`
	Photo    string  `json:"photo"`
	Value    float64 `json:"value"`
}

func FilteredResponse(user *User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		Provider:  user.Provider,
		Photo:     user.Photo,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
func OpenDb(config config.Config) *gorm.DB {
	// Connect to Postgres database
	connectionString := config.DBUrl
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&User{}, Position{})
	return db
}

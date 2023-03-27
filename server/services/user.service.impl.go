package services

import (
	"context"
	"encoding/json"
	"github.com/cryptowat-go/server/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	db  *gorm.DB
	ctx context.Context
}

var User *models.User

func NewUserServiceImpl(db *gorm.DB, ctx context.Context) UserService {
	return &UserServiceImpl{db, ctx}
}

// FindUserById Find user by id
func (uc *UserServiceImpl) FindUserById(id string) (*models.User, error) {
	err := uc.db.Where("id=?", id).First(&User).Error
	if err != nil {
		return nil, err
	}
	return User, nil
}

// FindUserByEmail Find user by email
func (uc *UserServiceImpl) FindUserByEmail(email string) (*models.User, error) {
	err := uc.db.Where("email=?", email).First(&User).Error
	if err != nil {
		return nil, err
	}
	return User, nil
}

// UpsertUser Update or insert database
func (uc *UserServiceImpl) UpsertUser(email string, data *models.UpdateDBUser) (*models.User, error) {
	userByte, _ := json.Marshal(data)
	err := json.Unmarshal(userByte, &User)
	if err != nil {
		return nil, err
	}
	err = uc.db.Where("email=?", email).First(&models.User{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			User.ID = uc.generateID()
			err = uc.db.Where("email=?", email).FirstOrCreate(&User).Error
			if err != nil {
			}
			return User, err
		}
		return User, err
	}
	err = uc.db.Save(User).Error
	return User, nil
}

func (uc *UserServiceImpl) generateID() string {
	id := uuid.New().String()
	err := uc.db.Where("id =?", id).First(User).Error
	if err == gorm.ErrRecordNotFound {
		return id
	}
	return uc.generateID()
}

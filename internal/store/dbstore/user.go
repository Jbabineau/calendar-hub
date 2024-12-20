package dbstore

import (
	"github.com/jbabineau/calendar-hub/internal/hash"
	"github.com/jbabineau/calendar-hub/internal/store"
	"gorm.io/gorm"
)

type UserStore struct {
	db           *gorm.DB
	passwordHash hash.PasswordHash
}

type NewUserStoreParams struct {
	DB *gorm.DB
	PasswordHash hash.PasswordHash
}

func NewUserStore(params NewUserStoreParams) *UserStore {
	return &UserStore{
		db: params.DB,
		passwordHash: params.PasswordHash,
	}
}

func (s *UserStore) CreateUser(email string, password string) error {
	hashedPassword, err := s.passwordHash.GeneratePassword(password)
	if err != nil {
		return err
	}

	return s.db.Create(&store.User{
		Email: email, 
		Password: hashedPassword,
	}).Error
}

func (s *UserStore) GetUser(email string) (*store.User, error) {
	var user store.User
	err := s.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, err
}
package repository

import (
	"fmt"
	"domain"
)

type UserRepository interface {
	Insert(userID, name, email string) error
	GeyByUserID(userID string) (*domain.User, error)
}
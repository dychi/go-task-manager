package persistence

import (
	"domain"
	"domain/repository"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

// User Registration
func (up userPersistence) Insert(userID, name, email string) error {
	stmt, err := fmt.Sprintf("INSERT INTO user(user_id, name, email) VALUES(%s, %s, %s)", userID, name, email)
	if err != nil {
		return err
	}
	// _, err = stmt.Exec(userID, name, email)
	fmt.Print(stmt)
	return err
}

// Get UserInfo by UserID
func (up userPersistence) GetByUserID(userID string) (*domain.User, error) {
	row := fmt.Sprintf("SELECT * FROM user WHERE user_id = %s", userID)
	fmt.Print(row)
	// convertToUser
	user := domain.User{}
	return &user, nil
}


// row型をuser型に紐付ける

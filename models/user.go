package model

import (
	"api/database"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Email    string `gorm:"size:255;not null;" json:"email"`
	Entries  []Entry
	gorm.Model
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return user, nil
	}

	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	result := database.Database.Last(&user, "username = ?", username)

	if result.Error == gorm.ErrRecordNotFound {
		return user, nil
	}
	return user, result.Error
}

func FindUserByEmail(email string) (User, error) {
	var user User
	result := database.Database.Last(&user, "email = ?", email)

	if result.Error == gorm.ErrRecordNotFound {
		return user, nil
	}
	return user, result.Error
}

func FindUserById(userId uint) (User, error) {
	var user User
	result := database.Database.Find(&user, "id = ?", userId)

	if result.Error == gorm.ErrRecordNotFound {
		return user, nil
	}
	return user, result.Error
}

package model

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User if you add sensitive fields, don't forget to clean them in setupLogin function.
// Otherwise, the sensitive information will be saved on local storage in plain text!
type User struct {
	Id        int64          `json:"id"`
	CreatedAt int64          `gorm:"type:int;column:created_at;not null"`
	UpdatedAt int64          `gorm:"type:int;column:updated_at;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `json:"username" gorm:"unique;index"`
	Password  string         `json:"password" gorm:"not null;"`
	Token     string         `json:"token" gorm:"not null;"`
}

func Password2Hash(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	return string(hashedPassword), err
}
func ValidatePasswordAndHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u *User) CreateUser() error {
	hashPassword, err := Password2Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashPassword
	result := DB.Save(u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) Save() (*User, error) {
	result := DB.Save(u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

// CheckUser check password
func (u *User) CheckUser() (*User, error) {
	password := u.Password
	username := strings.TrimSpace(u.Username)
	// find buy username or email
	DB.Where("username = ?", username).First(u)
	okay := ValidatePasswordAndHash(password, u.Password)
	if !okay {
		return nil, errors.New("用户名或密码错误！")
	}
	return u, nil
}

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func CheckUserNameExist(username string) (bool, error) {
	var user User
	var err error
	err = DB.Unscoped().First(&user, "username = ?", username).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// not exist, return false, nil
			return false, nil
		}
		// other error, return false, err
		return false, err
	}
	// exist, return true, nil
	return true, nil
}

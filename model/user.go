package model

// Import
import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html"
	"rpi-relay-be-golang/util/token"
	"strings"
)

// User declare object structure
type User struct {
	// Define DB structure for the `users` table
	gorm.Model
	Username 	string `gorm:"size:255;not null;unique" json:"username"`
	Password 	string `gorm:"size:255;not null;" json:"password"`
	FirstName	string `gorm:"size:255;not null;" json:"first_name"`
	LastName	string `gorm:"size:255;not null;" json:"last_name"`
	Email		string `gorm:"size:255;not null;unique" json:"email"`
}

// SaveUser - save data to DB
func (user *User) SaveUser() (*User, error) {
	var err error
	err = DB.Create(&user).Error

	if err != nil {
		return &User{}, err
	}

	return user, nil
}

// BeforeSave pre-processor
func (user *User) BeforeSave() error {

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Remove spaces in username
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return nil

}

// VerifyPassword logic
func VerifyPassword(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// LoginCheck attempt login
func LoginCheck(email string, password string) (string,error) {
	var err error

	user := User{}

	err = DB.Model(User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token,err := token.GenerateToken(user.ID)

	if err != nil {
		return "",err
	}

	return token,nil
}

// GetUserById - return User object by a given ID
func GetUserById(uid uint) (User,error) {
	var user User

	if err := DB.First(&user,uid).Error; err != nil {
		return user,errors.New("User not found!")
	}

	// Prepare the object, which will be returned as JSON
	user.PrepareGive()

	// Return User or Null
	return user,nil
}

// PrepareGive - hide sensitive data from Model
func (user *User) PrepareGive(){
	user.Password = ""
}
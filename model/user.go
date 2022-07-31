package model

// Import
import (
	"github.com/jinzhu/gorm"
)

// User declare object structure
type User struct {
	gorm.Model
	Username 	string `gorm:"size:255;not null;unique" json:"username"`
	Password 	string `gorm:"size:255;not null;" json:"password"`
	FirstName	string `gorm:"size:255;not null;" json:"first_name"`
	LastName	string `gorm:"size:255;not null;" json:"last_name"`
	Email		string `gorm:"size:255;not null;unique" json:"email"`
}
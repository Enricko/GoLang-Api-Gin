package models

import (
	"fmt"
	// "regexp"

	"gorm.io/gorm"
)

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"varchar(300);not null;unique" json:"username" binding:"required,min=6"`
	Password string `gorm:"varchar(300);not null;" json:"password" binding:"required,min=6"`
	gorm.Model
}

func ValidateUser(user *User) error {
    // Validate email format
    // if matched, _ := regexp.MatchString(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`, user.Username); !matched {
    //     return fmt.Errorf("Please provide a valid email address")
    // }
	// Validate name
	if user.Username == "" {
		return fmt.Errorf("Username field is required")
	}

    // Check uniqueness of email
    var existingUser User
    if err := DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil && existingUser.ID != user.ID {
        return fmt.Errorf("Username already exists")
    }

    // Validate level
    // if user.Level != "client" && user.Level != "admin" && user.Level != "owner" {
    //     return fmt.Errorf("Invalid level value")
    // }

    // Validate password length
    if len(user.Password) < 8 || len(user.Password) > 50 {
        return fmt.Errorf("Password must be between 8 and 50 characters")
    }

    return nil
}
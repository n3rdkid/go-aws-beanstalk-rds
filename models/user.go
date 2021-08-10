package models

import "gorm.io/gorm"

// User -> Model for User Table
type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// TableName -> Overriding default table name
func (User) TableName() string {
	return "users"
}

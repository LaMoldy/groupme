package models

import (
    "time"

    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Email       string    `json:"email"`
    Username    string    `json:"username"`
    Password    string    `json:"password"`
    FirstName   string    `json:"firstName"`
    LastName    string    `json:"lastName"`
    DateOfBirth time.Time `json:"dateOfBirth"`
}

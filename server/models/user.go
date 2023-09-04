package models

import (
    "time"
)

type User struct {
    Username    string    `json:"username"`
    Password    string    `json:"password"`
    FirstName   string    `json:"firstName"`
    LastName    string    `json:"lastName"`
    DateOfBirth time.Time `json:"dateOfBirth"`
}

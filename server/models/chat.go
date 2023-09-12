package models

import "gorm.io/gorm"

type Chat struct {
    gorm.Model
    Message string `json:"message"`
    Author  User `json:"user"`
}

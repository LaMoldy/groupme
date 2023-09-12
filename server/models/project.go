package models

import (
    "gorm.io/gorm"
)

type Project struct {
    gorm.Model
    Users       []User    `json:"users"`
    Messages    []Chat    `json:"messages"`
}

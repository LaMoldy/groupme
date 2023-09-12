package models

import "gorm.io/gorm"

type Log struct {
    gorm.Model
    Message string `json:"message"`
    Status  int    `json:"status"`
}

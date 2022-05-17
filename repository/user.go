package repository

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}
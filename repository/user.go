package repository

import "database/sql"

type UserRepository struct {
	db *sql.DB
}

// NewUserRepositoryを初期化
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

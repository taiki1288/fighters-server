package repository

import (
	"context"
	"database/sql"

	"github.com/taiki1288/fighters-server/entity"
)

type User interface {
	CreateUser(ctx context.Context, user *entity.User) error
	FindByUserID(ctx context.Context, userID string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, userID string) error
}

type UserRepository struct {
	db *sql.DB
}

// NewUserRepositoryを初期化
func NewUserRepository(db *sql.DB) User {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	insert := `INSERT INTO users(name, self_introduction, age, like_fighters, created_at, updated_at) VALUES(?, ?, ?, ?, NOW(), NOW())`
	stmt, err := repo.db.Prepare(insert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		user.ID, 
		user.Name, 
		user.SelfIntroduction,
		user.Age,
		user.LikeFighters,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) FindByUserID(ctx context.Context, userID string) (*entity.User, error) {
	query := `SELECT id, name, self_introduction, like_fighters, created_at, updated_at FROM users WHERE id = ?`
    stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(userID)

	user := &entity.User{}

	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.SelfIntroduction,
		&user.Age,
		&user.LikeFighters,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
    const (
        update = `UPDATE users SET name = ?, self_introduction = ?, age = ?, like_fighters = ?, updated_at=NOW() WHERE id = ?`
    )
    result, err := repo.db.Exec(update)
    if err != nil {
        return err
    }

    _, err = result.RowsAffected()
    if err != nil {
        return err
    }
    return nil
} 

func (repo *UserRepository) DeleteUser(ctx context.Context, userID string) error {
	delete := `DELETE FROM users WHERE id = ?`
	stmt, err := repo.db.Prepare(delete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(userID)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
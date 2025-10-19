package repository

import (
	"context"
	"database/sql"
	"errors"

	postgres "github.com/kittichai/core-framework/shared/framework/infrastructure/persistence/adapter/postgres"
	model "github.com/kittichai/core-model/shared/model/infrastructure/persistence/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
}

type UserRepositoryImpl struct {
	db *postgres.PostgresDB
}

func NewUserRepository(db *postgres.PostgresDB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *model.User) error {
	// implement creation using r.db, e.g. r.db.Client.Create(user) or appropriate DB call
	query := `
		INSERT INTO users (id, firstname, lastname, email, username, password_hash, created_at)
		VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.UserName,
		user.PasswordHash,
		user.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *model.User) error {
	// implement update using r.db, e.g. r.db.Client.Update(user) or appropriate DB call
	return nil
}
func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	query := "SELECT id, email, name FROM users WHERE email = $1"
	var schema model.User

	row := r.db.QueryRowContext(ctx, query, email)
	err := row.Scan(&schema.ID, &schema.Email, &schema.FirstName, &schema.LastName, &schema.UserName)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil // ไม่พบผู้ใช้
	}
	if err != nil {
		return nil, err
	}

	// 💡 แปลง Persistence Model เป็น Domain Entity ก่อนส่งออก
	return &schema, nil
}

func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	query := "SELECT id, email, name FROM users WHERE username = $1"
	var schema model.User

	row := r.db.QueryRowContext(ctx, query, username)
	err := row.Scan(&schema.ID, &schema.Email, &schema.FirstName, &schema.LastName, &schema.UserName)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil // ไม่พบผู้ใช้
	}
	if err != nil {
		return nil, err
	}

	// 💡 แปลง Persistence Model เป็น Domain Entity ก่อนส่งออก
	return &schema, nil
}

package userrepository

import (
	"database/sql"

	"github.com/BennoAlif/go-backend-starter/src/entities"
)

type sUserRepository struct {
	DB *sql.DB
}

type UserRepository interface {
	Create(*ParamsCreateUser) (*entities.User, error)
	FindByEmail(*string) (*entities.User, error)
}

func New(db *sql.DB) UserRepository {
	return &sUserRepository{DB: db}
}

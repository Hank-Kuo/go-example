package task

import (
	"context"

	"go-example/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*models.Task, error)
	Get(ctx context.Context, taskID int) (*models.Task, error)
	Create(ctx context.Context, task *models.Task) error
	Update(ctx context.Context, taskID int, name string, status int) error
	Delete(ctx context.Context, taskID int) error
}

type taskRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) Repository {
	return &taskRepo{db: db}
}

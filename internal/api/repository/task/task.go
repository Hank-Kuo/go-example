package task

import (
	"context"

	"github.com/Hank-Kuo/go-example/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetAll(ctx context.Context, cursorID int, limit int) ([]*models.Task, error)
	Get(ctx context.Context, taskID int) (*models.Task, error)
	Create(ctx context.Context, task *models.Task) (*models.Task, error)
	Update(ctx context.Context, taskID int, name string, status int) (*models.Task, error)
	Delete(ctx context.Context, taskID int) error
}

type taskRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) Repository {
	return &taskRepo{db: db}
}

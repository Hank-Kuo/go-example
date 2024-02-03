package task

import (
	"context"

	"go-example/config"
	taskRepo "go-example/internal/api/repository/task"
	"go-example/internal/models"
	"go-example/pkg/logger"
	"go-example/pkg/tracer"

	"github.com/pkg/errors"
)

type Service interface {
	GetAll(ctx context.Context) ([]*models.Task, error)
	Get(ctx context.Context, taskID int) (*models.Task, error)
	Create(ctx context.Context, task *models.Task) error
	Update(ctx context.Context, taskID int, name string, status int) error
	Delete(ctx context.Context, taskID int) error
}

type taskSrv struct {
	cfg      *config.Config
	taskRepo taskRepo.Repository
	logger   logger.Logger
}

func NewService(cfg *config.Config, taskRepo taskRepo.Repository, logger logger.Logger) Service {
	return &taskSrv{
		cfg:      cfg,
		taskRepo: taskRepo,
		logger:   logger,
	}
}

func (srv *taskSrv) GetAll(ctx context.Context) ([]*models.Task, error) {
	c, span := tracer.NewSpan(ctx, "TaskService.GetAll", nil)
	defer span.End()

	tasks, err := srv.taskRepo.GetAll(c)
	if err != nil {
		tracer.AddSpanError(span, err)
		return nil, errors.Wrap(err, "TaskService.GetAll")
	}
	return tasks, nil
}

func (srv *taskSrv) Get(ctx context.Context, taskID int) (*models.Task, error) {
	c, span := tracer.NewSpan(ctx, "TaskService.Get", nil)
	defer span.End()

	task, err := srv.taskRepo.Get(c, taskID)
	if err != nil {
		tracer.AddSpanError(span, err)
		return nil, errors.Wrap(err, "TaskService.Get")
	}
	return task, nil
}

func (srv *taskSrv) Create(ctx context.Context, task *models.Task) error {
	c, span := tracer.NewSpan(ctx, "TaskService.Create", nil)
	defer span.End()

	err := srv.taskRepo.Create(c, task)
	if err != nil {
		tracer.AddSpanError(span, err)
		return errors.Wrap(err, "TaskService.Create")
	}

	return nil
}

func (srv *taskSrv) Update(ctx context.Context, taskID int, name string, status int) error {
	c, span := tracer.NewSpan(ctx, "TaskService.Update", nil)
	defer span.End()

	err := srv.taskRepo.Update(c, taskID, name, status)
	if err != nil {
		tracer.AddSpanError(span, err)
		return errors.Wrap(err, "TaskService.Update")
	}

	return nil
}

func (srv *taskSrv) Delete(ctx context.Context, taskID int) error {
	c, span := tracer.NewSpan(ctx, "TaskService.Delete", nil)
	defer span.End()

	err := srv.taskRepo.Delete(c, taskID)
	if err != nil {
		tracer.AddSpanError(span, err)
		return errors.Wrap(err, "TaskService.Delete")
	}
	return nil
}

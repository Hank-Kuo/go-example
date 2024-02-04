package task

import (
	"context"

	"github.com/Hank-Kuo/go-example/config"
	taskRepo "github.com/Hank-Kuo/go-example/internal/api/repository/task"
	"github.com/Hank-Kuo/go-example/internal/dto"
	"github.com/Hank-Kuo/go-example/internal/models"
	"github.com/Hank-Kuo/go-example/pkg/customError"
	"github.com/Hank-Kuo/go-example/pkg/logger"
	"github.com/Hank-Kuo/go-example/pkg/tracer"
	"github.com/Hank-Kuo/go-example/pkg/utils"

	"github.com/pkg/errors"
)

type Service interface {
	GetAll(ctx context.Context, query *dto.TaskGetAllQueryDto) (*dto.TaskGetAllResDto, error)
	Get(ctx context.Context, taskID int) (*models.Task, error)
	Create(ctx context.Context, task *models.Task) (*models.Task, error)
	Update(ctx context.Context, taskID int, name string, status int) (*models.Task, error)
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

func (srv *taskSrv) GetAll(ctx context.Context, query *dto.TaskGetAllQueryDto) (*dto.TaskGetAllResDto, error) {
	c, span := tracer.NewSpan(ctx, "TaskService.GetAll", nil)
	defer span.End()

	cursor, err := utils.DecodeCursor("task_id", query.Cursor)
	if err != nil {
		tracer.AddSpanError(span, err)
		return nil, customError.ErrBadQueryParams
	}

	tasks, err := srv.taskRepo.GetAll(c, cursor, query.Limit+1)
	if err != nil {
		tracer.AddSpanError(span, err)
		return nil, errors.Wrap(err, "TaskService.GetAll")
	}

	// update tasks & add nextCursor
	var nextCursor string
	if len(tasks) > query.Limit {
		nextCursor = utils.EncodeCursor("task_id", tasks[query.Limit].ID)
		tasks = tasks[:query.Limit]
	}

	return &dto.TaskGetAllResDto{
		Tasks:      tasks,
		NextCursor: nextCursor,
	}, nil

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

func (srv *taskSrv) Create(ctx context.Context, task *models.Task) (*models.Task, error) {
	c, span := tracer.NewSpan(ctx, "TaskService.Create", nil)
	defer span.End()

	t, err := srv.taskRepo.Create(c, task)
	if err != nil {
		tracer.AddSpanError(span, err)
		return nil, errors.Wrap(err, "TaskService.Create")
	}

	return t, nil
}

func (srv *taskSrv) Update(ctx context.Context, taskID int, name string, status int) (*models.Task, error) {
	c, span := tracer.NewSpan(ctx, "TaskService.Update", nil)
	defer span.End()

	t, err := srv.taskRepo.Update(c, taskID, name, status)
	if err != nil {
		tracer.AddSpanError(span, err)
		return nil, errors.Wrap(err, "TaskService.Update")
	}

	return t, nil
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

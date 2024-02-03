package task

import (
	"context"

	"go-example/internal/models"
	"go-example/pkg/tracer"

	"github.com/pkg/errors"
)

func (r *taskRepo) GetAll(ctx context.Context) ([]*models.Task, error) {
	ctx, span := tracer.NewSpan(ctx, "TaskRepo.GetAll", nil)
	defer span.End()

	tasks := []*models.Task{}
	if err := r.db.SelectContext(ctx, &tasks, "SELECT * FROM task"); err != nil {
		tracer.AddSpanError(span, err)
		return nil, errors.Wrap(err, "TaskRepo.GetAll")
	}

	return tasks, nil
}

func (r *taskRepo) Get(ctx context.Context, taskID int) (*models.Task, error) {
	ctx, span := tracer.NewSpan(ctx, "TaskRepo.Get", nil)
	defer span.End()

	task := models.Task{}
	if err := r.db.GetContext(ctx, &task, "SELECT * FROM task WHERE id = $1", taskID); err != nil {
		tracer.AddSpanError(span, err)
		return nil, errors.Wrap(err, "TaskRepo.Get")
	}

	return &task, nil
}

func (r *taskRepo) Create(ctx context.Context, task *models.Task) error {
	ctx, span := tracer.NewSpan(ctx, "TaskRepo.Create", nil)
	defer span.End()

	sqlQuery := `INSERT INTO task(name, status) VALUES ($1, $2)`

	_, err := r.db.ExecContext(ctx, sqlQuery, task.Name, task.Status)
	if err != nil {
		tracer.AddSpanError(span, err)
		return errors.Wrap(err, "TaskRepo.Create")
	}
	return nil
}

func (r *taskRepo) Update(ctx context.Context, taskID int, name string, status int) error {
	ctx, span := tracer.NewSpan(ctx, "TaskRepo.Update", nil)
	defer span.End()

	sqlQuery := `UPDATE task SET name = $1, status = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3`
	_, err := r.db.ExecContext(ctx, sqlQuery, name, status, taskID)

	if err != nil {
		tracer.AddSpanError(span, err)
		return errors.Wrap(err, "TaskRepo.Update")
	}
	return nil
}

func (r *taskRepo) Delete(ctx context.Context, taskID int) error {
	ctx, span := tracer.NewSpan(ctx, "TaskRepo.Delete", nil)
	defer span.End()

	sqlQuery := `DELETE FROM task WHERE id = $1`
	_, err := r.db.ExecContext(ctx, sqlQuery, taskID)

	if err != nil {
		tracer.AddSpanError(span, err)
		return errors.Wrap(err, "TaskRepo.Delete")
	}
	return nil
}

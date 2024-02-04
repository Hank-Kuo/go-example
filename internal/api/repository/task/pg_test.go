package task

import (
	"context"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	models "github.com/Hank-Kuo/go-example/internal/models"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	ctx := context.Background()
	taskRepo := NewRepo(sqlxDB)

	t.Run("TaskRepo GetAll", func(t *testing.T) {
		mockData := []models.Task{
			{ID: 1, Name: "Task 1", Status: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 2, Name: "Task 2", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 3, Name: "Task 3", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}

		rows := sqlmock.NewRows([]string{"id", "name", "status", "created_at", "updated_at"})

		for _, d := range mockData {
			rows.AddRow(d.ID, d.Name, d.Status, d.CreatedAt, d.UpdatedAt)
		}
		mock.ExpectQuery("SELECT * FROM task WHERE id >= $1 ORDER BY id LIMIT $2").WithArgs(1, 10).WillReturnRows(rows)

		tasks, err := taskRepo.GetAll(ctx, 1, 10)

		assert.NoError(t, err)
		assert.NotNil(t, tasks)
		assert.Equal(t, 3, len(tasks))
	})

}

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	ctx := context.Background()
	taskRepo := NewRepo(sqlxDB)

	t.Run("TaskRepo Get", func(t *testing.T) {

		mockData := []models.Task{
			{ID: 1, Name: "Task 1", Status: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}

		rows := sqlmock.NewRows([]string{"id", "name", "status", "created_at", "updated_at"})

		for _, d := range mockData {
			rows.AddRow(d.ID, d.Name, d.Status, d.CreatedAt, d.UpdatedAt)
		}
		mock.ExpectQuery("SELECT * FROM task WHERE id = $1").WithArgs(1).WillReturnRows(rows)

		task, err := taskRepo.Get(ctx, 1)

		assert.NoError(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, "Task 1", task.Name)
	})
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	ctx := context.Background()
	taskRepo := NewRepo(sqlxDB)

	t.Run("TaskRepo Delete", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM task WHERE id = $1").WithArgs(12).WillReturnResult(sqlmock.NewResult(12, 1))
		err = taskRepo.Delete(ctx, 12)
		assert.NoError(t, err)
	})
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	ctx := context.Background()
	taskRepo := NewRepo(sqlxDB)

	t.Run("TaskRepo Update", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "status", "created_at", "updated_at"}).AddRow(1, "Task 1", 1, time.Now(), time.Now())

		mock.ExpectQuery("UPDATE task SET name = $1, status = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3 RETURNING *").
			WithArgs("Task 1", 1, 1).WillReturnRows(rows)

		result, err := taskRepo.Update(ctx, 1, "Task 1", 1)

		assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 1, result.ID)
		assert.Equal(t, "Task 1", result.Name)
		assert.Equal(t, 1, result.Status)
	})
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	ctx := context.Background()
	taskRepo := NewRepo(sqlxDB)

	t.Run("TaskRepo Create", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "status", "created_at", "updated_at"}).AddRow(1, "Task 1", 0, time.Now(), time.Now())

		task := &models.Task{
			ID:     1,
			Name:   "Task 1",
			Status: 0,
		}

		mock.ExpectQuery(`INSERT INTO task (name, status) VALUES ($1, $2) RETURNING *`).
			WithArgs("Task 1", 0).WillReturnRows(rows)

		result, err := taskRepo.Create(ctx, task)

		assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 1, result.ID)
		assert.Equal(t, "Task 1", result.Name)
		assert.Equal(t, 0, result.Status)
	})
}

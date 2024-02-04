package task

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Hank-Kuo/go-example/internal/api/repository/task/mock"
	"github.com/Hank-Kuo/go-example/internal/dto"
	"github.com/Hank-Kuo/go-example/internal/models"
	"github.com/Hank-Kuo/go-example/pkg/logger"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mocktaskRepo := mock.NewMockRepository(ctrl)
	srv := NewService(nil, mocktaskRepo, apiLogger)

	t.Run("TaskService Create", func(t *testing.T) {
		inputTask := &models.Task{
			Name:   "Task 1",
			Status: 0,
		}
		mockTask := &models.Task{
			ID:        1,
			Name:      "Task 1",
			Status:    0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mocktaskRepo.EXPECT().Create(gomock.Any(), gomock.Eq(inputTask)).Return(mockTask, nil)

		ctx := context.Background()
		result, err := srv.Create(ctx, inputTask)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		assert.Equal(t, 1, result.ID)
		assert.Equal(t, "Task 1", result.Name)
		assert.Equal(t, 0, result.Status)
	})
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mocktaskRepo := mock.NewMockRepository(ctrl)
	srv := NewService(nil, mocktaskRepo, apiLogger)

	t.Run("TaskService GetAll", func(t *testing.T) {
		mockTasks := []*models.Task{
			{ID: 1, Name: "Task 1", Status: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 2, Name: "Task 2", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 3, Name: "Task 3", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}
		limit := 3
		mocktaskRepo.EXPECT().GetAll(gomock.Any(), 1, limit+1).Return(mockTasks, nil)

		ctx := context.Background()

		query := &dto.TaskGetAllQueryDto{
			Cursor: "",
			Limit:  limit,
		}
		result, err := srv.GetAll(ctx, query)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 3, len(result.Tasks))
	})
}
func TestGet(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mocktaskRepo := mock.NewMockRepository(ctrl)
	srv := NewService(nil, mocktaskRepo, apiLogger)

	t.Run("TaskService Get", func(t *testing.T) {
		mockTask := &models.Task{
			ID:        1,
			Name:      "Task 1",
			Status:    0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mocktaskRepo.EXPECT().Get(gomock.Any(), 1).Return(mockTask, nil)

		ctx := context.Background()
		result, err := srv.Get(ctx, 1)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "Task 1", result.Name)
	})
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mocktaskRepo := mock.NewMockRepository(ctrl)
	srv := NewService(nil, mocktaskRepo, apiLogger)

	t.Run("TaskService Update", func(t *testing.T) {

		mockTask := &models.Task{
			ID:        1,
			Name:      "Task 1",
			Status:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mocktaskRepo.EXPECT().Update(gomock.Any(), 1, "Task 1", 1).Return(mockTask, nil)

		ctx := context.Background()
		result, err := srv.Update(ctx, 1, "Task 1", 1)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		assert.Equal(t, 1, result.ID)
		assert.Equal(t, "Task 1", result.Name)
		assert.Equal(t, 1, result.Status)
	})
}

func TestDelete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mocktaskRepo := mock.NewMockRepository(ctrl)
	srv := NewService(nil, mocktaskRepo, apiLogger)

	t.Run("TaskService Delete", func(t *testing.T) {

		mocktaskRepo.EXPECT().Delete(gomock.Any(), 1).Return(nil)

		ctx := context.Background()
		err := srv.Delete(ctx, 1)
		assert.NoError(t, err)
	})
}

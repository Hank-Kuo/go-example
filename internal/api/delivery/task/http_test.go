package task

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	mock_taskSrv "github.com/Hank-Kuo/go-example/internal/api/service/task/mock"
	"github.com/Hank-Kuo/go-example/internal/dto"
	"github.com/Hank-Kuo/go-example/internal/models"
	"github.com/Hank-Kuo/go-example/pkg/logger"
)

func TestHttpHandler_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockTaskSrv := mock_taskSrv.NewMockService(ctrl)

	handler := &httpHandler{
		taskSrv: mockTaskSrv,
		logger:  apiLogger,
	}

	w := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(w)
	engine.GET("/api/tasks", handler.GetAll)

	t.Run("TaskHttpDelivery GetAll", func(t *testing.T) {
		query := &dto.TaskGetAllQueryDto{
			Cursor: "",
			Limit:  10,
		}
		mockTask := &dto.TaskGetAllResDto{
			Tasks:      []*models.Task{},
			NextCursor: "",
		}

		mockTaskSrv.EXPECT().GetAll(gomock.Any(), query).Return(mockTask, nil)

		req, _ := http.NewRequest("GET", "/api/tasks", nil)
		engine.ServeHTTP(w, req)

		expectedBody := `{"status":"success","message":"get tasks successfully","data":{"tasks":[],"next_cursor":""}}`

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedBody, w.Body.String())
	})
}

func TestHttpHandler_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockTaskSrv := mock_taskSrv.NewMockService(ctrl)

	handler := &httpHandler{
		taskSrv: mockTaskSrv,
		logger:  apiLogger,
	}

	w := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(w)
	engine.GET("/api/tasks/:id", handler.Get)

	t.Run("TaskHttpDelivery Get", func(t *testing.T) {
		mockTaskSrv.EXPECT().Get(gomock.Any(), 1).Return(
			&models.Task{
				ID:        1,
				Name:      "Task 1",
				Status:    0,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

		req, _ := http.NewRequest("GET", "/api/tasks/1", nil)
		engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestHttpHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockTaskSrv := mock_taskSrv.NewMockService(ctrl)

	handler := &httpHandler{
		taskSrv: mockTaskSrv,
		logger:  apiLogger,
	}

	w := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(w)
	engine.POST("/api/tasks", handler.Create)

	inputTask := &models.Task{
		Name:   "Task 1",
		Status: 0,
	}
	returnTask := &models.Task{
		ID:        1,
		Name:      "Task 1",
		Status:    0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	t.Run("TaskHttpDelivery Create", func(t *testing.T) {
		mockTaskSrv.EXPECT().Create(gomock.Any(), gomock.Eq(inputTask)).Return(returnTask, nil)

		reqBody := `{"name": "Task 1", "status": 0}`
		req, err := http.NewRequest("POST", "/api/tasks", strings.NewReader(reqBody))
		engine.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestHttpHandler_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockTaskSrv := mock_taskSrv.NewMockService(ctrl)

	handler := &httpHandler{
		taskSrv: mockTaskSrv,
		logger:  apiLogger,
	}

	w := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(w)
	engine.PUT("/api/tasks/:id", handler.Update)

	t.Run("TaskHttpDelivery Update", func(t *testing.T) {
		returnTask := &models.Task{
			ID:        1,
			Name:      "Task 1",
			Status:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mockTaskSrv.EXPECT().Update(gomock.Any(), 1, "Task 1", 1).Return(returnTask, nil)

		reqBody := `{"name": "Task 1", "status": 1}`
		req, err := http.NewRequest("PUT", "/api/tasks/1", strings.NewReader(reqBody))
		engine.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

}

func TestHttpHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockTaskSrv := mock_taskSrv.NewMockService(ctrl)

	handler := &httpHandler{
		taskSrv: mockTaskSrv,
		logger:  apiLogger,
	}

	w := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(w)
	engine.DELETE("/api/tasks/:id", handler.Delete)

	t.Run("TaskHttpDelivery Delete", func(t *testing.T) {
		mockTaskSrv.EXPECT().Delete(gomock.Any(), 1).Return(nil)
		reqBody := `{"name": "Task 1", "status": 1}`
		req, err := http.NewRequest("DELETE", "/api/tasks/1", strings.NewReader(reqBody))
		engine.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

}

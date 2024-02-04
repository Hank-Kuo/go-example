package task

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	taskSrv "github.com/Hank-Kuo/go-example/internal/api/service/task"
	"github.com/Hank-Kuo/go-example/internal/dto"
	"github.com/Hank-Kuo/go-example/internal/models"
	httpErr "github.com/Hank-Kuo/go-example/pkg/customError/httpError"
	"github.com/Hank-Kuo/go-example/pkg/logger"
	httpResponse "github.com/Hank-Kuo/go-example/pkg/response/http_response"
	"github.com/Hank-Kuo/go-example/pkg/tracer"
)

type httpHandler struct {
	taskSrv taskSrv.Service
	logger  logger.Logger
}

func NewHttpHandler(e *gin.RouterGroup, taskSrv taskSrv.Service, logger logger.Logger) {
	handler := &httpHandler{
		taskSrv: taskSrv,
		logger:  logger,
	}

	e.GET("/tasks", handler.GetAll)
	e.GET("/task/:id", handler.Get)
	e.POST("/tasks", handler.Create)
	e.PUT("/tasks/:id", handler.Update)
	e.DELETE("/tasks/:id", handler.Delete)
}

// GetAll
// @Summary get all tasks
// @Description get all tasks
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {object} models.Task
// @Router /tasks [get]
func (h *httpHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, span := tracer.NewSpan(ctx, "TaskHttpHandler.GetAll", nil)
	defer span.End()

	var queryParams dto.TaskGetAllQueryDto
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		httpResponse.Fail(err, h.logger).ToJSON(c)
		return
	}

	tasks, err := h.taskSrv.GetAll(ctx, &queryParams)
	if err != nil {
		tracer.AddSpanError(span, err)
		httpResponse.Fail(err, h.logger).ToJSON(c)
		return
	}
	httpResponse.OK(http.StatusOK, "get tasks successfully", tasks).ToJSON(c)
}

// Get
// @Summary get task
// @Description get task
// @Tags Task
// @Accept  json
// @Produce  json
// @Success 200 {object} http_response.ResponseBody{Data:models.Task}
// @Router /task/:id [get]
func (h *httpHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, span := tracer.NewSpan(ctx, "TaskHttpHandler.Get", nil)
	defer span.End()

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httpResponse.Fail(httpErr.NewBadQueryParamsError(err), h.logger).ToJSON(c)
		return
	}

	task, err := h.taskSrv.Get(ctx, taskID)

	if err != nil {
		tracer.AddSpanError(span, err)
		httpResponse.Fail(err, h.logger).ToJSON(c)
		return
	}
	httpResponse.OK(http.StatusOK, "get task successfully", task).ToJSON(c)
}

// Create
// @Summary Create new task
// @Description create new task
// @Tags Task
// @Accept  json
// @Produce  json
// @Success 201 {object} models.Task
// @Router /tasks [post]
func (h *httpHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, span := tracer.NewSpan(ctx, "TaskHttpHandler.Create", nil)
	defer span.End()

	var body dto.TaskReqDto
	if err := c.ShouldBindJSON(&body); err != nil {
		httpResponse.Fail(err, h.logger).ToJSON(c)
		return
	}

	task, err := h.taskSrv.Create(ctx, &models.Task{
		Name:   body.Name,
		Status: *body.Status,
	})
	if err != nil {
		tracer.AddSpanError(span, err)
		httpResponse.Fail(err, h.logger).ToJSON(c)
		return
	}

	httpResponse.OK(http.StatusOK, "create task successfully", task).ToJSON(c)
}

// Update
// @Summary update task
// @Description update task
// @Tags Task
// @Accept  json
// @Produce  json
// @Success 201 {object} models.Task
// @Router /tasks/:id [put]
func (h *httpHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, span := tracer.NewSpan(ctx, "TaskHttpHandler.Create", nil)
	defer span.End()

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httpResponse.Fail(httpErr.NewBadQueryParamsError(err), h.logger).ToJSON(c)
		return
	}

	var body dto.TaskReqDto
	if err := c.ShouldBindJSON(&body); err != nil {
		httpResponse.Fail(err, h.logger).ToJSON(c)
		return
	}

	task, err := h.taskSrv.Update(ctx, taskID, body.Name, *body.Status)
	if err != nil {
		tracer.AddSpanError(span, err)
		httpResponse.Fail(err, h.logger).ToJSON(c)
		return
	}

	httpResponse.OK(http.StatusOK, "update task successfully", task).ToJSON(c)
}

// Delete
// @Summary delete task
// @Description deletee task
// @Tags Task
// @Accept  json
// @Produce  json
// @Success 200
// @Router /tasks/:id [delete]
func (h *httpHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, span := tracer.NewSpan(ctx, "TaskHttpHandler.Delete", nil)
	defer span.End()

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httpResponse.Fail(httpErr.NewBadQueryParamsError(err), h.logger).ToJSON(c)
		return
	}

	err = h.taskSrv.Delete(ctx, taskID)

	if err != nil {
		tracer.AddSpanError(span, err)
		httpResponse.Fail(err, h.logger).ToJSON(c)
		return
	}
	httpResponse.OK(http.StatusOK, "delete task successfully", nil).ToJSON(c)
}

package task

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	taskSrv "go-example/internal/api/service/task"
	"go-example/internal/models"
	httpErr "go-example/pkg/customError/httpError"
	"go-example/pkg/logger"
	httpResponse "go-example/pkg/response/http_response"
	"go-example/pkg/tracer"
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

func (h *httpHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, span := tracer.NewSpan(ctx, "TaskHttpHandler.GetAll", nil)
	defer span.End()

	tasks, err := h.taskSrv.GetAll(ctx)
	if err != nil {
		tracer.AddSpanError(span, err)
		httpResponse.Fail(err, h.logger).ToJSON(c)
		return
	}
	httpResponse.OK(http.StatusOK, "get tasks successfully", tasks).ToJSON(c)
}

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

func (h *httpHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, span := tracer.NewSpan(ctx, "TaskHttpHandler.Create", nil)
	defer span.End()

	var body models.Task
	if err := c.ShouldBindJSON(&body); err != nil {
		httpResponse.Fail(err, h.logger).ToJSON(c)
		return
	}

	if err := h.taskSrv.Create(ctx, &body); err != nil {
		tracer.AddSpanError(span, err)
		httpResponse.Fail(err, h.logger).ToJSON(c)
		return
	}

	httpResponse.OK(http.StatusOK, "create task successfully", nil).ToJSON(c)
}

func (h *httpHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, span := tracer.NewSpan(ctx, "TaskHttpHandler.Create", nil)
	defer span.End()

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httpResponse.Fail(httpErr.NewBadQueryParamsError(err), h.logger).ToJSON(c)
		return
	}

	var body models.Task
	if err := c.ShouldBindJSON(&body); err != nil {
		httpResponse.Fail(err, h.logger).ToJSON(c)
		return
	}

	if err := h.taskSrv.Update(ctx, taskID, body.Name, *body.Status); err != nil {
		tracer.AddSpanError(span, err)
		httpResponse.Fail(err, h.logger).ToJSON(c)
		return
	}

	httpResponse.OK(http.StatusOK, "update task successfully", nil).ToJSON(c)
}

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

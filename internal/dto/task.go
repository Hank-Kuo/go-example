package dto

import (
	"github.com/Hank-Kuo/go-example/internal/models"
)

type TaskGetAllQueryDto struct {
	Cursor string `form:"cursor"`
	Limit  int    `form:"limit,default=10"`
}

type TaskGetAllResDto struct {
	Tasks      []*models.Task `json:"tasks"`
	NextCursor string         `json:"next_cursor"`
}

type TaskReqDto struct {
	Name   string `json:"name" binding:"required"`
	Status *int   `json:"status" binding:"required,min=0,max=1"`
}

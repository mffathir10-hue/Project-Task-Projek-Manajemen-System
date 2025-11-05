package serviceroute

import (
	. "gintugas/modules/components/Tasks/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService TaskService
}

func NewTaskController(taskService TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (c *TaskHandler) CreateTask(ctx *gin.Context) {
	task, err := c.taskService.CreateTask(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"task":    task,
	})
}

func (c *TaskHandler) GetProjectTasks(ctx *gin.Context) {
	tasks, err := c.taskService.GetProjectTasks(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Tasks retrieved successfully",
		"tasks":   tasks,
	})
}

func (c *TaskHandler) GetTaskByID(ctx *gin.Context) {
	task, err := c.taskService.GetTaskByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task retrieved successfully",
		"task":    task,
	})
}

func (c *TaskHandler) UpdateTask(ctx *gin.Context) {
	task, err := c.taskService.UpdateTask(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
		"task":    task,
	})
}

func (c *TaskHandler) DeleteTask(ctx *gin.Context) {
	if err := c.taskService.DeleteTask(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}

package serviceroute

import (
	. "gintugas/modules/components/Tasks/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService TaskService
}

func NewTaskController(taskService TaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
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

func (c *TaskController) GetProjectTasks(ctx *gin.Context) {
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

func (c *TaskController) GetTaskByID(ctx *gin.Context) {
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

func (c *TaskController) UpdateTask(ctx *gin.Context) {
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

func (c *TaskController) DeleteTask(ctx *gin.Context) {
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

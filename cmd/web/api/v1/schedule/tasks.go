package schedule

import (
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/errors"
	"meteo/internal/kit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p scheduleAPI) GetAllTasks(c *gin.Context) {
	tasks, err := p.repo.GetAllTasks(dto.Pageable{})
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": tasks})
}

func (p scheduleAPI) AddTask(c *gin.Context) {

	var task entities.Tasks

	if err := c.ShouldBind(&task); err != nil ||
		len(task.ID) == 0 ||
		len(task.Name) == 0 {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.AddTask(task)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": task.ID})
}

func (p scheduleAPI) EditTask(c *gin.Context) {

	var task entities.Tasks

	if err := c.ShouldBind(&task); err != nil ||
		len(task.ID) == 0 ||
		len(task.Name) == 0 {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.EditTask(task)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p scheduleAPI) DelTask(c *gin.Context) {
	if err := p.repo.DelTask(c.Param("id")); err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p scheduleAPI) RunTask(c *gin.Context) {
	var api struct {
		Api string `json:"api"`
	}

	if err := c.ShouldBind(&api); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	_, err := kit.PutInt(api.Api, nil)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

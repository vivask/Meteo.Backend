package schedule

import (
	"meteo/internal/dto"
	"meteo/internal/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p scheduleAPI) GetAllTasks(c *gin.Context) {
	tasks, err := p.repo.GetAllTasks(dto.Pageable{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": tasks})
}

func (p scheduleAPI) AddTask(c *gin.Context) {

	var task entities.Tasks

	if err := c.ShouldBind(&task); err != nil ||
		len(task.ID) == 0 ||
		len(task.Name) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.repo.AddTask(task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p scheduleAPI) EditTask(c *gin.Context) {

	var task entities.Tasks

	if err := c.ShouldBind(&task); err != nil ||
		len(task.ID) == 0 ||
		len(task.Name) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.repo.EditTask(task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p scheduleAPI) DelTask(c *gin.Context) {
	if err := p.repo.DelTask(c.Param("id")); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

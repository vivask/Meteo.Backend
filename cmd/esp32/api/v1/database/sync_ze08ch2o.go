package database

import (
	"meteo/internal/entities"
	"net/http"

	lock "meteo/internal/repo/esp32"

	"github.com/gin-gonic/gin"
)

func (p databaseAPI) GetNotSyncZe08ch2o(c *gin.Context) {
	bmx280, err := p.repo.GetNotSyncZe08ch2o()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEEER",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bmx280)
}

func (p databaseAPI) AddSyncZe08ch2o(c *gin.Context) {

	var bmx280 []entities.Ze08ch2o

	if err := c.ShouldBind(&bmx280); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "DATABASEERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.repo.AddSyncZe08ch2o(bmx280)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEEER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) LockZe08ch2o(c *gin.Context) {
	err := lock.LockZe08ch2o(false)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEEER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) UnlockZe08ch2o(c *gin.Context) {
	err := lock.UnlockZe08ch2o(false)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEEER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) SyncZe08ch2o(c *gin.Context) {
	err := p.repo.SyncZe08ch2o()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEEER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) ReplaceZe08ch2o(c *gin.Context) {

	var bmx280 []entities.Ze08ch2o

	if err := c.ShouldBind(&bmx280); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "DATABASEERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.repo.ReplaceZe08ch2o(bmx280)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEEER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

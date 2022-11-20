package database

import (
	"meteo/internal/entities"
	"net/http"

	lock "meteo/internal/repo/esp32"

	"github.com/gin-gonic/gin"
)

func (p databaseAPI) GetNotSyncBmx280(c *gin.Context) {
	bmx280, err := p.repo.GetNotSyncBmx280()
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

func (p databaseAPI) AddSyncBmx280(c *gin.Context) {

	var bmx280 []entities.Bmx280

	if err := c.ShouldBind(&bmx280); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "DATABASEERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.repo.AddSyncBmx280(bmx280)
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

func (p databaseAPI) LockBmx280(c *gin.Context) {
	err := lock.LockBmx280(false)
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

func (p databaseAPI) UnlockBmx280(c *gin.Context) {
	err := lock.UnlockBmx280(false)
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

func (p databaseAPI) SyncBmx280(c *gin.Context) {
	err := p.repo.SyncBmx280()
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

func (p databaseAPI) ReplaceBmx280(c *gin.Context) {

	var bmx280 []entities.Bmx280

	if err := c.ShouldBind(&bmx280); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "DATABASEERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.repo.ReplaceBmx280(bmx280)
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

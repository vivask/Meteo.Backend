package database

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p databaseAPI) SyncEsp32Tables(c *gin.Context) {
	err := p.repo.SyncBmx280()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEEER",
				"message": err.Error()})
		return
	}

	err = p.repo.SyncDs18b20()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEEER",
				"message": err.Error()})
		return
	}

	err = p.repo.SyncMics6814()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEEER",
				"message": err.Error()})
		return
	}

	err = p.repo.SyncZe08ch2o()
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

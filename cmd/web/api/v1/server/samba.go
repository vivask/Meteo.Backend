package server

import (
	"meteo/internal/kit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p serverAPI) SambaStart(c *gin.Context) {

	_, err := kit.PutMain("/server/samba/start", nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) SambaStop(c *gin.Context) {

	_, err := kit.PutMain("/server/samba/stop", nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) SambaRestart(c *gin.Context) {

	_, err := kit.PutMain("/server/samba/restart", nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

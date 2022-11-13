package server

import (
	"meteo/internal/kit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p serverAPI) TransmissionStartJobs(c *gin.Context) {

	_, err := kit.PutMain("/server/transmission/jobs/start", nil)

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

func (p serverAPI) TransmissionStopJobs(c *gin.Context) {

	_, err := kit.PutMain("/server/transmission/jobs/stop", nil)

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

func (p serverAPI) TransmissionStart(c *gin.Context) {

	_, err := kit.PutMain("/server/transmission/start", nil)

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

func (p serverAPI) TransmissionStop(c *gin.Context) {

	_, err := kit.PutMain("/server/transmission/stop", nil)

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

func (p serverAPI) TransmissionRestart(c *gin.Context) {

	_, err := kit.PutMain("/server/transmission/restart", nil)

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

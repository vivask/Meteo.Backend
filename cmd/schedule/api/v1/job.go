package v1

import (
	"fmt"
	"meteo/internal/kit"
	"meteo/internal/log"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p scheduleAPI) JobsReload(c *gin.Context) {
	err := p.reloadJobs()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SCHEDULEER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p scheduleAPI) JobRun(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "SCHEDULEER",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	job, err := p.repo.GetJobByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SCHEDULEER",
				"message": err.Error()})
		return
	}

	_, err = kit.PutInt(job.Task.Api, job.Params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SCHEDULEER",
				"message": fmt.Sprintf("Not implemented Api [%s] for shedule task: %v", job.Task.Api, err)})
		return
	}
	c.Status(http.StatusOK)
}

func (p scheduleAPI) GetCronJobs(c *gin.Context) {
	jobs := p.getCronJobs()
	log.Infof("CRON: %v", jobs)
	c.JSON(http.StatusOK, jobs)
}

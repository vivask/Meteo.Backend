package v1

import (
	"meteo/internal/errors"
	"meteo/internal/kit"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p scheduleAPI) JobsReload(c *gin.Context) {
	err := p.reloadJobs()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p scheduleAPI) JobRun(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	job, err := p.repo.GetJobByID(id)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	_, err = kit.PutInt(job.Task.Api, job.Params)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p scheduleAPI) GetCronJobs(c *gin.Context) {
	jobs := p.getCronJobs()
	//log.Infof("CRON: %v", jobs)
	c.JSON(http.StatusOK, jobs)
}

package schedule

import (
	"encoding/json"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/errors"
	"meteo/internal/kit"
	"meteo/internal/log"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p scheduleAPI) GetAllJobs(c *gin.Context) {
	jobs, err := p.repo.GetAllJobs(dto.Pageable{})
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": jobs})
}

func (p scheduleAPI) AddJob(c *gin.Context) {

	var err error

	var job entities.Jobs

	if err = c.ShouldBind(&job); err != nil ||
		len(job.Note) == 0 ||
		len(job.Executor.ID) == 0 ||
		len(job.Task.ID) == 0 ||
		len(job.Period.ID) == 0 {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	var data struct {
		Id       uint32 `json:"id"`
		Activate bool   `json:"activate"`
	}

	data.Id, data.Activate, err = p.repo.AddJob(job)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p scheduleAPI) EditJob(c *gin.Context) {

	var err error
	var job entities.Jobs

	if err := c.ShouldBind(&job); err != nil ||
		len(job.Note) == 0 ||
		len(job.Executor.ID) == 0 ||
		len(job.Task.ID) == 0 ||
		len(job.Period.ID) == 0 {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	var data struct {
		Activate bool `json:"activate"`
	}

	data.Activate, err = p.repo.EditJob(job)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (p scheduleAPI) DelJob(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := p.repo.DeleteJob(id); err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p scheduleAPI) ActivateJob(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := p.repo.ActivateJob(id); err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p scheduleAPI) DeactivateJob(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := p.repo.DeactivateJob(id, true); err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p scheduleAPI) RunJob(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := p.repo.RunJob(id); err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p scheduleAPI) GetAllPeriods(c *gin.Context) {
	periods, err := p.repo.GetAllPeriods(dto.Pageable{})
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": periods})
}

func (p scheduleAPI) GetAllExecutors(c *gin.Context) {
	executors, err := p.repo.GetAllExecutors(dto.Pageable{})
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": executors})
}

func (p scheduleAPI) GetCronJobs(c *gin.Context) {
	body, err := kit.GetInt("/schedule/cron/get")
	if err != nil {
		log.Warningf("Local Schedule Server not responding: %v", err)
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	} else {
		var cronJobs []entities.CronJobs
		err = json.Unmarshal(body, &cronJobs)
		if err != nil {
			c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": cronJobs})
	}
}

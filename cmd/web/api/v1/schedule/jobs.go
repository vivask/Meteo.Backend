package schedule

import (
	"encoding/json"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/log"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p scheduleAPI) GetAllJobs(c *gin.Context) {
	jobs, err := p.repo.GetAllJobs(dto.Pageable{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": jobs})
}

func (p scheduleAPI) AddJob(c *gin.Context) {

	var job entities.Jobs

	if err := c.ShouldBind(&job); err != nil ||
		len(job.Note) == 0 ||
		len(job.Executor.ID) == 0 ||
		len(job.Task.ID) == 0 ||
		len(job.Period.ID) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	if job.Day.ID == 0 {
		job.Day.ID = 1
	}

	err := p.repo.AddJob(job)
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

func (p scheduleAPI) EditJob(c *gin.Context) {

	var job entities.Jobs

	if err := c.ShouldBind(&job); err != nil ||
		len(job.Note) == 0 ||
		len(job.Executor.ID) == 0 ||
		len(job.Task.ID) == 0 ||
		len(job.Period.ID) == 0 ||
		job.Day.ID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.repo.EditJob(job)
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

func (p scheduleAPI) DelJob(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}

	if err := p.repo.DeleteJob(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p scheduleAPI) ActivateJob(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}

	if err := p.repo.ActivateJob(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p scheduleAPI) DeactivateJob(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}

	if err := p.repo.DeactivateJob(id, true); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p scheduleAPI) RunJob(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}

	if err := p.repo.RunJob(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p scheduleAPI) GetAllPeriods(c *gin.Context) {
	periods, err := p.repo.GetAllPeriods(dto.Pageable{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": periods})
}

func (p scheduleAPI) GetAllDays(c *gin.Context) {
	days, err := p.repo.GetAllDays(dto.Pageable{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": days})
}

func (p scheduleAPI) GetAllExecutors(c *gin.Context) {
	executors, err := p.repo.GetAllExecutors(dto.Pageable{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": executors})
}

func (p scheduleAPI) GetCronJobs(c *gin.Context) {
	body, err := kit.GetInt("/schedule/cron/get")
	if err != nil {
		log.Warningf("Local Schedule Server not responding: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	} else {
		var cronJobs []entities.CronJobs
		err = json.Unmarshal(body, &cronJobs)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
		log.Infof("JOBS: %v", cronJobs)
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": cronJobs})
	}
}

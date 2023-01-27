package media

import (
	"meteo/internal/errors"
	"meteo/internal/kit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p mediaAPI) TransmissionStartJobs(c *gin.Context) {

	_, err := kit.PutMain("/media/transmission/jobs/start", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mediaAPI) TransmissionStopJobs(c *gin.Context) {

	_, err := kit.PutMain("/media/transmission/jobs/stop", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mediaAPI) TransmissionStart(c *gin.Context) {

	_, err := kit.PutMain("/media/transmission/start", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mediaAPI) TransmissionStop(c *gin.Context) {

	_, err := kit.PutMain("/media/transmission/stop", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mediaAPI) TransmissionRestart(c *gin.Context) {

	_, err := kit.PutMain("/media/transmission/restart", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

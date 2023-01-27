package media

import (
	"meteo/internal/errors"
	"meteo/internal/kit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p mediaAPI) SambaStart(c *gin.Context) {

	_, err := kit.PutMain("/media/samba/start", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mediaAPI) SambaStop(c *gin.Context) {

	_, err := kit.PutMain("/media/samba/stop", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mediaAPI) SambaRestart(c *gin.Context) {

	_, err := kit.PutMain("/media/samba/restart", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

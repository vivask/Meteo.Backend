package v1

import (
	"meteo/internal/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p sshclientAPI) GetHostFinger(c *gin.Context) {
	var touch entities.Touch
	if err := c.ShouldBind(&touch); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "SSHERR-1",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	finger, err := p.ssh.GetFinger(touch.User, touch.Host)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SSHERR-1",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, finger)
}

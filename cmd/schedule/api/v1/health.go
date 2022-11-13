package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p scheduleAPI) GetHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}

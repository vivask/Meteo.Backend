package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p webAPI) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, "healthy")
}

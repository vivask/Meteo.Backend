package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p proxyAPI) GetHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p esp32API) GetHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}

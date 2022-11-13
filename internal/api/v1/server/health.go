package server

import (
	"meteo/internal/config"
	"meteo/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p serverAPI) GetHealth(c *gin.Context) {
	err := p.HealthRadius()
	if err != nil {
		log.Errorf("radius health error: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVEREER",
				"message": err.Error()})
		return
	}
	if config.Default.App.Server == "main" {
		err = p.HealthStorage()
		if err != nil {
			log.Errorf("storage health error: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "SERVEREER",
					"message": err.Error()})
			return
		}
		err = p.HealthSamba()
		if err != nil {
			log.Errorf("samba health error: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "SERVEREER",
					"message": err.Error()})
			return
		}
		err = p.HealthTransmission()
		if err != nil {
			log.Errorf("transmission health error: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "SERVEREER",
					"message": err.Error()})
			return
		}
	}
	c.Status(http.StatusOK)
}

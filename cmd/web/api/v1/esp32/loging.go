package esp32

import (
	"meteo/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p esp32API) GetAllLogging(c *gin.Context) {
	log, err := p.repo.GetAllLoging(dto.Pageable{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": log})
}

func (p esp32API) JournalClear(c *gin.Context) {
	err := p.repo.JournalClear()
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

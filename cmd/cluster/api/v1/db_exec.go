package v1

import (
	"fmt"
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p clusterAPI) DbExec(c *gin.Context) {

	var cb entities.Callback

	if err := c.ShouldBind(&cb); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "CLUSTERERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.repo.ExecRaw(cb)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "CLUSTERERR",
				"message": err.Error()})
		return
	}

	err = updateHomezone(cb.Query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "CLUSTERERR",
				"message": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func updateHomezone(query string) error {
	_, tblName, err := utils.ParseQuery(query)
	if err != nil {
		return fmt.Errorf("parse query error: %w", err)
	}
	if tblName == "homezones" {
		_, err = kit.PutInt("/proxy/zones/update", nil)
		if err != nil {
			return fmt.Errorf("proxy internal error: %w", err)
		}
	}
	return nil
}

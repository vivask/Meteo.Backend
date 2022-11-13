package proxy

import (
	"meteo/internal/entities"
	"meteo/internal/log"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p proxyAPI) GetAllZones(c *gin.Context) {
	hosts, err := p.repo.GetAllHomeZoneHosts()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": hosts})
}

func (p proxyAPI) AddZone(c *gin.Context) {

	var host entities.Homezone

	if err := c.ShouldBind(&host); err != nil ||
		len(host.Name) == 0 ||
		len(host.Address) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.repo.AddHomeZoneHost(host)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (p proxyAPI) EditZone(c *gin.Context) {

	var host entities.Homezone

	if err := c.ShouldBind(&host); err != nil ||
		len(host.Name) == 0 ||
		len(host.Address) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.repo.EditHomeZoneHost(host)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (p proxyAPI) DelZone(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}

	log.Infof("DelZone ID: %v", id)

	if err := p.repo.DelHomeZoneHost(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (p proxyAPI) SyncZones(c *gin.Context) {
}

func (p proxyAPI) ReloadZones(c *gin.Context) {
}

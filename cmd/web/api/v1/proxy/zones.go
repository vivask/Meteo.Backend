package proxy

import (
	"meteo/internal/entities"
	"meteo/internal/errors"
	"meteo/internal/log"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p proxyAPI) GetAllZones(c *gin.Context) {
	hosts, err := p.repo.GetAllHomeZoneHosts()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": hosts})
}

func (p proxyAPI) AddZone(c *gin.Context) {

	var host entities.Homezone

	if err := c.ShouldBind(&host); err != nil ||
		len(host.Name) == 0 ||
		len(host.Address) == 0 {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	id, err := p.repo.AddHomeZoneHost(host)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": id})
}

func (p proxyAPI) EditZone(c *gin.Context) {
	var host entities.Homezone

	if err := c.ShouldBind(&host); err != nil ||
		len(host.Name) == 0 ||
		len(host.Address) == 0 {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.EditHomeZoneHost(host)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (p proxyAPI) DelZone(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	log.Infof("DelZone ID: %v", id)

	if err := p.repo.DelHomeZoneHost(id); err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (p proxyAPI) SyncZones(c *gin.Context) {
}

func (p proxyAPI) ReloadZones(c *gin.Context) {
}

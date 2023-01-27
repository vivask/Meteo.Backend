package proxy

import (
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/errors"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p proxyAPI) GetAccessLists(c *gin.Context) {
	lists, err := p.repo.GetAccessLists(dto.Pageable{})
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": lists})
}

func (p proxyAPI) GetAllManualToVpn(c *gin.Context) {
	hosts, err := p.repo.GetAllManualToVpn(dto.Pageable{})
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": hosts})
}

func (p proxyAPI) AddManualToVpn(c *gin.Context) {

	var host entities.ToVpnManual

	if err := c.ShouldBind(&host); err != nil ||
		len(host.Name) == 0 ||
		len(host.AccesList.ID) == 0 {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	id, err := p.repo.AddManualToVpn(host)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": id})
}

func (p proxyAPI) EditManualToVpn(c *gin.Context) {

	var host entities.ToVpnManual

	if err := c.ShouldBind(&host); err != nil ||
		len(host.Name) == 0 ||
		len(host.AccesList.ID) == 0 {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.EditManualToVpn(host)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func (p proxyAPI) DelManualFromVpn(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := p.repo.DelManualFromVpn(id); err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p proxyAPI) GetAllAutoToVpn(c *gin.Context) {
	hosts, err := p.repo.GetAllAutoToVpn(dto.Pageable{})
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": hosts})
}

func (p proxyAPI) IgnoreAutoToVpn(c *gin.Context) {

	var hosts []entities.ToVpnAuto

	if err := c.ShouldBind(&hosts); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.IgnoreAutoToVpn(hosts)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p proxyAPI) DelAutoFromVpn(c *gin.Context) {

	var hosts []entities.ToVpnAuto

	if err := c.ShouldBind(&hosts); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.DelAutoFromVpn(hosts)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p proxyAPI) GetAllIgnoreAutoToVpn(c *gin.Context) {
	hosts, err := p.repo.GetAllIgnoreAutoToVpn(dto.Pageable{})
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": hosts})
}

func (p proxyAPI) RestoreAutoToVpn(c *gin.Context) {

	var hosts []entities.ToVpnIgnore

	if err := c.ShouldBind(&hosts); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.RestoreAutoToVpn(hosts)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p proxyAPI) DelIgnoreAutoToVpn(c *gin.Context) {

	var hosts []entities.ToVpnIgnore

	if err := c.ShouldBind(&hosts); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	if err := p.repo.DelIgnoreAutoToVpn(hosts); err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

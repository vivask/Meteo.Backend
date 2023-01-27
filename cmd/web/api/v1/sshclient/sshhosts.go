package sshclient

import (
	"encoding/json"
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/errors"
	"meteo/internal/kit"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p sshclientAPI) GetAllSshHosts(c *gin.Context) {
	hosts, err := p.repo.GetAllSshHosts(dto.Pageable{})
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": hosts})
}

func (p sshclientAPI) AddSshHost(c *gin.Context) {

	var host entities.SshHosts

	if err := c.ShouldBind(&host); err != nil ||
		len(host.Host) == 0 ||
		host.SshKeys.ID == 0 {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}
	touch := entities.Touch{User: "touch", Host: host.Host}
	body, err := kit.PutInt("/sshclient/finger/get", touch)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	err = json.Unmarshal(body, &host.Finger)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	id, err := p.repo.AddSshHost(host)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": id})
}

func (p sshclientAPI) EditSshHost(c *gin.Context) {

	var host entities.SshHosts

	if err := c.ShouldBind(&host); err != nil ||
		len(host.Host) == 0 ||
		host.SshKeys.ID == 0 {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.EditSshHost(host)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p sshclientAPI) DelSshHost(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := p.repo.DelSshHost(id); err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

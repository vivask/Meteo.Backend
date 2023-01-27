package sshclient

import (
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/errors"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p sshclientAPI) GetAllGitUsers(c *gin.Context) {
	users, err := p.repo.GetAllGitUsers(dto.Pageable{})
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": users})
}

func (p sshclientAPI) AddGitUser(c *gin.Context) {

	var user entities.GitUsers

	if err := c.ShouldBind(&user); err != nil ||
		len(user.Username) == 0 ||
		len(user.Password) == 0 ||
		user.SshKeys.ID == 0 {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	id, err := p.repo.AddGitUser(user)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": id})
}

func (p sshclientAPI) EditGitUser(c *gin.Context) {

	var user entities.GitUsers

	if err := c.ShouldBind(&user); err != nil ||
		len(user.Username) == 0 ||
		len(user.Password) == 0 ||
		user.SshKeys.ID == 0 {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.EditGitUser(user)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p sshclientAPI) DelGitUser(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := p.repo.DelGitUser(id); err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p sshclientAPI) GetAllGitServices(c *gin.Context) {
	services, err := p.repo.GetAllGitServices(dto.Pageable{})
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": services})
}

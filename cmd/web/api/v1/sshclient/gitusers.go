package sshclient

import (
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p sshclientAPI) GetAllGitUsers(c *gin.Context) {
	users, err := p.repo.GetAllGitUsers(dto.Pageable{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
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
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.repo.AddGitUser(user)
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

func (p sshclientAPI) EditGitUser(c *gin.Context) {

	var user entities.GitUsers

	if err := c.ShouldBind(&user); err != nil ||
		len(user.Username) == 0 ||
		len(user.Password) == 0 ||
		user.SshKeys.ID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.repo.EditGitUser(user)
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

func (p sshclientAPI) DelGitUser(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	if err := p.repo.DelGitUser(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p sshclientAPI) GetAllGitServices(c *gin.Context) {
	services, err := p.repo.GetAllGitServices(dto.Pageable{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": services})
}

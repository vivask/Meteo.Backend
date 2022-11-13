package sshclient

import (
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p sshclientAPI) GetAllSshKeys(c *gin.Context) {
	keys, err := p.repo.GetAllSshKeys(dto.Pageable{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": keys})
}

func (p sshclientAPI) AddSshKey(c *gin.Context) {

	var key entities.SshKeys

	if err := c.ShouldBind(&key); err != nil ||
		len(key.Owner) == 0 ||
		len(key.Finger) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.repo.AddSshKey(key)
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

func (p sshclientAPI) DelSshKey(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	if err := p.repo.DelSshKey(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

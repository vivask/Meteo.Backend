package radius

import (
	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/errors"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p radiusAPI) GetAllUsers(c *gin.Context) {
	users, err := p.repo.GetAllUsers(dto.Pageable{})
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": users})
}

func (p radiusAPI) AddUser(c *gin.Context) {
	var user entities.Radcheck

	if err := c.ShouldBind(&user); err != nil ||
		len(user.UserName) == 0 ||
		len(user.Value) == 0 ||
		len(user.Attribute) == 0 ||
		len(user.Op) == 0 {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	id, err := p.repo.AddUser(user)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": id})
}

func (p radiusAPI) EditUser(c *gin.Context) {
	var user entities.Radcheck

	if err := c.ShouldBind(&user); err != nil ||
		len(user.UserName) == 0 ||
		len(user.Value) == 0 ||
		len(user.Attribute) == 0 ||
		len(user.Op) == 0 {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.EditUser(user)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p radiusAPI) DeleteUser(c *gin.Context) {
	id, err := utils.StringToUint32(c.Param("id"))
	if err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := p.repo.DelUser(id); err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

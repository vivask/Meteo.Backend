package radius

import (
	"meteo/internal/dto"
	"meteo/internal/errors"
	"meteo/internal/log"
	"meteo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p radiusAPI) GetAllAccounting(c *gin.Context) {
	var page dto.Pageable

	if err := c.ShouldBind(&page); err != nil {
		log.Error(err)
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	accounts, err := p.repo.GetAllAccounting(page)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": accounts})
}

func (p radiusAPI) GetVerifiedAccounting(c *gin.Context) {
	var page dto.Pageable

	if err := c.ShouldBind(&page); err != nil {
		log.Error(err)
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	accounts, err := p.repo.GetVerifiedAccounting(page)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": accounts})
}

func (p radiusAPI) GetNotVerifiedAccounting(c *gin.Context) {
	var page dto.Pageable

	if err := c.ShouldBind(&page); err != nil {
		log.Error(err)
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	accounts, err := p.repo.GetNotVerifiedAccounting(page)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": accounts})
}

func (p radiusAPI) GetAlarmAccounting(c *gin.Context) {
	var page dto.Pageable

	if err := c.ShouldBind(&page); err != nil {
		log.Error(err)
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	accounts, err := p.repo.GetAlarmAccounting(page)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": accounts})
}

func (p radiusAPI) GetAllVerified(c *gin.Context) {

	verified, err := p.repo.GetAllVerified(dto.Pageable{})
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": verified})
}

func (p radiusAPI) Verified(c *gin.Context) {
	id, err := utils.StringToInt(c.Param("id"))
	if err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := p.repo.Verify(id); err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p radiusAPI) ExcludeUser(c *gin.Context) {
	id, err := utils.StringToInt(c.Param("id"))
	if err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := p.repo.ExcludeUser(id); err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p radiusAPI) ClearAccounting(c *gin.Context) {
	id := c.Param("id")

	if err := p.repo.DeleteAccounting(id); err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

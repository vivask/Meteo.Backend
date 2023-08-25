package database

import (
	"meteo/internal/entities"
	"meteo/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p databaseAPI) GetAllSshKeys(c *gin.Context) {
	keys, err := p.repo.GetAllSshKeys()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, keys)
}

func (p databaseAPI) ReplaceSshKeys(c *gin.Context) {

	var keys []entities.SshKeys

	if err := c.ShouldBind(&keys); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceSshKeys(keys)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) GetAllSshHosts(c *gin.Context) {
	keys, err := p.repo.GetAllSshHosts()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, keys)
}

func (p databaseAPI) ReplaceSshHosts(c *gin.Context) {

	var hosts []entities.SshHosts

	if err := c.ShouldBind(&hosts); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceSshHosts(hosts)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) GetAllGitUsers(c *gin.Context) {
	keys, err := p.repo.GetAllGitUsers()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, keys)
}

func (p databaseAPI) ReplaceGitUsers(c *gin.Context) {

	var users []entities.GitUsers

	if err := c.ShouldBind(&users); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceGitUsers(users)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) GetAllHomezones(c *gin.Context) {
	keys, err := p.repo.GetAllHomezones()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, keys)
}

func (p databaseAPI) ReplaceHomezone(c *gin.Context) {

	var hosts []entities.Homezone

	if err := c.ShouldBind(&hosts); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceHomezone(hosts)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) GetAllToVpnManual(c *gin.Context) {
	keys, err := p.repo.GetAllToVpnManual()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, keys)
}

func (p databaseAPI) ReplaceToVpnManual(c *gin.Context) {

	var hosts []entities.ToVpnManual

	if err := c.ShouldBind(&hosts); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceToVpnManual(hosts)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) GetAllToVpnAuto(c *gin.Context) {
	keys, err := p.repo.GetAllToVpnAuto()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, keys)
}

func (p databaseAPI) ReplaceToVpnAuto(c *gin.Context) {

	var hosts []entities.ToVpnAuto

	if err := c.ShouldBind(&hosts); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceToVpnAuto(hosts)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) GetAllToVpnIgnore(c *gin.Context) {
	keys, err := p.repo.GetAllToVpnIgnore()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, keys)
}

func (p databaseAPI) ReplaceToVpnIgnore(c *gin.Context) {

	var hosts []entities.ToVpnIgnore

	if err := c.ShouldBind(&hosts); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceToVpnIgnore(hosts)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) GetAllTasks(c *gin.Context) {
	keys, err := p.repo.GetAllTasks()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, keys)
}

func (p databaseAPI) ReplaceTasks(c *gin.Context) {

	var tasks []entities.Tasks

	if err := c.ShouldBind(&tasks); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceTasks(tasks)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) GetAllJobs(c *gin.Context) {
	keys, err := p.repo.GetAllJobs()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, keys)
}

func (p databaseAPI) ReplaceJobs(c *gin.Context) {

	var jobs []entities.Jobs

	if err := c.ShouldBind(&jobs); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceJobs(jobs)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) GetAllRadcheck(c *gin.Context) {
	keys, err := p.repo.GetAllRadcheck()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, keys)
}

func (p databaseAPI) ReplaceRadcheck(c *gin.Context) {

	var users []entities.Radcheck

	if err := c.ShouldBind(&users); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceRadcheck(users)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) GetAllRadverified(c *gin.Context) {
	keys, err := p.repo.GetAllRadverified()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, keys)
}

func (p databaseAPI) ReplaceRadverified(c *gin.Context) {

	var users []entities.Radverified

	if err := c.ShouldBind(&users); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceRadverified(users)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) GetAllUsers(c *gin.Context) {
	keys, err := p.repo.GetAllUsers()
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, keys)
}

func (p databaseAPI) ReplaceUser(c *gin.Context) {

	var users []entities.User

	if err := c.ShouldBind(&users); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, errors.ErrInvalidInputs))
		return
	}

	err := p.repo.ReplaceUser(users)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

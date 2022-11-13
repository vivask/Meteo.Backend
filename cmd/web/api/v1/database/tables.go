package database

import (
	"meteo/internal/dto"
	"meteo/internal/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p databaseAPI) GetAllTables(c *gin.Context) {
	tables, err := p.repo.GetAllTables(dto.Pageable{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEERR",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": tables})
}

func (p databaseAPI) GetAllSTypes(c *gin.Context) {
	types, err := p.repo.GetAllSTypes(dto.Pageable{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEERR",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": types})
}

func (p databaseAPI) AddTable(c *gin.Context) {

	var table entities.SyncTables

	if err := c.ShouldBind(&table); err != nil ||
		len(table.ID) == 0 ||
		len(table.Params) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "DATABASEERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.repo.AddTable(table)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) EditTable(c *gin.Context) {

	var table entities.SyncTables

	if err := c.ShouldBind(&table); err != nil ||
		len(table.ID) == 0 ||
		len(table.Params) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "DATABASEERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.repo.EditTable(table)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) DelTable(c *gin.Context) {

	if err := p.repo.DelTable(c.Param("id")); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) DelTables(c *gin.Context) {

	var tables []entities.SyncTables

	if err := c.ShouldBind(&tables); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "DATABASEERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	if err := p.repo.DelTables(tables); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p databaseAPI) ImportTableContent(c *gin.Context) {
	if err := p.repo.ImportTableContent(c.Param("id")); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEERR",
				"message": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (p databaseAPI) ImportTablesContent(c *gin.Context) {
	var tables []entities.SyncTables

	if err := c.ShouldBind(&tables); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "DATABASEERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	if err := p.repo.ImportTablesContent(tables); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEERR",
				"message": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (p databaseAPI) SaveTableContent(c *gin.Context) {
	if err := p.repo.SaveTableContent(c.Param("id")); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEERR",
				"message": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (p databaseAPI) SaveTablesContent(c *gin.Context) {
	var tables []entities.SyncTables

	if err := c.ShouldBind(&tables); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "DATABASEERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	if err := p.repo.SaveTablesContent(tables); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "DATABASEERR",
				"message": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

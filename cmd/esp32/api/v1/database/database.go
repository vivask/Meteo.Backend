package database

import (
	repo "meteo/internal/repo/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DatabaseAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
}

type databaseAPI struct {
	repo repo.DatabaseService
}

// NewDatabaseAPI get database service instance
func NewDatabaseAPI(db *gorm.DB) DatabaseAPI {
	return &databaseAPI{repo: repo.NewDatabaseService(db)}
}

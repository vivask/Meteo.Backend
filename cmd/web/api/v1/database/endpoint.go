package database

import "github.com/gin-gonic/gin"

func (p databaseAPI) RegisterAPIV1(router *gin.RouterGroup) {
	database := router.Group("/database")
	database.GET("/tables/get", p.GetAllTables)
	database.POST("/table/add", p.AddTable)
	database.POST("/table/edit", p.EditTable)
	database.DELETE("/table/:id", p.DelTable)
	database.POST("/delete/tables", p.DelTables)
	database.GET("/stypes/get", p.GetAllSTypes)
	database.PUT("/table/import/:id", p.ImportTableContent)
	database.POST("/import/tables", p.ImportTablesContent)
	database.PUT("/table/save/:id", p.SaveTableContent)
	database.POST("/save/tables", p.SaveTablesContent)
}

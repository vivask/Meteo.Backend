package database

import "github.com/gin-gonic/gin"

func (p databaseAPI) RegisterAPIV1(router *gin.RouterGroup) {
	database := router.Group("/database")
	database.GET("/bmx280/get", p.GetNotSyncBmx280)
	database.POST("/bmx280", p.AddSyncBmx280)
	database.PUT("/lock/bmx280", p.LockBmx280)
	database.PUT("/unlock/bmx280", p.UnlockBmx280)
	database.PUT("/sync/bmx280", p.SyncBmx280)
	database.PUT("/replace/bmx280", p.ReplaceBmx280)
	database.PUT("/sync", p.SyncEsp32Tables)

	database.GET("/ds18b20/get", p.GetNotSyncDs18b20)
	database.POST("/ds18b20", p.AddSyncDs18b20)
	database.PUT("/lock/ds18b20", p.LockDs18b20)
	database.PUT("/unlock/ds18b20", p.UnlockDs18b20)
	database.PUT("/sync/ds18b20", p.SyncDs18b20)
	database.PUT("/replace/ds18b20", p.ReplaceDs18b20)

	database.GET("/mics6814/get", p.GetNotSyncMics6814)
	database.POST("/mics6814", p.AddSyncMics6814)
	database.PUT("/lock/mics6814", p.LockMics6814)
	database.PUT("/unlock/mics6814", p.UnlockMics6814)
	database.PUT("/sync/mics6814", p.SyncMics6814)
	database.PUT("/replace/mics6814", p.ReplaceMics6814)

	database.GET("/ze08ch2o/get", p.GetNotSyncZe08ch2o)
	database.POST("/ze08ch2o", p.AddSyncZe08ch2o)
	database.PUT("/lock/ze08ch2o", p.LockZe08ch2o)
	database.PUT("/unlock/ze08ch2o", p.UnlockZe08ch2o)
	database.PUT("/sync/ze08ch2o", p.SyncZe08ch2o)
	database.PUT("/replace/ze08ch2o", p.ReplaceZe08ch2o)

}

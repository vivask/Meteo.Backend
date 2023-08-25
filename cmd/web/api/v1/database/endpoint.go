package database

import "github.com/gin-gonic/gin"

func (p databaseAPI) RegisterAPIV1(router *gin.RouterGroup) {
	database := router.Group("/database")
	database.GET("/tables", p.GetAllTables)
	database.PUT("/tables", p.AddTable)
	database.POST("/tables", p.EditTable)
	database.POST("/tables/delete", p.DelTables)
	database.POST("/tables/drop", p.DropTables)
	database.GET("/stypes", p.GetAllSTypes)
	database.PUT("/import", p.ImportTablesContent)
	database.PUT("/save", p.SaveTablesContent)

	database.PUT("sync/bmx280/:direction", p.SyncBmx280)
	database.PUT("sync/mics6814/:direction", p.SyncMics6814)
	database.PUT("sync/radsens/:direction", p.SyncRadsens)
	database.PUT("sync/ze08ch2o/:direction", p.SyncZe08ch2o)
	database.PUT("sync/ds18b20/:direction", p.SyncDs18b20)

	database.PUT("replace/bmx280/:direction", p.ReplaceBmx280)
	database.PUT("replace/mics6814/:direction", p.ReplaceMics6814)
	database.PUT("replace/radsens/:direction", p.ReplaceRadsens)
	database.PUT("replace/ze08ch2o/:direction", p.ReplaceZe08ch2o)
	database.PUT("replace/ds18b20/:direction", p.ReplaceDs18b20)

	database.PUT("replace/ssh_keys/:direction", p.ReplaceSshKeys)
	database.PUT("replace/ssh_hosts/:direction", p.ReplaceSshHosts)
	database.PUT("replace/git_users/:direction", p.ReplaceGitUsers)
	database.PUT("replace/homezones/:direction", p.ReplaceHomezone)
	database.PUT("replace/tovpn_manuals/:direction", p.ReplaceToVpnManual)
	database.PUT("replace/tovpn_autos/:direction", p.ReplaceToVpnAuto)
	database.PUT("replace/tovpn_ignores/:direction", p.ReplaceToVpnIgnore)
	database.PUT("replace/tasks/:direction", p.ReplaceTasks)
	database.PUT("replace/jobs/:direction", p.ReplaceJobs)
	database.PUT("replace/radcheck/:direction", p.ReplaceRadcheck)
	database.PUT("replace/radverified/:direction", p.ReplaceRadverified)
	database.PUT("replace/users/:direction", p.ReplaceUser)
}

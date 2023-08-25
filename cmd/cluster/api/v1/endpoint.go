package v1

import "github.com/gin-gonic/gin"

func (p clusterAPI) RegisterAPIV1(router *gin.RouterGroup) {
	cluster := router.Group("/cluster")
	cluster.GET("/leader/get", p.IsLeader)
	cluster.GET("/health", p.GetHealth)
	cluster.POST("/database/exec", p.DbExec)
	cluster.GET("/logging", p.GetLogging)
	cluster.PUT("/logging", p.ClearLogging)
	cluster.GET("/logging/empty", p.IsEmptyLog)
	cluster.GET("/main/state", p.GetMainState)
	cluster.GET("/backup/state", p.GetBackupState)
}

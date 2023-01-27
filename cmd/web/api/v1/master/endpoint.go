package master

import "github.com/gin-gonic/gin"

func (p mainAPI) RegisterAPIV1(router *gin.RouterGroup) {
	main := router.Group("/main")
	main.GET("/state", p.GetServicesState)
	main.PUT("/restart/:id", p.RestartServerCont)
	main.PUT("/stop/:id", p.StopServerCont)
	main.PUT("/start/:id", p.StartServerCont)
	main.PUT("/reboot", p.Reboot)
	main.PUT("/shutdown", p.Shutdown)
	main.GET("/logging/:id", p.GetLogging)
	main.PUT("/logging/:id", p.ClearLogging)
}

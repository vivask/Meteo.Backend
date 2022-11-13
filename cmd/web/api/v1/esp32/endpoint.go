package esp32

import "github.com/gin-gonic/gin"

func (p esp32API) RegisterProtectedAPIV1(router *gin.RouterGroup) {
	esp32 := router.Group("/esp32")
	esp32.GET("/status/get", p.GetEsp32Settings)
	esp32.GET("/settings/get", p.GetEsp32Settings)
	esp32.POST("/settings/set", p.SetEsp32Settings)
	esp32.GET("/loging/get", p.GetAllLogging)
	esp32.PUT("/loging/clear", p.JournalClear)
	esp32.POST("/upgrade", p.UpgradeEsp32)
	esp32.GET("/upgrade/status/get", p.GetUpgradeStatus)
	esp32.PUT("/upgrade/terminate", p.TerminateUpgrade)
	esp32.PUT("/reboot", p.SetRebootEsp32)
	esp32.PUT("/ap", p.SetApMode)
	esp32.PUT("/mics6814/co/chk", p.Mics6814CoChk)
	esp32.PUT("/mics6814/no2/chk", p.Mics6814No2Chk)
	esp32.PUT("/mics6814/nh3/chk", p.Mics6814Nh3Chk)
	esp32.PUT("/bme280/temperature/chk", p.Bme280TemperatureChk)
	esp32.PUT("/radsens/static/chk", p.RadsensStaticChk)
	esp32.PUT("/radsens/dynamic/chk", p.RadsensDynamicChk)
	esp32.PUT("/radsens/hv", p.RadsensHVSet)
	esp32.PUT("/ds18b20/temperature/chk", p.Ds18b20TemperatureChk)
	esp32.PUT("/ze08ch2o/ch2o/chk", p.Ze08ch2oChk)
}

func (p esp32API) RegisterPublicAPIV1(router *gin.RouterGroup) {
	esp32 := router.Group("/esp32")
	esp32.GET("/peripheral/get", p.GetHomePageData)
}

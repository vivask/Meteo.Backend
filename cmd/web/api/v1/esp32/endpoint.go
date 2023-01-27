package esp32

import "github.com/gin-gonic/gin"

func (p esp32API) RegisterProtectedAPIV1(router *gin.RouterGroup) {
	esp32 := router.Group("/esp32")
	esp32.GET("/status", p.GetEsp32Settings)
	esp32.GET("/settings", p.GetEsp32Settings)
	esp32.PUT("/settings", p.SetEsp32Settings)
	esp32.GET("/logging", p.GetAllLogging)
	esp32.PUT("/logging", p.JournalClear)
	esp32.PUT("/upgrade", p.UpgradeEsp32)
	esp32.GET("/upgrade/status", p.GetUpgradeStatus)
	esp32.PUT("/upgrade/terminate", p.TerminateUpgrade)
	esp32.PUT("/reboot", p.SetRebootEsp32)
	esp32.PUT("/setup", p.SetApMode)
	esp32.PUT("/mics6814/co", p.Mics6814CoChk)
	esp32.PUT("/mics6814/no2", p.Mics6814No2Chk)
	esp32.PUT("/mics6814/nh3", p.Mics6814Nh3Chk)
	esp32.PUT("/bme280/temperature", p.Bme280TemperatureChk)
	esp32.PUT("/radsens/static", p.RadsensStaticChk)
	esp32.PUT("/radsens/dynamic", p.RadsensDynamicChk)
	esp32.PUT("/radsens/hv", p.RadsensHVSet)
	esp32.PUT("/ds18b20/temperature", p.Ds18b20TemperatureChk)
	esp32.PUT("/ze08ch2o/ch2o", p.Ze08ch2oChk)
}

func (p esp32API) RegisterPublicAPIV1(router *gin.RouterGroup) {
	esp32 := router.Group("/esp32")
	esp32.GET("/peripheral", p.GetHomePageData)
	esp32.POST("/bmx280/min/day", p.GetBmx280MinByHours)
	esp32.POST("/bmx280/max/day", p.GetBmx280MaxByHours)
	esp32.POST("/bmx280/avg/day", p.GetBmx280AvgByHours)
	esp32.POST("/bmx280/min/week", p.GetBmx280MinByDays)
	esp32.POST("/bmx280/max/week", p.GetBmx280MaxByDays)
	esp32.POST("/bmx280/avg/week", p.GetBmx280AvgByDays)
	esp32.POST("/bmx280/min/month", p.GetBmx280MinByDays)
	esp32.POST("/bmx280/max/month", p.GetBmx280MaxByDays)
	esp32.POST("/bmx280/avg/month", p.GetBmx280AvgByDays)
	esp32.POST("/bmx280/min/year", p.GetBmx280MinByMonths)
	esp32.POST("/bmx280/max/year", p.GetBmx280MaxByMonths)
	esp32.POST("/bmx280/avg/year", p.GetBmx280AvgByMonths)

	esp32.POST("/ze08ch2o/min/day", p.GetZe08MinByHours)
	esp32.POST("/ze08ch2o/max/day", p.GetZe08MaxByHours)
	esp32.POST("/ze08ch2o/avg/day", p.GetZe08AvgByHours)
	esp32.POST("/ze08ch2o/min/week", p.GetZe08MinByDays)
	esp32.POST("/ze08ch2o/max/week", p.GetZe08MaxByDays)
	esp32.POST("/ze08ch2o/avg/week", p.GetZe08AvgByDays)
	esp32.POST("/ze08ch2o/min/month", p.GetZe08MinByDays)
	esp32.POST("/ze08ch2o/max/month", p.GetZe08MaxByDays)
	esp32.POST("/ze08ch2o/avg/month", p.GetZe08AvgByDays)
	esp32.POST("/ze08ch2o/min/year", p.GetZe08MinByMonths)
	esp32.POST("/ze08ch2o/max/year", p.GetZe08MaxByMonths)
	esp32.POST("/ze08ch2o/avg/year", p.GetZe08AvgByMonths)

	esp32.POST("/ds18b20/min/day", p.GetDs18b20MinByHours)
	esp32.POST("/ds18b20/max/day", p.GetDs18b20MaxByHours)
	esp32.POST("/ds18b20/avg/day", p.GetDs18b20AvgByHours)
	esp32.POST("/ds18b20/min/week", p.GetDs18b20MinByDays)
	esp32.POST("/ds18b20/max/week", p.GetDs18b20MaxByDays)
	esp32.POST("/ds18b20/avg/week", p.GetDs18b20AvgByDays)
	esp32.POST("/ds18b20/min/month", p.GetDs18b20MinByDays)
	esp32.POST("/ds18b20/max/month", p.GetDs18b20MaxByDays)
	esp32.POST("/ds18b20/avg/month", p.GetDs18b20AvgByDays)
	esp32.POST("/ds18b20/min/year", p.GetDs18b20MinByMonths)
	esp32.POST("/ds18b20/max/year", p.GetDs18b20MaxByMonths)
	esp32.POST("/ds18b20/avg/year", p.GetDs18b20AvgByMonths)

	esp32.POST("/mics6814/min/day", p.GetMics6814MinByHours)
	esp32.POST("/mics6814/max/day", p.GetMics6814MaxByHours)
	esp32.POST("/mics6814/avg/day", p.GetMics6814AvgByHours)
	esp32.POST("/mics6814/min/week", p.GetMics6814MinByDays)
	esp32.POST("/mics6814/max/week", p.GetMics6814MaxByDays)
	esp32.POST("/mics6814/avg/week", p.GetMics6814AvgByDays)
	esp32.POST("/mics6814/min/month", p.GetMics6814MinByDays)
	esp32.POST("/mics6814/max/month", p.GetMics6814MaxByDays)
	esp32.POST("/mics6814/avg/month", p.GetMics6814AvgByDays)
	esp32.POST("/mics6814/min/year", p.GetMics6814MinByMonths)
	esp32.POST("/mics6814/max/year", p.GetMics6814MaxByMonths)
	esp32.POST("/mics6814/avg/year", p.GetMics6814AvgByMonths)

	esp32.POST("/radsens/min/day", p.GetRadsensMinByHours)
	esp32.POST("/radsens/max/day", p.GetRadsensMaxByHours)
	esp32.POST("/radsens/avg/day", p.GetRadsensAvgByHours)
	esp32.POST("/radsens/min/week", p.GetRadsensMinByDays)
	esp32.POST("/radsens/max/week", p.GetRadsensMaxByDays)
	esp32.POST("/radsens/avg/week", p.GetRadsensAvgByDays)
	esp32.POST("/radsens/min/month", p.GetRadsensMinByDays)
	esp32.POST("/radsens/max/month", p.GetRadsensMaxByDays)
	esp32.POST("/radsens/avg/month", p.GetRadsensAvgByDays)
	esp32.POST("/radsens/min/year", p.GetRadsensMinByMonths)
	esp32.POST("/radsens/max/year", p.GetRadsensMaxByMonths)
	esp32.POST("/radsens/avg/year", p.GetRadsensAvgByMonths)
}

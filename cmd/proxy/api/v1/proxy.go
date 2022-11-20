package v1

import (
	"meteo/internal/entities"
	"meteo/internal/kit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p proxyAPI) GetState(c *gin.Context) {

	state := entities.ProxyState{
		Active:  p.IsStarted(),
		Main:    kit.IsMain(),
		AdBlock: p.dns.GetAdBlock(),
		Cache:   p.dns.GetCache(),
		Unlock:  p.dns.GetUnlock(),
	}
	c.JSON(http.StatusOK, state)
}

func (p proxyAPI) SetStart(c *gin.Context) {
	err := p.Start()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "PROXYEER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p proxyAPI) SetStop(c *gin.Context) {
	p.stop()
	c.Status(http.StatusOK)
}

func (p proxyAPI) SetAdBlockOn(c *gin.Context) {
	p.dns.SetAdBlock(true)
	c.Status(http.StatusOK)
}

func (p proxyAPI) SetAdBlockOff(c *gin.Context) {
	p.dns.SetAdBlock(false)
	c.Status(http.StatusOK)
}

func (p proxyAPI) SetCacheOn(c *gin.Context) {
	p.dns.SetCache(true)
	c.Status(http.StatusOK)
}

func (p proxyAPI) SetCacheOff(c *gin.Context) {
	p.dns.SetCache(false)
	c.Status(http.StatusOK)
}

func (p proxyAPI) SetUnlockOn(c *gin.Context) {
	p.dns.SetUnlock(true)
	c.Status(http.StatusOK)
}

func (p proxyAPI) SetUnlockOff(c *gin.Context) {
	p.dns.SetUnlock(false)
	c.Status(http.StatusOK)
}

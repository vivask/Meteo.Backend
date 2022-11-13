package proxy

import (
	"encoding/json"
	"meteo/internal/entities"
	"meteo/internal/kit"
	"meteo/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var state [2]entities.ProxyState

func (p proxyAPI) GetProxyesState(c *gin.Context) {

	body, err := kit.GetInt("/proxy/status/get")
	if err != nil {
		log.Warningf("Local Proxy Server not responding: %v", err)
	} else {
		err = json.Unmarshal(body, &state[0])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	}

	body, err = kit.GetExt("/proxy/status/get")
	if err != nil {
		log.Warningf("Remote Proxy Server not responding: %v", err)
	} else {
		err = json.Unmarshal(body, &state[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": state})
}

func SetProxyStart(c *gin.Context, master bool) {
	if (master && kit.IsMain()) || (!master && !kit.IsMain()) {
		_, err := kit.PutInt("/proxy/server/start", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	} else {
		_, err := kit.PutExt("/proxy/server/start", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	}
	c.Status(http.StatusOK)
}

func SetProxyStop(c *gin.Context, master bool) {
	if (master && kit.IsMain()) || (!master && !kit.IsMain()) {
		_, err := kit.PutInt("/proxy/server/stop", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	} else {
		_, err := kit.PutExt("/proxy/server/stop", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	}
	c.Status(http.StatusOK)
}

func SetProxyAdBlockOn(c *gin.Context, master bool) {
	if (master && kit.IsMain()) || (!master && !kit.IsMain()) {
		_, err := kit.PutInt("/proxy/adblock/on", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	} else {
		_, err := kit.PutExt("/proxy/adblock/on", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	}
	c.Status(http.StatusOK)
}

func SetProxyAdBlockOff(c *gin.Context, master bool) {
	if (master && kit.IsMain()) || (!master && !kit.IsMain()) {
		_, err := kit.PutInt("/proxy/adblock/off", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	} else {
		_, err := kit.PutExt("/proxy/adblock/off", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	}
	c.Status(http.StatusOK)
}

func SetProxyCacheOn(c *gin.Context, master bool) {
	if (master && kit.IsMain()) || (!master && !kit.IsMain()) {
		_, err := kit.PutInt("/proxy/cache/on", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	} else {
		_, err := kit.PutExt("/proxy/cache/on", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	}
	c.Status(http.StatusOK)
}

func SetProxyCacheOff(c *gin.Context, master bool) {
	if (master && kit.IsMain()) || (!master && !kit.IsMain()) {
		_, err := kit.PutInt("/proxy/cache/off", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	} else {
		_, err := kit.PutExt("/proxy/cache/off", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	}
	c.Status(http.StatusOK)
}

func SetProxyUnlockOn(c *gin.Context, master bool) {
	if (master && kit.IsMain()) || (!master && !kit.IsMain()) {
		_, err := kit.PutInt("/proxy/unlock/on", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	} else {
		_, err := kit.PutExt("/proxy/unlock/on", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	}
	c.Status(http.StatusOK)
}

func SetProxyUnlockOff(c *gin.Context, master bool) {
	if (master && kit.IsMain()) || (!master && !kit.IsMain()) {
		_, err := kit.PutInt("/proxy/unlock/off", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	} else {
		_, err := kit.PutExt("/proxy/unlock/off", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code":    http.StatusInternalServerError,
					"error":   "WEBERR",
					"message": err.Error()})
			return
		}
	}
	c.Status(http.StatusOK)
}

func (p proxyAPI) SetMasterProxyStart(c *gin.Context) {
	SetProxyStart(c, true)
}

func (p proxyAPI) SetMasterProxyStop(c *gin.Context) {
	SetProxyStop(c, true)
}

func (p proxyAPI) SetMasterProxyAdBlockOn(c *gin.Context) {
	SetProxyAdBlockOn(c, true)
}

func (p proxyAPI) SetMasterProxyAdBlockOff(c *gin.Context) {
	SetProxyAdBlockOff(c, true)
}

func (p proxyAPI) SetMasterProxyCacheOn(c *gin.Context) {
	SetProxyCacheOn(c, true)
}

func (p proxyAPI) SetMasterProxyCacheOff(c *gin.Context) {
	SetProxyCacheOff(c, true)
}

func (p proxyAPI) SetMasterProxyUnlockOn(c *gin.Context) {
	SetProxyUnlockOn(c, true)
}

func (p proxyAPI) SetMasterProxyUnlockOff(c *gin.Context) {
	SetProxyUnlockOff(c, true)
}

func (p proxyAPI) SetSlaveProxyStart(c *gin.Context) {
	SetProxyStart(c, false)
}

func (p proxyAPI) SetSlaveProxyStop(c *gin.Context) {
	SetProxyStop(c, false)
}

func (p proxyAPI) SetSlaveProxyAdBlockOn(c *gin.Context) {
	SetProxyAdBlockOn(c, false)
}

func (p proxyAPI) SetSlaveProxyAdBlockOff(c *gin.Context) {
	SetProxyAdBlockOff(c, false)
}

func (p proxyAPI) SetSlaveProxyCacheOn(c *gin.Context) {
	SetProxyCacheOn(c, false)
}

func (p proxyAPI) SetSlaveProxyCacheOff(c *gin.Context) {
	SetProxyCacheOff(c, false)
}

func (p proxyAPI) SetSlaveProxyUnlockOn(c *gin.Context) {
	SetProxyUnlockOn(c, false)
}

func (p proxyAPI) SetSlaveProxyUnlockOff(c *gin.Context) {
	SetProxyUnlockOff(c, false)
}

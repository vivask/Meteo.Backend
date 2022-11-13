package v1

import (
	"meteo/cmd/proxy/api/v1/tools"
	"meteo/internal/config"
	"meteo/internal/kit"
	"meteo/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p proxyAPI) AdBlockLoad(c *gin.Context) {
	if kit.IsLeader() {

		list, loaded := tools.UpdateFile(config.Default.Proxy.AdSources)
		p.dns.SetBlackList(list)
		p.checkLoadedBlackLists(loaded)
		_, err := kit.PostExt("/proxy/adblock/update", list.GetData())
		if err != nil {
			log.Errorf("update external blocklist error: %v", err)
		}
	}
	c.Status(http.StatusOK)
}

func (p proxyAPI) AdBlockUpdate(c *gin.Context) {
	data := map[string]struct{}{}

	if err := c.ShouldBind(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "PROXYERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}
	p.dns.SetBlackListData(data)

	c.Status(http.StatusOK)
}

func (p proxyAPI) checkLoadedBlackLists(loaded []bool) {
	for idx, load := range loaded {
		if !load {
			message := "Не удалось загрузить список блокировки: " + config.Default.Proxy.AdSources[idx]
			_, err := kit.PostInt("/messanger/telegram", message)
			if err != nil {
				log.Errorf("can't send telegram message: %v", err)
			}
		}
	}
}

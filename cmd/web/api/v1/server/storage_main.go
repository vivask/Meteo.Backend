package server

import (
	"meteo/internal/kit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p serverAPI) MainStorageMount(c *gin.Context) {

	_, err := kit.PutMain("/server/storage/mount", nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) MainStorageUmount(c *gin.Context) {

	_, err := kit.PutMain("/server/storage/umount", nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) MainStorageRemount(c *gin.Context) {

	_, err := kit.PutMain("/server/storage/remount", nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

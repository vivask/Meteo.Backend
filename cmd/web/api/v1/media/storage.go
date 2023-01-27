package media

import (
	"meteo/internal/errors"
	"meteo/internal/kit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p mediaAPI) StorageMount(c *gin.Context) {

	_, err := kit.PutMain("/media/storage/mount", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mediaAPI) StorageUmount(c *gin.Context) {

	_, err := kit.PutMain("/media/storage/umount", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

func (p mediaAPI) StorageRemount(c *gin.Context) {

	_, err := kit.PutMain("/media/storage/remount", nil)

	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Status(http.StatusOK)
}

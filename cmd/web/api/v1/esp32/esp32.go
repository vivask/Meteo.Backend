package esp32

import (
	"meteo/internal/entities"
	"meteo/internal/log"
	repo "meteo/internal/repo/esp32"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Esp32API interface {
	RegisterProtectedAPIV1(router *gin.RouterGroup)
	RegisterPublicAPIV1(router *gin.RouterGroup)
}

type esp32API struct {
	repo repo.Esp32Service
}

// NewEsp32API get esp32 service instance
func NewEsp32API(db *gorm.DB) Esp32API {
	return &esp32API{repo: repo.NewEsp32Service(db)}
}

func (p esp32API) GetEsp32Settings(c *gin.Context) {
	settings, err := p.repo.GetSettings()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": settings})
}

func (p esp32API) SetEsp32Settings(c *gin.Context) {
	var settings entities.Settings
	if err := c.ShouldBind(&settings); err != nil {
		log.Errorf("unmarshal error: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code":    http.StatusBadRequest,
				"error":   "WEBERR",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	err := p.repo.SetSettings(&settings)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "WEBERR",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

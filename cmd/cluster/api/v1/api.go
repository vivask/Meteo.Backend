package v1

import (
	vip "meteo/cmd/cluster/internal"
	"meteo/internal/entities"
	repo "meteo/internal/repo/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ClusterAPI api interface
type ClusterAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
	IsLeader(*gin.Context)
}

type clusterAPI struct {
	repo repo.DatabaseService
}

// NewClusterAPI get mesanger service instance
func NewClusterAPI(db *gorm.DB) ClusterAPI {
	return &clusterAPI{repo: repo.NewDatabaseService(db)}
}

func (p clusterAPI) IsLeader(c *gin.Context) {
	state := entities.Cluster{
		Leader:      vip.IsLeader(),
		AliveRemote: vip.IsAliveRemote(),
	}
	c.JSON(http.StatusOK, state)
}

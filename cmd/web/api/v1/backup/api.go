package backup

import "github.com/gin-gonic/gin"

type BackupAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
}

type backupAPI struct {
}

// NewBackupAPI get backup server instance
func NewBackupAPI() BackupAPI {
	return &backupAPI{}
}

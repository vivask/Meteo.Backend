package nut

import (
	"fmt"
	"meteo/internal/config"
	"meteo/internal/utils"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func (p nutAPI) GetState(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": nil})
}

func ReadUpsDriver() (string, error) {
	cmd := fmt.Sprintf("upsc %s@localhost:3493", config.Default.Nut.Driver)
	shell := utils.NewShell(cmd)
	err, out, _ := shell.Run(1)
	if err != nil {
		return out, fmt.Errorf("upsdriver health error: %w", err)
	}
	matched, _ := regexp.MatchString("ups.test.result: Done and passed", out)
	if !matched {
		return out, fmt.Errorf("upsdriver not healthy: %s", out)
	}
	return out, nil
}

package utils

import (
	"fmt"
	"meteo/internal/config"
)

func GetCurrentApi() string {
	return fmt.Sprintf("/api/v%s", config.Default.App.Version)
}

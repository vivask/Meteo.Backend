package nut

import "github.com/gin-gonic/gin"

type NutAPI interface {
	RegisterAPIV1(router *gin.RouterGroup)
}

type nutAPI struct {
}

// NewNutAPI get nut service instance
func NewNutAPI() NutAPI {
	return &nutAPI{}
}

package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRouterAPIV1 group for api/v1/*
func RegisterRouterAPIV1(router *gin.RouterGroup, db *gorm.DB) {
	userAPI := NewUserAPI(db)
	router.GET("/users", userAPI.GetAllUser)
	router.GET("/users/:id", userAPI.GetUser)

	prodAPI := NewProductAPI(db)
	router.GET("/products", prodAPI.FindAll)
	router.GET("/products/:id", prodAPI.GetProduct)
	router.DELETE("/products/:id", prodAPI.DeleteProduct)
}

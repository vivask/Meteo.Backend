package v1

import (
	"meteo/internal/config"
	"meteo/internal/log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func (p authAPI) LinkJWT() {
	p.midleware = p.jwtMidleware()
	errInit := p.midleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("midleware.MiddlewareInit() Error:" + errInit.Error())
	}
	p.public.POST("/api/v1/signup", p.Signup)
	p.public.POST("/api/v1/login", p.midleware.LoginHandler)
	p.public.NoRoute(p.midleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Infof("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// Refresh time can be longer than token timeout
	p.protected.GET("/refresh_token", p.midleware.RefreshHandler)
	p.protected.Use(p.midleware.MiddlewareFunc())
	p.protected.GET("/logout", p.midleware.LogoutHandler)
	p.protected.GET("/user", p.CurrentUser)
}

func (p authAPI) jwtMidleware() *jwt.GinJWTMiddleware {

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte(config.Default.Auth.JwtKey),
		Timeout:         time.Duration(config.Default.Auth.JwtExpiration) * time.Hour,
		MaxRefresh:      time.Duration(config.Default.Auth.JwtRefreshExpiration) * time.Hour,
		PrivKeyFile:     config.Default.Auth.AccessTokenPrivateKeyPath,
		PubKeyFile:      config.Default.Auth.AccessTokenPublicKeyPath,
		IdentityKey:     identityKey,
		PayloadFunc:     p.Payload,
		IdentityHandler: p.Identity,
		Authenticator:   p.Authenticator,
		Authorizator:    p.Authorizator,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}

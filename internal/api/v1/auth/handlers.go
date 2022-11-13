package v1

import (
	"fmt"
	"meteo/internal/entities"

	"meteo/internal/log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func (p authAPI) Authenticator(c *gin.Context) (interface{}, error) {
	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userID := loginVals.Username
	password := loginVals.Password

	user, err := p.repo.GetUserByName(userID)
	if err != nil {
		return "", fmt.Errorf("get user error: %w", err)
	}

	if userID == user.Username && doPasswordsMatch(user.Password, password) {
		return &entities.User{
			Username: userID,
		}, nil
	}

	log.Error("Authentication fail")
	return nil, jwt.ErrFailedAuthentication
}

func (p authAPI) Authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*entities.User); ok {
		user, err := p.repo.GetUserByName(v.Username)
		if err != nil {
			log.Error(err)
			return false
		}
		if v.Username == user.Username {
			return true
		}
	}
	return false
}

func (p authAPI) Payload(data interface{}) jwt.MapClaims {
	if v, ok := data.(*entities.User); ok {
		return jwt.MapClaims{
			identityKey: v.Username,
		}
	}
	return jwt.MapClaims{}
}

func (p authAPI) Identity(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	currentUser := &entities.User{
		Username: claims[identityKey].(string),
	}
	//log.Warningf("Identity: %v", p.currentUser)
	return currentUser
}

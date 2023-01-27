package errors

import (
	"meteo/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ErrReplyUnknown  = "Unknown error"
	ErrInvalidInputs = "Invalid inputs. Please check your inputs"
)

const tagUnhandlerError = "[UnhandlerError]:"
const tagAppError = "[AppError]:"

// GinError middleware
func GinError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if errors := c.Errors.ByType(gin.ErrorTypeAny); len(errors) > 0 {
			err := errors[0].Err
			if err, ok := err.(*Error); ok {
				log.Error(tagAppError, err)
				c.AbortWithStatusJSON(err.Code, gin.H{"code": err.Code, "message": err.Error()})
				return
			}
			log.Error(tagUnhandlerError, err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": ErrReplyUnknown})
			return
		}
	}
}

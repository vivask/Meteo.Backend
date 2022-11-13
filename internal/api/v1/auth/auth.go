package v1

import (
	"meteo/internal/entities"
	repo "meteo/internal/repo/auth"
	"meteo/internal/utils"
	"net/http"
	"strings"

	"meteo/internal/log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProxyAPI api controller of produces
type AuthAPI interface {
	LinkJWT()
	Signup(*gin.Context)
	CurrentUser(*gin.Context)
	Protected() *gin.RouterGroup
}

type authAPI struct {
	repo      repo.AuthService
	public    *gin.Engine
	protected *gin.RouterGroup
	midleware *jwt.GinJWTMiddleware
}

type signup struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"email"`
	Password string `form:"password" json:"password" binding:"required"`
}

// NewAuthAPI get product service instance
func NewAuthAPI(public *gin.Engine, db *gorm.DB) AuthAPI {
	return &authAPI{
		repo:      repo.NewAuthService(db),
		public:    public,
		protected: public.Group("/api/v1/admin"),
	}
}

func (p authAPI) Signup(c *gin.Context) {

	var signVals signup

	if err := c.ShouldBind(&signVals); err != nil ||
		len(signVals.Username) == 0 ||
		len(signVals.Password) == 0 ||
		len(signVals.Email) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "VALIDATEERR-1",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	user := entities.User{
		Username: signVals.Username,
		Email:    signVals.Email,
		Password: signVals.Password,
	}

	hashedPass, err := hashPassword(user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error":   "VALIDATEERR-1",
				"message": "Error generating password hash"})
		return
	}
	user.Password = hashedPass
	user.Token = utils.GenerateRandomString(15)

	err = p.repo.Create(user)
	if err != nil {
		log.Errorf("unable create user, error: %v", err)
		errMsg := err.Error()
		if strings.Contains(errMsg, PgDuplicateKeyMsg) {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error":   "VALIDATEERR-1",
					"message": PgDuplicateKeyMsg})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"error":   "VALIDATEERR-1",
					"message": errMsg})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

func (p authAPI) CurrentUser(c *gin.Context) {

	claims := jwt.ExtractClaims(c)
	currentUser := &entities.User{
		Username: claims[identityKey].(string),
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": currentUser})
}

func (p authAPI) Protected() *gin.RouterGroup {
	return p.protected
}

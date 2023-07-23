package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/vashu992/Home-Automation/model"
	"github.com/vashu992/Home-Automation/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Middleware function for token authentication
func (api APIRoutes) AuthMiddlewareComplete() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader(model.Authorization)
		if tokenString == "" {
			util.Log(model.LogLevelInfo, model.ServerPackageLavel,
				model.AuthMiddleware, "token string empty", nil)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return model.SecretKey, nil
		})
		if err != nil || !token.Valid {
			util.Log(model.LogLevelError, model.ServerPackageLavel,
				model.AuthMiddleware, "token value is not valid", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}
		util.Log(model.LogLevelInfo, model.ServerPackageLavel,
			model.AuthMiddleware, "token value parsed", token)
		
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			email := claims["email"].(string)
			password := claims["password"].(string)
			util.Log(model.LogLevelInfo, model.ServerPackageLavel,
				model.AuthMiddleware, "token value for email and password", email+" password = "+password)
			db, err := gorm.Open(postgres.Open(model.DSN), &gorm.Config{})
			defer func() {
				sqldb, err := db.DB()
				if err != nil {
					util.Log(model.LogLevelError, model.ServerPackageLavel,
						model.AuthMiddleware, "error in geting sql object for middleware", nil)
				}
				sqldb.Close()
				util.Log(model.LogLevelInfo, model.ServerPackageLavel,
					model.AuthMiddleware, "middleware db connection closed", nil)
			}()
			if err != nil {
				util.Log(model.LogLevelError, model.PgressPackageLevel, model.NewStore,
					"error while connecting database", err)
				panic(err)
			}
			var user model.User
			util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.AuthMiddlewareComplete,
				"reading user data from db based on email", email+" pass = "+password)
			resp := db.Where("email = ? AND password = ?", email, password).First(&user)
			if resp.Error != nil {
				util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetUser,
					"error while reading user data", resp.Error)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "error in geting user data from Database"})
				c.Abort()
				return
			}
			util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetUser,
				"returning user data", user)
		}
		c.Next()
	}
}
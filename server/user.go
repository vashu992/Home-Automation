package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/vashu992/Home-Automation/model"
	"github.com/vashu992/Home-Automation/util"
)

func (server *Server) GetUsers(c *gin.Context) (*[]model.User, error) {

	// validation is to be done here
	// DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetUsers, "reading all data ", nil)
	users, err := server.Pgress.GetUsers()
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetUsers,
			"error while reading user data from pgress ", err)
		return users, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetUsers,
		"returning all user data to api and setting response ", users)
	c.JSON(http.StatusOK, users)
	return users, nil
}

func (server *Server) GetUsersByFilter(c *gin.Context) (*[]model.User, error) {

	// validation is to be done here
	// DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetUsers, "reading all user data ", nil)
	condition := server.readQueryParams(c)
	users, err := server.Pgress.GetUsersByFilter(condition)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetUsers,
			"error while reading user data from pgress", err)
		return users, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetUsers,
		"returning all user data to api and setting response ", users)
	c.JSON(http.StatusOK, users)
	return users, nil
}

func (server *Server) GetUser(c *gin.Context) (*model.User, error) {

	// validation is to be done here
	// DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetUser,
		"reading user data from pgress ", nil)
	user, err := server.Pgress.GetUser(c.Param("id"))
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetUser,
			"error while reading user data from pgress ", err)
		return user, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetUser,
		"returning user data to api and setting response ", user)
	c.JSON(http.StatusOK, user)
	return user, nil
}

// Sighup API handler
func (server *Server) SighUp(c *gin.Context) {
	var user model.User

	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.CreateUser,
		"unmarshaling user data ", nil)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now().UTC()
	err := server.Pgress.SignUp(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.SignUp,
			"error in saving user record ", user)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to SighUp User "})
		return
	}

	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		model.Email:    user.Email,
		model.Password: user.Password,
		model.UserID:   user.ID,
		model.Expire:   time.Now().Add(model.TokenExpiration).Unix(), // token expiration time
		// Additional data can be added here
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(model.SecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token "})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token ": tokenString})
}

// SighIn API handler
func (server *Server) SighIn(c *gin.Context) {
	var user model.UserSignIn
	err := c.ShouldBindJSON(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateUser,
			"error while unmarshaling payload ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user Data from payload "})
		return
	}

	userResp, err := server.Pgress.SingIn(user)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.SignIn,
			"error getting user data from pgress for emailID ", user.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user data gor given user "})
		return
	}
	if userResp.Email != user.Email || userResp.Password != user.Password {
		util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.SignIn,
			"user data not matched , database response ", userResp)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate user data"})
		return

	}

	// Create a new token
	newtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		model.Email:    user.Email,
		model.Password: user.Password,
		model.UserID:   userResp.ID,
		model.Expire:   time.Now().Add(model.TokenExpiration).Unix(), // Token expiration time
		// Additional data can be added here
	})

	// Sigh the newtoken with the secret key
	tokenString, err := newtoken.SignedString(model.SecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (server *Server) CreateUser(c *gin.Context) error {

	var user model.User
	// Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.CreateUser,
		"unmarshaling user data ", nil)

	err := c.ShouldBindJSON(user)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateUser,
			"error while unmarshaling payload ", err)
		return fmt.Errorf("")
	}
	user.CreatedAt = time.Now().UTC()
	user.ID = uuid.New()
	// validation is to be done here
	// DB call
	err = server.Pgress.CreateUser(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateUser,
			"error while creating record from pgress ", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.CreateUser,
		"successfully created user record and setting response ", user)
	c.JSON(http.StatusCreated, user)
	return nil

}

func (server *Server) UpdateUser(c *gin.Context) error {

	var user model.User
	// Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.UpdateUser,
		"unmarshaling user data ", nil)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdateUser,
			"error while unmarshaling payload ", err)
		return fmt.Errorf("")
	}
	// validation is to be done here
	// DB call
	user.UpdatedAt = time.Now().UTC()
	err = server.Pgress.UpdateUser(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdateUser,
			"error while updating record from pgress ", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.UpdateUser,
		"successfully updated user record and setting response ", user)
	c.JSON(http.StatusOK, user)
	return nil

}

func (server *Server) DeleteUser(c *gin.Context) error {

	// validation is to be done here
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.DeleteUser,
		"reading user id ", nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeleteUser,
			"missing user id ", nil)
		return fmt.Errorf("")
	}
	// DB call
	err := server.Pgress.DeleteUser(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeleteUser,
			"error while deleting user record from pgress ", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.DeleteUser,
		"successfully deleted user record ", nil)
	return nil
}

package server

import (
	"github.com/GoSome/fileUpdater/pkg/core"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"html"
	"net/http"
	"strings"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a *App) PostRegister(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := core.User{}

	u.Username = input.Username
	u.Password = html.EscapeString(strings.TrimSpace(input.Password))

	//_, err := u.SaveUser()

	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	temp, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	//c.JSON(http.StatusOK, gin.H{"message": "registration success"})
	c.JSON(http.StatusOK, gin.H{"User": u, "password": temp})
}

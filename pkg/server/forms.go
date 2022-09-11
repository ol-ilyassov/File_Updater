package server

import (
	"github.com/GoSome/fileUpdater/pkg/config"
	"github.com/GoSome/fileUpdater/pkg/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseConfig struct {
	ServerPort string             `json:"server_port" binding:"required"`
	ServerHost string             `json:"server_host" binding:"required"`
	DisableUi  bool               `json:"disable_ui"`
	Updaters   []core.FileUpdater `json:"updaters"`
}

// Aimed for Self Config

func (a *App) GetContentForm(c *gin.Context) {
	config.Parse(false)
	response := BaseConfig{
		ServerPort: config.Config.ServerPort,
		ServerHost: config.Config.ServerHost,
		DisableUi:  config.Config.DisableUI,
		Updaters:   config.Config.FileUpdaters,
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}

func (a *App) UpdateFileForm(c *gin.Context) {
	return
}

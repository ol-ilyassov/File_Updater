package server

import (
	"encoding/json"
	"fmt"
	"github.com/GoSome/fileUpdater/pkg/config"
	"github.com/GoSome/fileUpdater/pkg/core"
	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
	var req BaseConfig
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "response": err.Error()})
		return
	}
	result, err := yaml.Marshal(req)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	getConfig, _ := c.Get("cfg")
	cfg := getConfig.(core.ServerConfigs)
	updater := cfg.GetUpdaterByName("selfConfig")
	if updater == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "response": "Not Found"})
		return
	}

	f := strings.NewReader(string(result))
	if err := updater.UpdateFile(f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "response": "Config updated successfully"})
	return
}

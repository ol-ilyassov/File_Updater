/*
@Description: just go
@Author: skipper
@Date: 2020/1/13
@Time: 5:13 PM
@ProjectName fileUpdater
*/
package server

import (
	"github.com/GoSome/fileUpdater/pkg/middleware"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/GoSome/fileUpdater/pkg/config"
	"github.com/GoSome/fileUpdater/pkg/core"
	"github.com/gin-gonic/gin"
)

func Run(cfg core.ServerConfigs) {
	app := App{
		Options: cfg,
		Engine:  gin.Default(),
	}

	cfg.RunProcess()
	app.Engine.Use(config.Inject)
	app.Engine.Use(middleware.CORSMiddleware())
	//app.Engine.POST("/api/register", app.PostRegister)
	app.Engine.POST("/api/login", app.PostLogin)

	app.Engine.GET("/api/updaters", app.GetUpdaters)
	app.Engine.GET("/api/updater", app.GetUpdater)
	app.Engine.GET("/api/content", app.GetContent)
	app.Engine.POST("/api/content", app.UpdateFile)
	app.Engine.GET("/api/form", app.GetContentForm)
	app.Engine.POST("/api/form", app.UpdateFileForm)
	app.Engine.POST("/api/exec", app.Exec)

	protected := app.Engine.Group("/protected")
	protected.Use(middleware.JwtAuthMiddleware())

	protected.GET("/user", app.CurrentUser)
	//protected.GET("/api/form", app.GetContentForm)
	//protected.POST("/api/form", app.UpdateFileForm)

	if !cfg.DisableUI {
		sub, err := fs.Sub(app.Options.HttpData, "build/static")
		if err != nil {
			panic(err)
		}
		app.Engine.StaticFS("/static/", http.FS(sub))
		app.Engine.GET("/", app.Index)
		app.Engine.NoRoute(app.Index)
	}
	log.Fatal(app.Engine.Run(cfg.ServerHost + ":" + cfg.ServerPort))
}

func (a *App) Index(c *gin.Context) {
	f := a.Options.HttpData
	path := c.Request.URL.Path
	fileName := "build" + path
	if path == "/" {
		c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		fileName = "build/index.html"
	}
	indexFile, err := f.Open(fileName)
	//use frontend route
	if err != nil {
		c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		fileName = "build/index.html"
		indexFile, _ = f.Open(fileName)
		c.Status(http.StatusOK)
		io.Copy(c.Writer, indexFile)
		return
	}
	c.Status(http.StatusOK)
	io.Copy(c.Writer, indexFile)
}

package server

import (
	"github.com/GoSome/fileUpdater/pkg/core"
	"github.com/gin-gonic/gin"
)

type App struct {
	Options core.ServerConfigs
	Engine  *gin.Engine
}

// DB Connection

// var DB *gorm.DB

//func ConnectDB() {
//	err := godotenv.Load(".env")
//	if err != nil {
//		log.Fatalf("Error loading .env file")
//	}
//
//	DBdriver := os.Getenv("DB_DRIVER")
//	DBhost := os.Getenv("DB_HOST")
//	DBuser := os.Getenv("DB_USER")
//	DBpassword := os.Getenv("DB_PASSWORD")
//	DBname := os.Getenv("DB_NAME")
//	DBport := os.Getenv("DB_PORT")
//
//	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBuser, DBpassword, DBhost, DBport, DBname)
//
//	DB, err = gorm.Open(DBdriver, DBURL)
//
//	if err != nil {
//		fmt.Println("> DB Connection failure", DBdriver)
//		log.Fatal("Connection error:", err)
//	} else {
//		fmt.Println("> DB Connection success", DBdriver)
//	}
//
//	DB.AutoMigrate(&User{})
//}

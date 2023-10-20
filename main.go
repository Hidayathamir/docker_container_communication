package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	var db *gorm.DB
	var err error
	dsn := "myappuser:myapppassword@tcp(myapp-mysql:3306)/myappdatabase?parseTime=true"
	gormConfig := gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	for {
		db, err = gorm.Open(mysql.Open(dsn), &gormConfig)
		if err == nil {
			break
		}
		log.Warn(err)
		time.Sleep(5 * time.Second)
	}

	router := gin.Default()

	router.GET("/name", func(ctx *gin.Context) {
		type result struct {
			Name string
		}
		r := result{}

		db.Raw("select 'hidayat' as name").First(&r)

		ctx.JSON(http.StatusOK, gin.H{
			"name": r.Name,
		})
	})

	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Error(err)
	}
}

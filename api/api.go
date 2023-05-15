package api

import (
	"fmt"
	"log"
	"time"

	"github.com/MyriadFlow/airbot/api/controllers"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		controllers.ApplyRoutes(api)
	}
}
func Init() {
	ginApp := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	ginApp.Use(cors.New(config))

	ginApp.Use(helmet.Default())

	ginApp.Use(func(ctx *gin.Context) {
		ctx.Set("cache", cache.New(60*time.Minute, 10*time.Minute))
		ctx.Next()
	})
	ApplyRoutes(ginApp)
	err := ginApp.Run(":8080")
	if err != nil {
		log.Fatal("Failed to Start HTTP Server: ", err)
	} else {
		fmt.Println("api is up")
	}
}

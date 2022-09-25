package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/rayyanhunerkar/gin-tryout/docs"
	"github.com/rayyanhunerkar/gin-tryout/pkg/books"
	"github.com/rayyanhunerkar/gin-tryout/pkg/common/db"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Gin Tryout
// @version         1.0
// @description     Gin Tryout

// @host      localhost:3000
// @BasePath  /
func main() {

	port := ":3000"
	dbUrl := "postgres://postgres@localhost:5432/gin-tryout"

	router := gin.Default()
	dbHandler := db.Init(dbUrl)

	books.RegisterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	router.GET("/health", HealthCheck)

	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(port)
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}

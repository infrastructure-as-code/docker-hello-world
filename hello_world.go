package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
)

func serviceInfoMiddleware() gin.HandlerFunc {
	var hostname, err = os.Hostname()
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Hostname", hostname)
		c.Next()
	}
}

func getVersion() string {
	return "$Id$"
}

func helloFunc(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func healthFunc(c *gin.Context) {
	c.String(http.StatusOK, "")
}

func versionFunc(c *gin.Context) {
	c.String(http.StatusOK, getVersion())
}

func setupRouter(routePrefix string) *gin.Engine {
	router := gin.Default()
	ginprom := ginprometheus.NewPrometheus("gin")
	ginprom.Use(router)
	router.Use(serviceInfoMiddleware())
	router.GET("/health", healthFunc)

	rg := router.Group(routePrefix)
	rg.GET("/", helloFunc)
	rg.GET("/version", versionFunc)
	return router
}

func main() {
	optRoutePrefix := flag.String("route-prefix", "/", "Route prefix")
	flag.Parse()

	router := setupRouter(*optRoutePrefix)
	router.Run()
}

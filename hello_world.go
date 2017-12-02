package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
)

var hostname = getHostname()

func getHostname() string {
	var name, err = os.Hostname()
	if err != nil {
		panic(err)
	}
	return name
}

func helloFunc(c *gin.Context) {
	c.Writer.Header().Set("X-Hostname", hostname)
	c.String(http.StatusOK, "Hello, World!")
}

func healthFunc(c *gin.Context) {
	c.Writer.Header().Set("X-Hostname", hostname)
	c.String(http.StatusOK, "")
}

func setupRouter(routePrefix string) *gin.Engine {
	router := gin.Default()
	ginprom := ginprometheus.NewPrometheus("gin")
	ginprom.Use(router)
	router.GET("/health", healthFunc)

	rg := router.Group(routePrefix)
	rg.GET("/", helloFunc)
	return router
}

func main() {
	optRoutePrefix := flag.String("route-prefix", "/", "Route prefix")
	flag.Parse()

	router := setupRouter(*optRoutePrefix)
	router.Run()
}

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
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
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	})
	router.GET("/health", healthFunc)

	rg := router.Group(routePrefix)
	rg.GET("/", helloFunc)
	rg.GET("/version", versionFunc)
	return router
}

func main() {
	optRoutePrefix := flag.String("route-prefix", "/", "Route prefix")
	optVersion := flag.Bool("version", false, "Print version")
	flag.Parse()

	if *optVersion {
		fmt.Fprintf(os.Stderr, "Version: %s\nArch: %s\n",
			getVersion(),
			runtime.GOARCH,
		)
	} else {
		router := setupRouter(*optRoutePrefix)
		_ = router.Run()
	}
}

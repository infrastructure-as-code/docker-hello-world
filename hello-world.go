// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// A minimal example of how to include Prometheus instrumentation.
package main

import (
	"net/http"
  "os"

  "github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

func main() {
  router := gin.Default()
  router.GET("/", helloFunc)
  router.GET("/health", healthFunc)
  router.GET("/metrics", gin.WrapH(promhttp.Handler()))
  router.Run()
}

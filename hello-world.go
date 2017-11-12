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
	"flag"
  "fmt"
	"log"
	"net/http"
  "os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
var hostname = getHostname()

func getHostname() string {
  var name, err = os.Hostname()
  if err != nil {
    panic(err)
  }
  return name
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("X-Hostname", hostname)
  fmt.Fprintf(w, "Hello, World!\n")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("X-Hostname", hostname)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}

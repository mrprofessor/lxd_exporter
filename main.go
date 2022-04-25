package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	version = "staging-UNVERSIONED"

	port = kingpin.Flag(
		"port", "Provide the port to listen on").Default("9666").Int16()
)

func main() {
	logger := log.New(os.Stderr, "lxd_exporter: ", log.LstdFlags)

	kingpin.Version(version)
	kingpin.Parse()

	http.Handle("/metrics", promhttp.Handler())

	servingPort := fmt.Sprintf(":%d", *port)
	logger.Print("Server listening on ", servingPort)
	logger.Fatal(http.ListenAndServe(servingPort, nil))
}

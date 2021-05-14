package util

import (
    "net/http"

    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitUtil() {
    initConfig()
    initHttpClient()
    initLog()
    initMetrics()
}

func initMetrics() {
    go func() {
        // create a new mux server
        server := http.NewServeMux()
        // register a new handler for the /metrics endpoint
        server.Handle("/metrics", promhttp.Handler())
        // start an http server using the mux server
        _ = http.ListenAndServe(":9002", server)
    }()
}

package main

import (
	"base_go_web/cmd/api/config"
	"fmt"
	"golang.org/x/net/http2"
	"net/http"
)

func main() {

	router := config.InitRouter(true)
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", 30001),
		Handler: router,
	}
	http2.ConfigureServer(server, &http2.Server{})
	server.ListenAndServe()

}

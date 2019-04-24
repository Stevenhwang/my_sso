package main

import (
	"net/http"
	"time"

	"niubility_sso/router"
)

func main() {
	router.SetRouter()
	s := &http.Server{
		Addr:         ":1234",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	router.Server.Logger.Fatal(router.Server.StartServer(s))
}

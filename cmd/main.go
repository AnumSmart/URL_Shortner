package main

import (
	"fmt"
	"net/http"
	"os"
	"server/configs"
	"server/internal/auth"
	"server/internal/link"
	"server/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()

	//Handler
	link.NewLinkHandler(router, link.LinkHandlerDep{})

	auth.NewAuthHandler(router, auth.AuthHandlerDep{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("server is listening on port 8081")
	dir, _ := os.Getwd()
	fmt.Println("Current directory:", dir)
	server.ListenAndServe()

}

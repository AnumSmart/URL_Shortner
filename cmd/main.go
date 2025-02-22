package main

import (
	"fmt"
	"net/http"
	"os"
	"server/configs"
	"server/internal/auth"
	"server/internal/link"
	"server/internal/user"
	"server/pkg/db"
	"server/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)

	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDep{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDep{LinkRepository: linkRepository})

	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("server is listening on port 8081")
	dir, _ := os.Getwd()
	fmt.Println("Current directory:", dir)
	server.ListenAndServe()

}

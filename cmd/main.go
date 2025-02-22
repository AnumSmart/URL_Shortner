package main

import (
	"fmt"
	"net/http"
	"os"
	"server/configs"
	"server/internal/auth"
	"server/internal/link"
	"server/internal/stat"
	"server/internal/user"
	"server/pkg/db"
	"server/pkg/event"
	"server/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	//Repositories-----------------------------------------
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	//Services---------------------------------------------
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	//Handlers----------------------------------------------
	auth.NewAuthHandler(router, auth.AuthHandlerDep{
		Config:      conf,
		AuthService: authService,
	})

	link.NewLinkHandler(router, link.LinkHandlerDep{
		LinkRepository: linkRepository,
		Config:         conf,
		//StatRepository: statRepository,
		EventBus: eventBus,
	})

	stat.NewStatHandler(router, stat.StatHandlerDep{
		StatRepository: statRepository,
		Config:         conf,
	})

	//Middlewares-------------------------------------------
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	//Server------------------------------------------------
	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	go statService.AddClick()

	fmt.Println("server is listening on port 8081")
	dir, _ := os.Getwd()
	fmt.Println("Current directory:", dir)
	server.ListenAndServe()

}

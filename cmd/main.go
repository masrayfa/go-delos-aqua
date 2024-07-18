package main

import (
	"log"
	"net/http"

	"github.com/masrayfa/go-delos-aqua/database"
	"github.com/masrayfa/go-delos-aqua/internals/controller"
	"github.com/masrayfa/go-delos-aqua/internals/repository"
	"github.com/masrayfa/go-delos-aqua/internals/service"
)

func main() {

	dbPool := database.NewDBPool()

	// repository
	userRepository := repository.NewUserRepository()
	farmsRepository := repository.NewFarmRepository()
	pondsRepository := repository.NewPondsRepository()

	// service
	userService := service.NewUserService(userRepository, dbPool)
	farmService := service.NewFarmService(farmsRepository, dbPool)
	pondsService := service.NewPondsService(pondsRepository, dbPool)

	// controller
	userController := controller.NewUserController(userService)
	farmController := controller.NewFarmController(farmService)
	pondsController := controller.NewPondsController(pondsService)

	// router
	mainRouter := NewRouter()

	userRouter := NewUserRouter(userController)
	farmRouter := NewFarmRouter(farmController)
	pondsRouter := NewPondsRouter(pondsController)

	// user router
	mainRouter.appRouter.Handler("GET", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))
	mainRouter.appRouter.Handler("POST", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))
	mainRouter.appRouter.Handler("PUT", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))
	mainRouter.appRouter.Handler("DELETE", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))

	// farm router
	mainRouter.appRouter.Handler("GET", "/api/v1/farm/*path", http.StripPrefix("/api/v1/farm", farmRouter))
	mainRouter.appRouter.Handler("POST", "/api/v1/farm/*path", http.StripPrefix("/api/v1/farm", farmRouter))
	mainRouter.appRouter.Handler("PUT", "/api/v1/farm/*path", http.StripPrefix("/api/v1/farm", farmRouter))
	mainRouter.appRouter.Handler("DELETE", "/api/v1/farm/*path", http.StripPrefix("/api/v1/farm", farmRouter))

	// ponds router
	mainRouter.appRouter.Handler("GET", "/api/v1/ponds/*path", http.StripPrefix("/api/v1/ponds", pondsRouter))
	mainRouter.appRouter.Handler("POST", "/api/v1/ponds/*path", http.StripPrefix("/api/v1/ponds", pondsRouter))
	mainRouter.appRouter.Handler("PUT", "/api/v1/ponds/*path", http.StripPrefix("/api/v1/ponds", pondsRouter))
	mainRouter.appRouter.Handler("DELETE", "/api/v1/ponds/*path", http.StripPrefix("/api/v1/ponds", pondsRouter))

	server := http.Server {
		Addr: ":8080",
		Handler: mainRouter.appRouter,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	log.Println("Server is running on port 8080")
}
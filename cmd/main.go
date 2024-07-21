package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/masrayfa/go-delos-aqua/database"
	"github.com/masrayfa/go-delos-aqua/internals/controller"
	"github.com/masrayfa/go-delos-aqua/internals/dependencies"
	"github.com/masrayfa/go-delos-aqua/internals/helper"
	"github.com/masrayfa/go-delos-aqua/internals/middleware"
	"github.com/masrayfa/go-delos-aqua/internals/repository"
	"github.com/masrayfa/go-delos-aqua/internals/service"
)

func main() {
	dbPool := database.NewDBPool()

	validate := dependencies.NewValidator()

	// repository
	userRepository := repository.NewUserRepository()
	farmsRepository := repository.NewFarmRepository()
	pondsRepository := repository.NewPondsRepository()

	// service
	userService := service.NewUserService(userRepository, dbPool, validate)
	farmService := service.NewFarmService(farmsRepository, dbPool, validate)
	pondsService := service.NewPondsService(pondsRepository, dbPool, validate)

	// controller
	userController := controller.NewUserController(userService)
	farmController := controller.NewFarmController(farmService)
	pondsController := controller.NewPondsController(pondsService)

	// router
	mainRouter := NewRouter()

	userRouter := NewUserRouter(userController)
	farmRouter := NewFarmRouter(farmController)
	pondsRouter := NewPondsRouter(pondsController)

	// middleware
	stats := middleware.NewStats()

	// user router
	mainRouter.appRouter.Handler("GET", "/api/v1/user/*path", helper.StoreOriginalPath(http.StripPrefix("/api/v1/users", stats.Middleware(userRouter))))
	mainRouter.appRouter.Handler("POST", "/api/v1/user/*path", helper.StoreOriginalPath(http.StripPrefix("/api/v1/users", stats.Middleware(userRouter))))
	mainRouter.appRouter.Handler("PUT", "/api/v1/user/*path", helper.StoreOriginalPath(http.StripPrefix("/api/v1/users", stats.Middleware(userRouter))))
	mainRouter.appRouter.Handler("DELETE", "/api/v1/user/*path", helper.StoreOriginalPath(http.StripPrefix("/api/v1/users", stats.Middleware(userRouter))))

	// farm router
	mainRouter.appRouter.Handler("GET", "/api/v1/farms/*path", helper.StoreOriginalPath(http.StripPrefix("/api/v1/farms", stats.Middleware(farmRouter))))
	mainRouter.appRouter.Handler("POST", "/api/v1/farms/*path", helper.StoreOriginalPath(http.StripPrefix("/api/v1/farms", farmRouter)))
	mainRouter.appRouter.Handler("PUT", "/api/v1/farms/*path", helper.StoreOriginalPath(http.StripPrefix("/api/v1/farms", farmRouter)))
	mainRouter.appRouter.Handler("DELETE", "/api/v1/farms/*path", helper.StoreOriginalPath(http.StripPrefix("/api/v1/farms", farmRouter)))

	// ponds router
	mainRouter.appRouter.Handler("GET", "/api/v1/ponds/*path", helper.StoreOriginalPath(http.StripPrefix("/api/v1/ponds", stats.Middleware(pondsRouter))))
	mainRouter.appRouter.Handler("POST", "/api/v1/ponds/*path", helper.StoreOriginalPath(http.StripPrefix("/api/v1/ponds", stats.Middleware(pondsRouter))))
	mainRouter.appRouter.Handler("PUT", "/api/v1/ponds/*path",  helper.StoreOriginalPath(http.StripPrefix("/api/v1/ponds", stats.Middleware(pondsRouter))))
	mainRouter.appRouter.Handler("DELETE", "/api/v1/ponds/*path",  helper.StoreOriginalPath(http.StripPrefix("/api/v1/ponds", stats.Middleware(pondsRouter))))

	// stats router
	mainRouter.appRouter.GET(
		"/api/v1/stats", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			helper.WriteToResponseBody(w, stats.GetEndpointStats())
		},
	)

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
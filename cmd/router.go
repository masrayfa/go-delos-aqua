package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/masrayfa/go-delos-aqua/internals/controller"
)

type Router struct {
	appRouter *httprouter.Router
}

func NewRouter() *Router {
	return &Router{
		appRouter: httprouter.New(),
	}
}

func NewUserRouter(userController controller.UserController) *httprouter.Router {
	router := NewRouter()
	router.appRouter.GET("/", userController.FindAll)
	router.appRouter.GET("/:id", userController.FindById)
	router.appRouter.POST("/", userController.Create)
	router.appRouter.PUT("/:id", userController.Update)
	router.appRouter.DELETE("/:id", userController.Delete)

	return router.appRouter
}

func NewFarmRouter(farmController controller.FarmController) *httprouter.Router {
	router := NewRouter()
	router.appRouter.GET("/", farmController.FindAll)
	router.appRouter.GET("/:id", farmController.FindById)
	router.appRouter.POST("/", farmController.Create)
	router.appRouter.PUT("/:id", farmController.Update)
	router.appRouter.DELETE("/:id", farmController.Delete)

	return router.appRouter
}

func NewPondsRouter(pondsController controller.PondsController) *httprouter.Router {
	router := NewRouter()
	router.appRouter.GET("/", pondsController.FindAll)
	router.appRouter.GET("/:id", pondsController.FindById)
	router.appRouter.POST("/", pondsController.Create)
	router.appRouter.PUT("/:id", pondsController.Update)
	router.appRouter.DELETE("/:id", pondsController.Delete)

	return router.appRouter
}

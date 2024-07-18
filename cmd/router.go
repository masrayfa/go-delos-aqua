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
	router.appRouter.GET("/user", userController.FindAll)
	router.appRouter.GET("/user/:id", userController.FindById)
	router.appRouter.POST("/user", userController.Create)
	router.appRouter.PUT("/user", userController.Update)
	router.appRouter.DELETE("/user/:id", userController.Delete)

	return router.appRouter
}

func NewFarmRouter(farmController controller.FarmController) *httprouter.Router {
	router := NewRouter()
	router.appRouter.GET("/farm", farmController.FindAll)
	router.appRouter.GET("/farm/:id", farmController.FindById)
	router.appRouter.POST("/farm", farmController.Create)
	router.appRouter.PUT("/farm", farmController.Update)
	router.appRouter.DELETE("/farm/:id", farmController.Delete)

	return router.appRouter
}

func NewPondsRouter(pondsController controller.PondsController) *httprouter.Router {
	router := NewRouter()
	router.appRouter.GET("/ponds", pondsController.FindAll)
	router.appRouter.GET("/ponds/:id", pondsController.FindById)
	router.appRouter.POST("/ponds", pondsController.Create)
	router.appRouter.PUT("/ponds", pondsController.Update)
	router.appRouter.DELETE("/ponds/:id", pondsController.Delete)

	return router.appRouter
}

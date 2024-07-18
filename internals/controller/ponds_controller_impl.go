package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/masrayfa/go-delos-aqua/internals/service"
)

type PondsControllerImpl struct {
	pondsService service.PondsService
}

func NewPondsController(pondsService service.PondsService) PondsController {
	return &PondsControllerImpl{
		pondsService: pondsService,
	}
}

func (p *PondsControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (p *PondsControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (p *PondsControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (p *PondsControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (p *PondsControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("implement me")
}
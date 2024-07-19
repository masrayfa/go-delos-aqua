package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/masrayfa/go-delos-aqua/internals/helper"
	"github.com/masrayfa/go-delos-aqua/internals/models/web"
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
	ponds, err := p.pondsService.FindAll(request.Context())
	if err != nil {
			var statusCode int

		switch {
			case errors.Is(err, helper.ErrNotFound):
				statusCode = http.StatusNotFound
			case errors.Is(err, helper.ErrBadRequest):
				statusCode = http.StatusBadRequest
			case errors.Is(err, helper.ErrUnathorized):
				statusCode = http.StatusUnauthorized
			case errors.Is(err, helper.ErrForbidden):
				statusCode = http.StatusForbidden
			case errors.Is(err, helper.ErrConflict):
				statusCode = http.StatusConflict
			default:
				statusCode = http.StatusInternalServerError
			
			webResponse := web.Response {
				Code: statusCode,
				Message: http.StatusText(statusCode),
			}

			helper.WriteToResponseBody(writer, webResponse)
		}

		webResponse := web.Response{
			Code:    statusCode,
			Message: http.StatusText(statusCode),
			Data: ponds,
		}

		helper.WriteToResponseBody(writer, webResponse)	
		return 
	}
}

func (p *PondsControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	UserId, err := strconv.Atoi(param)
	if err != nil {
		helper.WriteToResponseBody(writer, err.Error())
		return
	}

	user, err := p.pondsService.FindById(request.Context(), UserId)
	if err != nil {
		var statusCode int

		switch {
			case errors.Is(err, helper.ErrNotFound):
				statusCode = http.StatusNotFound
			case errors.Is(err, helper.ErrBadRequest):
				statusCode = http.StatusBadRequest
			case errors.Is(err, helper.ErrUnathorized):
				statusCode = http.StatusUnauthorized
			case errors.Is(err, helper.ErrForbidden):
				statusCode = http.StatusForbidden
			case errors.Is(err, helper.ErrConflict):
				statusCode = http.StatusConflict
			default:
				statusCode = http.StatusInternalServerError
		}

		webResponse := web.Response{
			Code:    statusCode,
			Message: http.StatusText(statusCode),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.Response{
		Code:   http.StatusOK,
		Message: "Pond found",
		Data:  user,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (p *PondsControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pondCreateReq := web.PondCreateRequest{}
	helper.ReadFromRequestBody(request, &pondCreateReq)

	_, err := p.pondsService.Create(request.Context(), pondCreateReq)
	if err != nil {
		var statusCode int

		switch {
			case errors.Is(err, helper.ErrNotFound):
				statusCode = http.StatusNotFound
			case errors.Is(err, helper.ErrBadRequest):
				statusCode = http.StatusBadRequest
			case errors.Is(err, helper.ErrUnathorized):
				statusCode = http.StatusUnauthorized
			case errors.Is(err, helper.ErrForbidden):
				statusCode = http.StatusForbidden
			case errors.Is(err, helper.ErrConflict):
				statusCode = http.StatusConflict
			default:
				statusCode = http.StatusInternalServerError
		}

		webResponse := web.Response{
			Code:    statusCode,
			Message: http.StatusText(statusCode),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.Response{
		Code:   http.StatusCreated,
		Message: "Pond created successfully",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (p *PondsControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	pondId, err := strconv.Atoi(param)
	if err != nil {
				var statusCode int

		switch {
			case errors.Is(err, helper.ErrNotFound):
				statusCode = http.StatusNotFound
			case errors.Is(err, helper.ErrBadRequest):
				statusCode = http.StatusBadRequest
			case errors.Is(err, helper.ErrUnathorized):
				statusCode = http.StatusUnauthorized
			case errors.Is(err, helper.ErrForbidden):
				statusCode = http.StatusForbidden
			case errors.Is(err, helper.ErrConflict):
				statusCode = http.StatusConflict
			default:
				statusCode = http.StatusInternalServerError
		}

		webResponse := web.Response{
			Code:    statusCode,
			Message: http.StatusText(statusCode),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	pondUpdateReq := web.PondUpdateRequest{}
	helper.ReadFromRequestBody(request, &pondUpdateReq)

	_, err = p.pondsService.Update(request.Context(), pondUpdateReq, pondId)
	if err != nil {
		var statusCode int

		switch {
			case errors.Is(err, helper.ErrNotFound):
				statusCode = http.StatusNotFound
			case errors.Is(err, helper.ErrBadRequest):
				statusCode = http.StatusBadRequest
			case errors.Is(err, helper.ErrUnathorized):
				statusCode = http.StatusUnauthorized
			case errors.Is(err, helper.ErrForbidden):
				statusCode = http.StatusForbidden
			case errors.Is(err, helper.ErrConflict):
				statusCode = http.StatusConflict
			default:
				statusCode = http.StatusInternalServerError
		}

		webResponse := web.Response{
			Code:    statusCode,
			Message: http.StatusText(statusCode),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.Response{
		Code:   http.StatusOK,
		Message: "User updated successfully",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (p *PondsControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	pondId, err := strconv.Atoi(param)
	if err != nil {
				var statusCode int

		switch {
			case errors.Is(err, helper.ErrNotFound):
				statusCode = http.StatusNotFound
			case errors.Is(err, helper.ErrBadRequest):
				statusCode = http.StatusBadRequest
			case errors.Is(err, helper.ErrUnathorized):
				statusCode = http.StatusUnauthorized
			case errors.Is(err, helper.ErrForbidden):
				statusCode = http.StatusForbidden
			case errors.Is(err, helper.ErrConflict):
				statusCode = http.StatusConflict
			default:
				statusCode = http.StatusInternalServerError
		}

		webResponse := web.Response{
			Code:    statusCode,
			Message: http.StatusText(statusCode),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	err = p.pondsService.Delete(request.Context(), pondId)
	if err != nil {
		var statusCode int

		switch {
			case errors.Is(err, helper.ErrNotFound):
				statusCode = http.StatusNotFound
			case errors.Is(err, helper.ErrBadRequest):
				statusCode = http.StatusBadRequest
			case errors.Is(err, helper.ErrUnathorized):
				statusCode = http.StatusUnauthorized
			case errors.Is(err, helper.ErrForbidden):
				statusCode = http.StatusForbidden
			case errors.Is(err, helper.ErrConflict):
				statusCode = http.StatusConflict
			default:
				statusCode = http.StatusInternalServerError
		}

		webResponse := web.Response{
			Code:    statusCode,
			Message: http.StatusText(statusCode),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.Response{
		Code:   http.StatusOK,
		Message: "Pond deleted successfully",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
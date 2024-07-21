package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/masrayfa/go-delos-aqua/internals/helper"
	"github.com/masrayfa/go-delos-aqua/internals/models/web"
	"github.com/masrayfa/go-delos-aqua/internals/service"
)

type FarmControllerImpl struct {
	farmService service.FarmService
}

func NewFarmController(farmService service.FarmService) FarmController {
	return &FarmControllerImpl{
		farmService: farmService,
	}
}

func (controller *FarmControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	log.Println("@FarmControllerImpl.FindAll")

	farm, err := controller.farmService.FindAll(request.Context())
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
		Code:    http.StatusOK,
		Message: "Farm found",
		Data:    farm,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FarmControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	FarmId, err := strconv.Atoi(param)
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

	farm, err := controller.farmService.FindById(request.Context(), FarmId)
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
		Code:    http.StatusOK,
		Message: "Farm found",
		Data:    farm,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FarmControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var farmRequest web.FarmRequest
	err := helper.ReadFromRequestBody(request, &farmRequest)
	if err != nil {
		webResponse := web.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	_, err = controller.farmService.Create(request.Context(), farmRequest)
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
		Code:    http.StatusOK,
		Message: "Farm created",
		Data:    farmRequest,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FarmControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var farmRequest web.FarmRequest
	err := helper.ReadFromRequestBody(request, &farmRequest)
	if err != nil {
		webResponse := web.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	_, err = controller.farmService.Update(request.Context(), farmRequest)
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
		Code:    http.StatusOK,
		Message: "Farm updated",
		Data:    farmRequest,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FarmControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	FarmId, err := strconv.Atoi(param)
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

	err = controller.farmService.Delete(request.Context(), FarmId)
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
		Code:    http.StatusOK,
		Message: "Farm deleted",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
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

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (uc *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	user, err := uc.userService.FindAll(request.Context())
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
		Message: "User found",
		Data:  user,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (uc *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	UserId, err := strconv.Atoi(param)
	if err != nil {
		helper.WriteToResponseBody(writer, err.Error())
		return
	}

	user, err := uc.userService.FindById(request.Context(), UserId)
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
		Message: "User found",
		Data:  user,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (uc *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateReq := web.UserCreate{}
	helper.ReadFromRequestBody(request, &userCreateReq)

	_, err := uc.userService.Create(request.Context(), userCreateReq)
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
		Message: "User created successfully",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (uc *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateReq := web.UserUpdate{}
	helper.ReadFromRequestBody(request, &userUpdateReq)

	_, err := uc.userService.Update(request.Context(), userUpdateReq)
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
			Message: err.Error(),
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

func (uc *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	UserId, err := strconv.Atoi(param)
	if err != nil {
		helper.WriteToResponseBody(writer, err.Error())
		return
	}

	err = uc.userService.Delete(request.Context(), UserId)
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

	webReponse := web.Response{
		Code:   http.StatusOK,
		Message: "User deleted successfully",
	}

	helper.WriteToResponseBody(writer, webReponse)
}
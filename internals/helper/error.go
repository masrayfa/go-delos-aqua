package helper

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal server error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested Item is not found")

	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your Item already exist")

	// ErrUnathorized will throw if the user is not authorized
	ErrUnathorized = errors.New("you are not authorized")

	// ErrForbidden will throw if the user is not authorized to access the data
	ErrForbidden = errors.New("you are not authorized to access the data")

	// ErrBadRequest will throw if the request is invalid
	ErrBadRequest = errors.New("bad request")

)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
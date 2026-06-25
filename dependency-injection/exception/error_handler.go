package exception

import (
	"net/http"
	"restful-api/helper"
	"restful-api/model/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(
	writer http.ResponseWriter,
	request *http.Request,
	err interface{},
) {

	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationErrors(
	writer http.ResponseWriter,
	request *http.Request,
	err interface{},
) bool {
	exception, ok := err.(validator.ValidationErrors)
	if !ok {
		return false
	}

	writer.Header().Set("Context-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "BAD REQUEST",
		Data:   exception.Error(),
	}

	helper.WriteToResponseBody(writer, webResponse)
	return true
}

func notFoundError(
	writer http.ResponseWriter,
	request *http.Request,
	err interface{},
) bool {
	exception, ok := err.(NotFoundError)
	if !ok {
		return false
	}
	writer.Header().Set("Context-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusNotFound,
		Status: "Not Found",
		Data:   exception.Error,
	}

	helper.WriteToResponseBody(writer, webResponse)
	return true
}

func internalServerError(
	writer http.ResponseWriter,
	request *http.Request,
	err interface{},
) {

	writer.Header().Set("Context-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

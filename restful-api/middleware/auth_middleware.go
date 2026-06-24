package middleware

import (
	"net/http"
	"restful-api/helper"
	"restful-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (m *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		m.Handler.ServeHTTP(writer, request)
		return
	}

	writer.Header().Set("Context-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)

	webResponse := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "Unauthorized",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

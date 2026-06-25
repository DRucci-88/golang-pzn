package main

import (
	"net/http"
	"restful-api/helper"
	"restful-api/middleware"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: authMiddleware,
	}
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func main() {
	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

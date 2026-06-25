//go:build wireinject
// +build wireinject

package main

import (
	"net/http"
	"restful-api/app"
	"restful-api/controller"
	"restful-api/middleware"
	"restful-api/repository"
	"restful-api/service"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		app.NewRouter,
		NewValidator,
		repository.NewCategoryRepository,
		service.NewCategoryService,
		controller.NewCategoryController,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}

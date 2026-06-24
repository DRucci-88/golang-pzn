package app

import (
	"encoding/json"
	"net/http"
	"restful-api/controller"
	"restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/health", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		encoder := json.NewEncoder(w)
		encoder.Encode("Ok Good")
	})

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}

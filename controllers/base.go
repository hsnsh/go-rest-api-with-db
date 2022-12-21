package controllers

import "github.com/gorilla/mux"

type IBaseController interface {
	InitializeRoutes(Router *mux.Router)
}

package controllers

import (
	"github.com/gorilla/mux"
	"net/http"

	. "go-rest-api-with-db/helpers"
)

type bookController struct {
	_logger IFileLogger
}

func NewBookController(logger IFileLogger) IBaseController {
	return &bookController{
		_logger: logger,
	}
}

func (c *bookController) InitializeRoutes(Router *mux.Router) {
	Router.HandleFunc("/api/books", c.getAllBooks).Methods("GET")
}

func (c *bookController) getAllBooks(w http.ResponseWriter, r *http.Request) {
	defer HandlePanicAndRecovery(w)

	RespondWithJSON(w, http.StatusOK, "All Books listed")
}

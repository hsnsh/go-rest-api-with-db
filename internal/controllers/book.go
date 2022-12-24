package controllers

import (
	"github.com/HsnCorp/go-hsn-library/logger"
	"github.com/gorilla/mux"
	. "go-rest-api-with-db/internal/helpers"
	"net/http"
)

type bookController struct {
	logger logger.IFileLogger
	router *mux.Router
}

func RegisterBookController(appLogger logger.IFileLogger, appRouter *mux.Router) {
	bookController{logger: appLogger, router: appRouter}.initializeRoutes()
}

func (c bookController) initializeRoutes() {
	c.router.HandleFunc("/api/books", c.getAllBooks).Methods("GET")
}

func (c bookController) getAllBooks(w http.ResponseWriter, r *http.Request) {
	defer HandlePanicAndRecovery(w)

	RespondWithJSON(w, http.StatusOK, "All Books listed")
}

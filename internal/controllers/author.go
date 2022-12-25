package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/HsnCorp/go-hsn-library/logger"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	. "go-rest-api-with-db/internal/dtos"
	. "go-rest-api-with-db/internal/helpers"
	. "go-rest-api-with-db/internal/services"
	"net/http"
)

const baseUrl = "/api/authors"

type authorController struct {
	logger        logger.IFileLogger
	router        *mux.Router
	authorService IAuthorAppService
}

func RegisterAuthorController(appLogger logger.IFileLogger, appRouter *mux.Router, authorAppService IAuthorAppService) {
	authorController{logger: appLogger, router: appRouter, authorService: authorAppService}.initializeRoutes()
}

func (c authorController) initializeRoutes() {
	c.router.HandleFunc(baseUrl, c.getAllAuthors).Methods("GET")
	c.router.HandleFunc(baseUrl+"/{id}", c.getAuthorById).Methods("GET")
	c.router.HandleFunc(baseUrl, c.createAuthor).Methods("POST")
	c.router.HandleFunc(baseUrl+"/{id}", c.updateAuthor).Methods("PUT")
	c.router.HandleFunc(baseUrl+"/{id}", c.deleteAuthor).Methods("DELETE")
}

func (c authorController) getAllAuthors(w http.ResponseWriter, r *http.Request) {
	defer HandlePanicAndRecovery(w, c.logger)

	authors, errResult := c.authorService.GetAuthorList()
	if errResult != nil {
		c.logger.Warning(errResult.Error())
		RespondErrorWithMessage(w, errResult.Error())
		return
	}

	if authors == nil {
		authors = make([]AuthorDto, 0)
	}
	RespondOkWithData(w, authors)
}

func (c authorController) getAuthorById(w http.ResponseWriter, r *http.Request) {
	defer HandlePanicAndRecovery(w, c.logger)

	// Get variables from request url
	variables := mux.Vars(r)

	// Get ID parameter
	key := variables["id"]

	// Check ID Parameter is valid
	searchId, errParse := uuid.FromString(key)
	if errParse != nil {
		c.logger.Warning("Invalid ID; " + errParse.Error())
		RespondErrorWithMessage(w, "Invalid ID")
		return
	}

	// Get from on the Store
	author, errResult := c.authorService.GetAuthorById(searchId)
	if errResult != nil {
		c.logger.Warning(errResult.Error())
		RespondErrorWithMessage(w, errResult.Error())
		return
	}

	RespondOkWithData(w, author)
}

func (c authorController) createAuthor(w http.ResponseWriter, r *http.Request) {
	defer HandlePanicAndRecovery(w, c.logger)

	// Get create dto from request body
	var authorCreateDto AuthorCreateDto
	errDecode := json.NewDecoder(r.Body).Decode(&authorCreateDto)
	if errDecode != nil {
		c.logger.Warning(errDecode.Error())
		RespondErrorWithMessage(w, errDecode.Error())
		return
	}

	// Create on the store
	author, errCreate := c.authorService.CreateAuthor(authorCreateDto)
	if errCreate != nil {
		c.logger.Warning(errCreate.Error())
		RespondErrorWithMessage(w, errCreate.Error())
		return
	}

	RespondCreatedWithData(w, author)
}

func (c authorController) updateAuthor(w http.ResponseWriter, r *http.Request) {
	defer HandlePanicAndRecovery(w, c.logger)

	// Get variables from request url
	variables := mux.Vars(r)

	// Get ID parameter
	key := variables["id"]

	// Check ID Parameter is valid
	updateId, errParse := uuid.FromString(key)
	if errParse != nil {
		c.logger.Warning("Invalid ID, " + errParse.Error())
		RespondErrorWithMessage(w, "Invalid ID")
		return
	}

	// Get update dto from request body
	var authorUpdateDto AuthorUpdateDto
	errDecode := json.NewDecoder(r.Body).Decode(&authorUpdateDto)
	if errDecode != nil {
		c.logger.Warning(errDecode.Error())
		RespondErrorWithMessage(w, errDecode.Error())
		return
	}

	// Update on the Store
	author, errUpdate := c.authorService.UpdateAuthor(updateId, authorUpdateDto)
	if errUpdate != nil {
		c.logger.Warning(errUpdate.Error())
		RespondErrorWithMessage(w, errUpdate.Error())
		return
	}

	RespondOkWithData(w, author)
}

func (c authorController) deleteAuthor(w http.ResponseWriter, r *http.Request) {
	defer HandlePanicAndRecovery(w, c.logger)

	// Get variables from request url
	variables := mux.Vars(r)

	// Get ID parameter
	key := variables["id"]

	// Check ID Parameter is valid
	searchId, errParse := uuid.FromString(key)
	if errParse != nil {
		c.logger.Warning("Invalid ID, " + errParse.Error())
		RespondErrorWithMessage(w, "Invalid ID")
		return
	}

	// Delete on the Store
	errDelete := c.authorService.DeleteAuthor(searchId)
	if errDelete != nil {
		c.logger.Warning(errDelete.Error())
		RespondErrorWithMessage(w, errDelete.Error())
		return
	}

	RespondOkWithData(w, fmt.Sprintf("%s record successfully deleted", key))
}

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
	defer HandlePanicAndRecovery(w)

	authors, errResult := c.authorService.GetAuthorList()
	if errResult != nil {
		fmt.Println(errResult.Error())
		RespondWithError(w, http.StatusInternalServerError, errResult.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, authors)
}

func (c authorController) getAuthorById(w http.ResponseWriter, r *http.Request) {
	defer HandlePanicAndRecovery(w)

	// Get variables from request url
	variables := mux.Vars(r)

	// Get ID parameter
	key := variables["id"]

	// Check ID Parameter is valid
	searchId, err := uuid.FromString(key)
	if err != nil {
		fmt.Println(err.Error())
		RespondWithError(w, http.StatusBadRequest, "Invalid Author ID")
		return
	}

	// Get from on the Store
	author, errResult := c.authorService.GetAuthorById(searchId)
	if errResult != nil {
		fmt.Println(errResult.Error())
		RespondWithError(w, http.StatusNotFound, errResult.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, author)
}

func (c authorController) createAuthor(w http.ResponseWriter, r *http.Request) {
	defer HandlePanicAndRecovery(w)

	// Get create dto from request body
	var authorCreateDto AuthorCreateDto
	errDecode := json.NewDecoder(r.Body).Decode(&authorCreateDto)
	if errDecode != nil {
		fmt.Println(errDecode.Error())
		RespondWithError(w, http.StatusBadRequest, errDecode.Error())
		return
	}

	// Create on the store
	author, errCreate := c.authorService.CreateAuthor(authorCreateDto)
	if errCreate != nil {
		fmt.Println(errCreate.Error())
		RespondWithError(w, http.StatusBadRequest, errCreate.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, author)
}

func (c authorController) updateAuthor(w http.ResponseWriter, r *http.Request) {
	defer HandlePanicAndRecovery(w)

	// Get variables from request url
	variables := mux.Vars(r)

	// Get ID parameter
	key := variables["id"]

	// Check ID Parameter is valid
	updateId, errParse := uuid.FromString(key)
	if errParse != nil {
		fmt.Println(errParse.Error())
		RespondWithError(w, http.StatusBadRequest, errParse.Error())
		return
	}

	// Get update dto from request body
	var authorUpdateDto AuthorUpdateDto
	errDecode := json.NewDecoder(r.Body).Decode(&authorUpdateDto)
	if errDecode != nil {
		fmt.Println(errDecode.Error())
		RespondWithError(w, http.StatusBadRequest, errDecode.Error())
		return
	}

	// Update on the Store
	author, errUpdate := c.authorService.UpdateAuthor(updateId, authorUpdateDto)
	if errUpdate != nil {
		fmt.Println(errUpdate.Error())
		RespondWithError(w, http.StatusBadRequest, errUpdate.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, author)
}

func (c authorController) deleteAuthor(w http.ResponseWriter, r *http.Request) {
	defer HandlePanicAndRecovery(w)

	// Get variables from request url
	variables := mux.Vars(r)

	// Get ID parameter
	key := variables["id"]

	// Check ID Parameter is valid
	searchId, errParse := uuid.FromString(key)
	if errParse != nil {
		fmt.Println(errParse.Error())
		RespondWithError(w, http.StatusBadRequest, errParse.Error())
		return
	}

	// Delete on the Store
	errDelete := c.authorService.DeleteAuthor(searchId)
	if errDelete != nil {
		fmt.Println(errDelete.Error())
		RespondWithError(w, http.StatusBadRequest, errDelete.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, "Delete operation is successfully completed")
}

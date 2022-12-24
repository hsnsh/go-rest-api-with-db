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

type productController struct {
	logger         logger.IFileLogger
	router         *mux.Router
	productService IProductAppService
}

func RegisterProductController(appLogger logger.IFileLogger, appRouter *mux.Router, productAppService IProductAppService) {
	productController{logger: appLogger, router: appRouter, productService: productAppService}.initializeRoutes()
}

func (c productController) initializeRoutes() {
	c.router.HandleFunc("/api/products", c.getAllProducts).Methods("GET")
	c.router.HandleFunc("/api/products/{id}", c.getProductById).Methods("GET")
	c.router.HandleFunc("/api/products", c.createProduct).Methods("POST")
	c.router.HandleFunc("/api/products/{id}", c.updateProduct).Methods("PUT")
	c.router.HandleFunc("/api/products/{id}", c.deleteProduct).Methods("DELETE")
}

func (c productController) getAllProducts(w http.ResponseWriter, r *http.Request) {
	defer HandlePanicAndRecovery(w)

	products, errResult := c.productService.GetProductList()
	if errResult != nil {
		fmt.Println(errResult.Error())
		RespondWithError(w, http.StatusInternalServerError, errResult.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, products)
}

func (c productController) getProductById(w http.ResponseWriter, r *http.Request) {
	defer HandlePanicAndRecovery(w)

	// Get variables from request url
	variables := mux.Vars(r)

	// Get ID parameter
	key := variables["id"]

	// Check ID Parameter is valid
	searchId, err := uuid.FromString(key)
	if err != nil {
		fmt.Println(err.Error())
		RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	// Get from on the Store
	product, errResult := c.productService.GetProductById(searchId)
	if errResult != nil {
		fmt.Println(errResult.Error())
		RespondWithError(w, http.StatusNotFound, errResult.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, product)
}

func (c productController) createProduct(w http.ResponseWriter, r *http.Request) {
	defer HandlePanicAndRecovery(w)

	// Get create dto from request body
	var productCreateDto ProductCreateDto
	errDecode := json.NewDecoder(r.Body).Decode(&productCreateDto)
	if errDecode != nil {
		fmt.Println(errDecode.Error())
		RespondWithError(w, http.StatusBadRequest, errDecode.Error())
		return
	}

	// Create on the store
	product, errCreate := c.productService.CreateProduct(productCreateDto)
	if errCreate != nil {
		fmt.Println(errCreate.Error())
		RespondWithError(w, http.StatusBadRequest, errCreate.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, product)
}

func (c productController) updateProduct(w http.ResponseWriter, r *http.Request) {
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
	var productUpdateDto ProductUpdateDto
	errDecode := json.NewDecoder(r.Body).Decode(&productUpdateDto)
	if errDecode != nil {
		fmt.Println(errDecode.Error())
		RespondWithError(w, http.StatusBadRequest, errDecode.Error())
		return
	}

	// Update on the Store
	product, errUpdate := c.productService.UpdateProduct(updateId, productUpdateDto)
	if errUpdate != nil {
		fmt.Println(errUpdate.Error())
		RespondWithError(w, http.StatusBadRequest, errUpdate.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, product)
}

func (c productController) deleteProduct(w http.ResponseWriter, r *http.Request) {
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
	errDelete := c.productService.DeleteProduct(searchId)
	if errDelete != nil {
		fmt.Println(errDelete.Error())
		RespondWithError(w, http.StatusBadRequest, errDelete.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, "Delete operation is successfully completed")
}

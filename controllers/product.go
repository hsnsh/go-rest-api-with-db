package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/HsnCorp/go-hsn-library/logger"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"

	. "go-rest-api-with-db/dtos"
	. "go-rest-api-with-db/helpers"
	. "go-rest-api-with-db/services"
)

type productController struct {
	_logger            logger.IFileLogger
	_productAppService IProductAppService
}

func NewProductController(logger logger.IFileLogger, productAppService IProductAppService) IBaseController {
	return &productController{
		_logger:            logger,
		_productAppService: productAppService,
	}
}

func (c productController) InitializeRoutes(Router *mux.Router) {
	Router.HandleFunc("/api/products", c.getAllProducts).Methods("GET")
	Router.HandleFunc("/api/products/{id}", c.getProductById).Methods("GET")
	Router.HandleFunc("/api/products", c.createProduct).Methods("POST")
	Router.HandleFunc("/api/products/{id}", c.updateProduct).Methods("PUT")
	Router.HandleFunc("/api/products/{id}", c.deleteProduct).Methods("DELETE")
}

func (c productController) getAllProducts(w http.ResponseWriter, r *http.Request) {
	defer HandlePanicAndRecovery(w)

	products, errResult := c._productAppService.GetProductList()
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
	product, errResult := c._productAppService.GetProductById(searchId)
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
	product, errCreate := c._productAppService.CreateProduct(productCreateDto)
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
	product, errUpdate := c._productAppService.UpdateProduct(updateId, productUpdateDto)
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
	errDelete := c._productAppService.DeleteProduct(searchId)
	if errDelete != nil {
		fmt.Println(errDelete.Error())
		RespondWithError(w, http.StatusBadRequest, errDelete.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, "Delete operation is successfully completed")
}

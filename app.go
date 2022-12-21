package main

import (
	"github.com/google/uuid"
	c "go-rest-api-with-db/controllers"
	d "go-rest-api-with-db/domain"
	. "go-rest-api-with-db/helpers"
	r "go-rest-api-with-db/repositories"
	s "go-rest-api-with-db/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"

	"net/http"
	// tom: go get required
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB

	bookController    c.IBaseController
	productController c.IBaseController
}

// Public functions

func (a *App) Initialize(connectionString string) {

	appLogger.Info("Go RestAPI initializing...")

	a.initializeDatabase(connectionString)
	a.seedData()

	// Register Repositories
	dao := r.NewDAO(a.DB)

	// Register Services
	productAppService := s.NewProductAppService(appLogger, dao)

	// Initialize Controllers
	a.bookController = c.NewBookController(appLogger)
	a.productController = c.NewProductController(appLogger, productAppService)

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	appLogger.Info("Go RestAPI listening ... [ " + addr + " ]")
	if err := http.ListenAndServe(addr, a.Router); err != nil {
		appLogger.Fatal(err.Error())
	} else {
		appLogger.Info("Go RestAPI terminated.")
	}
}

// Private Functions

func (a *App) initializeDatabase(dsn string) {

	dsn = "data/test.db"
	db, conErr := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	//db, conErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if conErr != nil {
		appLogger.Fatal(conErr.Error())
	} else {
		appLogger.Info("Successfully connected to the database")
	}

	// Migrate the schema
	var models []interface{}
	models = append(models, &d.Book{})
	models = append(models, &d.Product{})

	if err := db.AutoMigrate(models...); err != nil {
		appLogger.Fatal(err.Error())
	} else {
		appLogger.Info("AutoMigrate completed")
	}

	a.DB = db
}

func (a *App) seedData() {

	id := uuid.New().String()

	// Create
	a.DB.Create(&d.Product{BaseEntity: d.BaseEntity{ID: id}, Name: "D42", Price: 100})

	// Read
	var product d.Product
	a.DB.First(&product, "id = ?", id)      // find product with integer primary key
	a.DB.First(&product, "name = ?", "D42") // find product with name D42

	// Update - update product's price to 200
	a.DB.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	a.DB.Model(&product).Updates(d.Product{Price: 200, Name: "F42"}) // non-zero fields
	a.DB.Model(&product).Updates(map[string]interface{}{"Price": 200, "Name": "F42"})

	// Delete - delete product
	a.DB.Delete(&product, "id = ?", id)
}

func (a *App) initializeRoutes() {

	a.Router = mux.NewRouter().StrictSlash(true)

	if os.Getenv("APP_ENV") != "prod" {
		a.Router.HandleFunc("/", a.handleIndex).Methods("GET")
	}

	//Register routes of controllers
	a.bookController.InitializeRoutes(a.Router)
	a.productController.InitializeRoutes(a.Router)
}

func (a *App) handleIndex(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, http.StatusOK, "Welcome GoLang RestAPI")
}

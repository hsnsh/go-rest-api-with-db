package main

import (
	"errors"
	"fmt"
	c "go-rest-api-with-db/controllers"
	d "go-rest-api-with-db/domain"
	. "go-rest-api-with-db/helpers"
	r "go-rest-api-with-db/repositories"
	s "go-rest-api-with-db/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

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

	gormConfig := gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // Slow SQL threshold
				LogLevel:                  logger.Silent, // Log level
				IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,         // Disable color
			},
		),
	}

	dsn = "data/test.db" // "file::memory:?cache=shared"
	db, conErr := gorm.Open(sqlite.Open(dsn), &gormConfig)

	//db, conErr := gorm.Open(postgres.Open(dsn), &gormConfig)
	if conErr != nil {
		appLogger.Fatal(conErr.Error())
	} else {
		appLogger.Info("Successfully connected to the database")
	}

	// Migrate the schema
	var models []interface{}
	models = append(models, &d.Book{})
	models = append(models, &d.Product{})
	models = append(models, &d.ProductLanguage{})
	models = append(models, &d.CategoryType{})

	if err := db.AutoMigrate(models...); err != nil {
		appLogger.Fatal(err.Error())
	} else {
		appLogger.Info("AutoMigrate completed")
	}

	a.DB = db
}

func (a *App) seedData() {

	createProduct := d.Product{Name: "D42", Price: 99}

	// Create
	r0 := a.DB.Create(&createProduct)
	if r0.Error != nil {
		fmt.Println("CREATE ERROR : " + r0.Error.Error())
	}
	id := createProduct.ID // uuid.NewV4()

	fmt.Println("CREATED PRODUCT ID: " + id.String())

	// Read
	var product d.Product
	result := a.DB.First(&product, "id = ?", id.String()) // find product with integer primary key
	// Check if returns RecordNotFound error
	errors.Is(result.Error, gorm.ErrRecordNotFound)

	a.DB.First(&product, "name = ?", "D42") // find product with name D42

	// Update - update product's price to 200
	r1 := a.DB.Updates(&product)
	if r1.Error != nil {
		fmt.Println("UPDATE ERROR1 : " + r1.Error.Error())
	}

	r2 := a.DB.Model(&product).Update("Price", 102)
	if r2.Error != nil {
		fmt.Println("UPDATE ERROR2 : " + r2.Error.Error())
	}

	// Update - update multiple fields
	r3 := a.DB.Model(&product).Updates(d.Product{Price: 222, Name: "F42"}) // non-zero fields
	if r3.Error != nil {
		fmt.Println("UPDATE ERROR3 : " + r3.Error.Error())
	}

	r4 := a.DB.Model(&product).Updates(map[string]interface{}{"Price": 444, "Name": "F42"})
	if r4.Error != nil {
		fmt.Println("UPDATE ERROR4 : " + r4.Error.Error())
	}

	//Delete - delete product
	r5 := a.DB.Delete(&d.Product{}, "id = ?", id.String())
	if r5.Error != nil {
		fmt.Println("DELETE ERROR : " + r5.Error.Error())
	}
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

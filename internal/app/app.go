package app

import (
	"errors"
	"fmt"
	. "github.com/HsnCorp/go-hsn-library/logger"
	. "go-rest-api-with-db/internal/controllers"
	"go-rest-api-with-db/internal/domain"
	. "go-rest-api-with-db/internal/helpers"
	r "go-rest-api-with-db/internal/repositories"
	s "go-rest-api-with-db/internal/services"
	"gorm.io/driver/postgres"
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

type app struct {
	appRouter *mux.Router
	appDB     *gorm.DB
	appLogger IFileLogger
}

func New(fileLogger IFileLogger) *app {
	return &app{appLogger: fileLogger}
}

// Public functions

func (a *app) Initialize(connectionString string) {

	a.appLogger.Info("Go RestAPI initializing...")

	a.initializeDatabase(connectionString)
	a.seedData()

	// Register Repositories
	dao := r.NewDAO(a.appDB)

	// Register Services
	productAppService := s.NewProductAppService(a.appLogger, dao)

	// Initialize Routes
	a.appRouter = mux.NewRouter().StrictSlash(true)
	if os.Getenv("APP_ENV") != "prod" {
		a.appRouter.HandleFunc("/", a.handleIndex).Methods("GET")
	}
	RegisterBookController(a.appLogger, a.appRouter)
	RegisterProductController(a.appLogger, a.appRouter, productAppService)
}

func (a *app) Run(addr string) {
	a.appLogger.Info("Go RestAPI listening ... [ " + addr + " ]")
	if err := http.ListenAndServe(addr, a.appRouter); err != nil {
		a.appLogger.Fatal(err.Error())
	} else {
		a.appLogger.Info("Go RestAPI terminated.")
	}
}

// Private Functions

func (a *app) initializeDatabase(dsn string) {

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

	//dsn = "data/test.db" // "file::memory:?cache=shared"
	//db, conErr := gorm.Open(sqlite.Open(dsn), &gormConfig)

	db, conErr := gorm.Open(postgres.Open(dsn), &gormConfig)
	if conErr != nil {
		a.appLogger.Fatal(conErr.Error())
	} else {
		a.appLogger.Info("Successfully connected to the database")
	}

	// Migrate the schema
	var models []interface{}
	models = append(models, &domain.Product{})
	models = append(models, &domain.ProductLanguage{})
	models = append(models, &domain.CategoryType{})

	models = append(models, &domain.Book{})
	models = append(models, &domain.BookContent{})
	models = append(models, &domain.BookFile{})
	models = append(models, &domain.Author{})
	models = append(models, &domain.Publisher{})
	models = append(models, &domain.PublisherBook{})

	if err := db.AutoMigrate(models...); err != nil {
		a.appLogger.Fatal(err.Error())
	} else {
		a.appLogger.Info("AutoMigrate completed")
	}

	a.appDB = db
}

func (a *app) seedData() {

	createProduct := domain.Product{Name: "D42", Price: 99}

	// Create
	r0 := a.appDB.Create(&createProduct)
	if r0.Error != nil {
		fmt.Println("CREATE ERROR : " + r0.Error.Error())
	}
	id := createProduct.ID // uuid.NewV4()

	fmt.Println("CREATED PRODUCT ID: " + id.String())

	// Read
	var product domain.Product
	result := a.appDB.First(&product, "id = ?", id.String()) // find product with integer primary key
	// Check if returns RecordNotFound error
	errors.Is(result.Error, gorm.ErrRecordNotFound)

	a.appDB.First(&product, "name = ?", "D42") // find product with name D42

	// Update - update product's price to 200
	r1 := a.appDB.Updates(&product)
	if r1.Error != nil {
		fmt.Println("UPDATE ERROR1 : " + r1.Error.Error())
	}

	r2 := a.appDB.Model(&product).Update("Price", 102)
	if r2.Error != nil {
		fmt.Println("UPDATE ERROR2 : " + r2.Error.Error())
	}

	// Update - update multiple fields
	r3 := a.appDB.Model(&product).Updates(domain.Product{Price: 222, Name: "F42"}) // non-zero fields
	if r3.Error != nil {
		fmt.Println("UPDATE ERROR3 : " + r3.Error.Error())
	}

	r4 := a.appDB.Model(&product).Updates(map[string]interface{}{"Price": 444, "Name": "F42"})
	if r4.Error != nil {
		fmt.Println("UPDATE ERROR4 : " + r4.Error.Error())
	}

	//Delete - delete product
	r5 := a.appDB.Delete(&domain.Product{}, "id = ?", id.String())
	if r5.Error != nil {
		fmt.Println("DELETE ERROR : " + r5.Error.Error())
	}
}

func (a *app) handleIndex(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, http.StatusOK, "Welcome GoLang RestAPI")
}

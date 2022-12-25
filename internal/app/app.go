package app

import (
	. "github.com/HsnCorp/go-hsn-library/logger"
	. "go-rest-api-with-db/internal/controllers"
	"go-rest-api-with-db/internal/domain"
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

	// Initialize Database
	a.initializeDatabase(connectionString)

	// Register Repositories
	dao := r.NewDataAccessLayer(a.appDB)

	// Register Services
	authorAppService := s.NewAuthorAppService(a.appLogger, dao)

	// Initialize Routes
	a.appRouter = mux.NewRouter().StrictSlash(true)
	if os.Getenv("APP_ENV") != "prod" {
		a.appRouter.HandleFunc("/", a.handleIndex).Methods("GET")
	}
	RegisterAuthorController(a.appLogger, a.appRouter, authorAppService)
}

func (a *app) GetDB() *gorm.DB {
	return a.appDB
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

func (a *app) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome GoLang RestAPI"))
}

package app

import (
	"fmt"
	. "github.com/HsnCorp/go-hsn-library/logger"
	"github.com/gorilla/mux"
	//_ "github.com/lib/pq"
	"go-rest-api-with-db/internal/config"
	. "go-rest-api-with-db/internal/controllers"
	"go-rest-api-with-db/internal/domain"
	r "go-rest-api-with-db/internal/repositories"
	s "go-rest-api-with-db/internal/services"
	"go-rest-api-with-db/internal/storage"
	"gorm.io/gorm"
	"net/http"
)

type app struct {
	appRouter   *mux.Router
	appDB       *gorm.DB
	appSettings *config.AppSettings
	appLogger   IFileLogger
}

func New(fileLogger IFileLogger, appSettings *config.AppSettings) *app {
	instance := app{appLogger: fileLogger, appSettings: appSettings}
	instance.initialize()
	return &instance
}

// Public functions

func (a *app) GetDB() *gorm.DB {
	return a.appDB
}

func (a *app) Run() {
	addr := fmt.Sprintf("%s:%s", a.appSettings.AppHost, a.appSettings.AppPort)

	a.appLogger.Info("Go RestAPI listening ... [ " + addr + " ]")
	if err := http.ListenAndServe(addr, a.appRouter); err != nil {
		a.appLogger.Fatal(err.Error())
	} else {
		a.appLogger.Info("Go RestAPI terminated.")
	}
}

// Private Functions
func (a *app) initialize() {

	a.appLogger.Info(a.appSettings.AppTitle + " initializing... [ v" + a.appSettings.AppVersion + " ]")

	// Initialize Database
	a.initializeDatabase()

	// Register Repositories
	dao := r.NewDataAccessLayer(a.appDB)

	// Register Services
	//s.NewAuthorService(s.WithAuthorRepository(dao.AuthorRepository()))
	authorAppService := s.NewAuthorAppService(a.appLogger, dao)

	// Initialize Routes
	a.appRouter = mux.NewRouter().StrictSlash(true)
	if a.appSettings.AppEnvironment != "prod" {
		a.appRouter.HandleFunc("/", a.handleIndex).Methods("GET")
	}
	RegisterAuthorController(a.appLogger, a.appRouter, authorAppService)
}

func (a *app) initializeDatabase() {

	config := storage.PostgresConfig{
		Host:     a.appSettings.DbHost,
		Port:     a.appSettings.DbPort,
		Password: a.appSettings.DbPass,
		User:     a.appSettings.DbUser,
		DbName:   a.appSettings.DbName,
		SslMode:  a.appSettings.DbSslMode,
	}

	db, conErr := storage.PostgresNewConnectionWithConfig(&config)
	if conErr != nil {
		a.appLogger.Fatal(conErr.Error())
	} else {
		a.appLogger.Info(fmt.Sprintf("Successfully connected to the database. [ %s on %s:%s ]",
			a.appSettings.DbName,
			a.appSettings.DbHost,
			a.appSettings.DbPort,
		))
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
		a.appLogger.Info("AutoMigrate completed.")
	}

	a.appDB = db
}

func (a *app) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome GoLang RestAPI"))
}

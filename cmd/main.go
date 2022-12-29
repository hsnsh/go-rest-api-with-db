package main

import (
	"fmt"
	"github.com/HsnCorp/go-hsn-library/logger"
	"github.com/joho/godotenv"
	"go-rest-api-with-db/internal/app"
	"go-rest-api-with-db/internal/config"
	"go-rest-api-with-db/internal/domain"
	"gorm.io/gorm"
	"os"
	"strings"
)

var (
	appLogger   logger.IFileLogger
	appSettings config.AppSettings
)

func init() {

	appLogger = logger.NewFileLogger()

	var err error
	err = godotenv.Load() // The Original .env
	if err != nil {
		appLogger.Fatal("Error loading .env file")
	}

	env := os.Getenv("APP_ENV")
	if len(env) < 1 {
		env = "dev"
	}

	err = godotenv.Load(".env." + env)
	if err != nil {
		appLogger.Fatal("Error loading .env." + env + " file")
	}

	appSettings.AppTitle = os.Getenv("APP_TITLE")
	appSettings.AppVersion = os.Getenv("APP_VERSION")
	appSettings.AppEnvironment = os.Getenv("APP_ENV")
	appSettings.AppHost = os.Getenv("APP_HOST_ADDRESS")
	appSettings.AppPort = os.Getenv("APP_HOST_PORT")
	appSettings.DbHost = os.Getenv("APP_DB_HOST")
	appSettings.DbPort = os.Getenv("APP_DB_PORT")
	appSettings.DbPass = os.Getenv("APP_DB_PASS")
	appSettings.DbUser = os.Getenv("APP_DB_USER")
	appSettings.DbName = os.Getenv("APP_DB_NAME")
	appSettings.DbSslMode = os.Getenv("APP_DB_SSLM")

	if appSettings.AppEnvironment == "test" {
		// Show application environment variables
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			if strings.HasPrefix(pair[0], "APP_") {
				appLogger.Info(fmt.Sprintf("%s: %s", pair[0], pair[1]))
			}
		}
	}
}

func main() {
	a := app.New(appLogger, &appSettings)

	// Seed Sample Data
	if appSettings.AppEnvironment != "prod" {
		//seedData(a.GetDB())
	}

	a.Run()
}

func seedData(appDB *gorm.DB) {

	createAuthor := domain.Author{Name: "Hasan SAHIN"}

	// Create
	r0 := appDB.Create(&createAuthor)
	if r0.Error != nil {
		fmt.Println("CREATE ERROR : " + r0.Error.Error())
	}
	id := createAuthor.ID // uuid.NewV4()

	fmt.Println("CREATED AUTHOR ID: " + id.String())

	//// Read
	//var author domain.Author
	//result := appDB.First(&author, "id = ?", id.String()) // find author with integer primary key
	//// Check if returns RecordNotFound error
	//errors.Is(result.Error, gorm.ErrRecordNotFound)
	//
	//appDB.First(&author, "name = ?", "D42") // find author with name D42

	//// Update - update product's price to 200
	//r1 := appDB.Updates(&product)
	//if r1.Error != nil {
	//	fmt.Println("UPDATE ERROR1 : " + r1.Error.Error())
	//}
	//
	//r2 := appDB.Model(&product).Update("Price", 102)
	//if r2.Error != nil {
	//	fmt.Println("UPDATE ERROR2 : " + r2.Error.Error())
	//}
	//
	//// Update - update multiple fields
	//r3 := appDB.Model(&product).Updates(domain.Product{Price: 222, Name: "F42"}) // non-zero fields
	//if r3.Error != nil {
	//	fmt.Println("UPDATE ERROR3 : " + r3.Error.Error())
	//}
	//
	//r4 := appDB.Model(&product).Updates(map[string]interface{}{"Price": 444, "Name": "F42"})
	//if r4.Error != nil {
	//	fmt.Println("UPDATE ERROR4 : " + r4.Error.Error())
	//}
	//
	////Delete - delete product
	//r5 := appDB.Delete(&domain.Product{}, "id = ?", id.String())
	//if r5.Error != nil {
	//	fmt.Println("DELETE ERROR : " + r5.Error.Error())
	//}
}

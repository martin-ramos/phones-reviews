package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	migration "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/joho/godotenv"
	"github.com/martin-ramos/phones-reviews/internals/database"
	"github.com/martin-ramos/phones-reviews/internals/logs"
)

const (
	migrationsRootFolder     = "file://migrations"
	migrationsScriptsVersion = 1
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	_ = logs.InitLogger()
	// Seteo de las variables para
	// la conexión a la base de datos
	user := goDotEnvVariable("USER_MYSQL")
	password := goDotEnvVariable("PASSWORD_MYSQL")
	host := goDotEnvVariable("HOST_MYSQL")
	port := goDotEnvVariable("PORT_MYSQL")
	db := goDotEnvVariable("DB_MYSQL")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, db)
	client := database.NewSqlClient(connection)
	doMigrate(client, "phones_review")
}

func doMigrate(client *database.MySqlClient, dbName string) {
	driver, _ := migration.WithInstance(client.DB, &migration.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationsRootFolder,
		dbName,
		driver,
	)

	if err != nil {
		logs.Log().Error(err.Error())
		return
	}

	current, _, _ := m.Version()
	logs.Log().Infof("current migrations version in %d", current)
	err = m.Migrate(migrationsScriptsVersion)
	if err != nil && err.Error() == "no change" {
		logs.Log().Info("no migration needed")
	}
}

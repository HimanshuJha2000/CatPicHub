package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/DevtronLabs/CatPicHub/internal/config"
	"github.com/DevtronLabs/CatPicHub/internal/constants"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pressly/goose"
	"log"
	"os"
)

var (
	flags = flag.NewFlagSet("migrate", flag.ExitOnError)
	dir   = flags.String(constants.MigrationDir, constants.DefaultMigrationDir, "directory with migration files")
	env   = flags.String(constants.Env, constants.Development, "Application env : prod/dev")
)

const (
	POSTGRES = "postgres"
)

func main() {
	// Get the database connection string
	config.LoadConfig(*env)
	driver := POSTGRES

	var databaseConfig config.Database

	databaseConfig = config.GetConfig().Database

	dbstring := databaseConfig.URL()
	if err := goose.SetDialect(driver); err != nil {
		log.Fatalln("goose run: %v", err)
	}

	db, err := sql.Open(driver, dbstring)
	if err != nil {
		log.Fatalln("-dbstring=%q: %v\n", dbstring, err)
	}

	defer db.Close()

	// Get the current working directory
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Specify the folder containing your migration files
	migrationsDir := fmt.Sprintf("%s/migrations", workingDir)

	// Run the migrations
	err = goose.Up(db, migrationsDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migrations ran successfully!")
}

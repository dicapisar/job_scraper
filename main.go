package main

import (
	"github.com/dicapisar/job_scraper/api"
	"github.com/dicapisar/job_scraper/database"
	"github.com/dicapisar/job_scraper/domain"
	"github.com/dicapisar/job_scraper/infra"
	repository2 "github.com/dicapisar/job_scraper/repository"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := database.NewConnection(config)

	if err != nil {
		log.Fatal("could not load database: " + err.Error())
	}

	err = domain.MigrateLinkedinJob(db)

	if err != nil {
		log.Fatal("could not migrate db")
	}

	repositoryGenerated := &repository2.Repository{
		DB: db,
	}

	infra.DBRepository = repositoryGenerated

	api.InitializeApi()
}

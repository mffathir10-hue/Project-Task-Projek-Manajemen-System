package main

import (
	"database/sql"
	"fmt"
	"gintugas/database"
	routers "gintugas/modules"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var (
	db     *sql.DB
	gormDB *gorm.DB
	err    error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Berhasil Koneksi Ke Database")

	gormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to create GORM connection:", err)
	}

	database.DBMigrate(db)

	InitiateRouter(db, gormDB)
}

func InitiateRouter(db *sql.DB, gormDB *gorm.DB) {
	router := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	routers.Initiator(router, db, gormDB)

	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}

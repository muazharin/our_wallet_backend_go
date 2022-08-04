package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/muazharin/our_wallet_backend_go/src/entities/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Gagal load .env file")
	}
	dbHost := os.Getenv("MySQL_HOST")
	dbUser := os.Getenv("MySQL_USER")
	dbPass := os.Getenv("MySQL_PASSWORD")
	dbName := os.Getenv("MySQL_DB_NAME")
	dbPort := os.Getenv("MySQL_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create connection to database")
	}
	fmt.Println("Berhasil terhubung ke MySql")
	db.AutoMigrate(database.Users{}, database.Wallets{}, database.OurWallet{}, database.Category{}, database.Transaction{}, database.TransactionFile{})
	return db
}

// database.go
package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB // Mengubah nama variabel menjadi DB (huruf kapital)

func ConnectDB() {
	var err error
	dsn := "myuser:secret@tcp(213.210.37.2:3306)/mydatabase"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error membuka koneksi ke database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Tidak dapat terhubung ke database: %v", err)
	}

	log.Println("Koneksi ke database berhasil!")
}

func CloseDB() {
	DB.Close()
}

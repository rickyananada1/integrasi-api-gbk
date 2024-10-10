// main.go
package main

import (
	"integrasi-gbk-online/config"
	"integrasi-gbk-online/scheduler"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Koneksi ke database
	config.ConnectDB()
	defer config.CloseDB()

	// Jalankan scheduler
	scheduler.StartScheduler()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Println("Program dihentikan.")
}

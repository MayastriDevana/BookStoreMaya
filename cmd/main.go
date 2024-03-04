package main

import (
	"bcas/bookstore-go/internals/routes"
	"bcas/bookstore-go/pkg"
	"log"
)

//Dependency Injection (DI) menggunakan parameter dari luar

func main() {
	// inisialisasi DB
	db, err := pkg.InitMySql()
	if err != nil {
		log.Fatal(err)
		// return
	}
	// inisialisasi Router
	router := routes.InitRouter(db)
	// inisialisasi Server
	server := pkg.InitServer(router)
	// jalankan server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

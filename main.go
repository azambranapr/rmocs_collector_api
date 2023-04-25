// Package main es el paquete principal
// en donde se ejecuta nuestra aplicación
// sss parsing member file
package main

import (
	"fmt"
	"github.com/azambranapr/rmocs_collector_api/src/core/services"
	"github.com/azambranapr/rmocs_collector_api/src/presentation/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome RMOCS 2023")

	var err = services.LoadFiles("certificates/app_rmocs_api.rsa", "certificates/app_rmocs_api.rsa.pub")
	if err != nil {
		log.Fatalf("error in certificates: %v", err)
	}

	err = godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(os.Getenv("PORT"))
	fmt.Println(os.Getenv("GIN_MODE"))

	router := routers.NewRouter()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	port := os.Getenv("PORT")
	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)
	err = router.Run(address)

	if err != nil {
		fmt.Printf("Error in starat router err: %v", err)
		return
	}

}

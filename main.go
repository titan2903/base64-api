package main

import (
	"base64-api/internal/routes"
	"base64-api/pkg/utils"
	"fmt"
)

func main() {
	e := routes.NewRoutes()
	port := utils.GoDotEnvVariable("PORT")
	if port == "" {
		port = "3000"
	}

	sPort := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(sPort))
	fmt.Printf("Successfully started on port %s\n", port)
}

package main

import (
	"fmt"
	"os"

	// "github.com/joho/godotenv"

)

func main() {
	fmt.Println("hello world")

	// loading the env values in local 
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("err loading: %v", err)
	// }
	// get the env value
	port := os.Getenv("PORT")
	fmt.Printf("hello world %s", port)
	if port == "" {
		port = "8080"
	}
	// fmt.Printf("hello world %s", port)

}

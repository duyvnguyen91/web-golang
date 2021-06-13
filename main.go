package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/duyvnguyen91/web-golang/server/routes"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error getting env, %v, err")
	}

	router := httprouter.New()

	router.POST("/auth/loging", routes.Login)
	router.POST("/auth/register", routes.Register)

	fmt.Println("Listening to port 8000")
	log.Fatal((http.ListenAndServe(":8000", router)))
}

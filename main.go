package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")
	godotenv.Load(".env")

    port := os.Getenv("PORT")
    // if port == "" {
    //     fmt.Println("PORT not set")
    // } else {
    //     fmt.Println("PORT:", port)
    // }

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, 
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", readinessHandler)
	v1Router.Get("/error", errorHandler)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	log.Printf("Server started on port %v", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
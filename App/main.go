package main

import (
	_ "apiGo/docs"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title API de Lista Aleatória com Swagger
// @description Esta API retorna uma lista aleatória de inteiros.
// @BasePath /
func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/random-list", randomListHandler) // Novo endpoint
	http.Handle("/swagger/", httpSwagger.WrapHandler)  // Swagger endpoint
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// @Summary Retorna uma lista aleatória
// @Description Retorna uma lista aleatória de inteiros.
// @Accept  json
// @Produce  json
// @Success 200 {array} int
// @Router /random-list [get]
func randomListHandler(w http.ResponseWriter, r *http.Request) {
	list := make([]int, 5)
	for i := range list {
		list[i] = rand.Intn(100)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

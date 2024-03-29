package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Pokemon struct {
	Name string
	Id   int
}

func main() {
	http.HandleFunc("GET /api/pokemon", func(w http.ResponseWriter, r *http.Request) {
		var pokemon = Pokemon{
			Name: "Pikapika",
			Id:   2,
		}
		okeBytes, err := json.MarshalIndent(pokemon, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(418)
		w.Write(okeBytes)
	})

	handler := http.FileServer(http.Dir("www"))

	http.Handle("/", handler)

	templateHandler := template.Must(template.ParseFiles("www/custom.html"))

	http.HandleFunc("/custom", func(writer http.ResponseWriter, request *http.Request) {
		type PageData struct {
			Title    string
			Pokemons []Pokemon
		}
		var data = PageData{
			Title: "Mes pokemons" + string(rand.Intn(1000)),
			Pokemons: []Pokemon{
				{Id: 2, Name: "Pikachu"},
				{Id: 3, Name: "Bulbi"},
			},
		}
		templateHandler.Execute(writer, data)
	})

	fmt.Println("Démarrage du serveur sur le port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

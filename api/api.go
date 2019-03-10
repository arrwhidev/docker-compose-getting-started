package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type director struct {
	Name string `json:"name"`
}

type film struct {
	Title    string   `json:"title"`
	Year     int      `json:"year"`
	Director director `json:"director"`
}

var spielberg = director{"Steven Spielberg"}
var lucas = director{"George Lucas"}
var kershner = director{"Irvin Kershner"}
var marquand = director{"Richard Marquand"}
var films = [...]film{
	{"Indiana Jones: Raiders of the Lost Ark", 1981, spielberg},
	{"Indiana Jones: Temple of Doom", 1984, spielberg},
	{"Indiana Jones: The Last Crusade", 1989, spielberg},
	{"Star Wars: Episode IV - A New Hope", 1977, lucas},
	{"Star Wars: Episode V - The Empire Strikes Back", 1980, kershner},
	{"Star Wars: Episode VI - Return of the Jedi", 1983, marquand},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        j, _ := json.Marshal(films)
        w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}

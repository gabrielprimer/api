package main

import (
	"encoding/json"
	"net/http"
)

// Definição da estrutura do Anime
type Anime struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Genre    string `json:"genre"`
	Synopsis string `json:"synopsis"`
}

// Função handler exportada para a rota /api/anime
func AnimeHandler(w http.ResponseWriter, r *http.Request) {  // Mudança aqui: AnimeHandler em vez de animeHandler
	// Criando um exemplo de anime
	anime := Anime{
		Name:     "Naruto",
		Year:     2002,
		Genre:    "Ação, Aventura",
		Synopsis: "Naruto Uzumaki, um jovem ninja, busca reconhecimento e sonha em se tornar o Hokage.",
	}

	// Define o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")

	// Retorna o JSON do anime
	json.NewEncoder(w).Encode(anime)
}

func Handler() {
	// Mapeia a rota /api/anime para a função AnimeHandler
	http.HandleFunc("/api/anime", AnimeHandler)

	// Inicia o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}

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

// Função handler para a rota /api/anime
func animeHandler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	// Mapeia a rota /api/anime para a função animeHandler
	http.HandleFunc("/api/anime", animeHandler)

	// Inicia o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}

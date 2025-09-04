package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Jogo struct {
	ID     string `json:"id"`
	Titulo string `json:"titulo"`
	Plataforma string `json:"plataforma"`
	Nota float32 `json:"nota"`
}

var jogos = []Jogo{
	{ID: "1", Titulo: "Dark Souls Remastered", Plataforma: "PC", Nota: 9.3},
	{ID: "2", Titulo: "Bloodborne", Plataforma: "PlayStation 4", Nota: 9.3},
	{ID: "3", Titulo: "Elden Ring", Plataforma: "PC", Nota: 9.4},
	{ID: "4", Titulo: "Sekiro: Shadows Die Twice", Plataforma: "PlayStation 4", Nota: 9.5},
	{ID: "5", Titulo: "Demon's Souls", Plataforma: "PlayStation 5", Nota: 9.4},
	{ID: "6", Titulo: "Dark Souls II", Plataforma: "PlayStation 3", Nota: 10.0},

}

func getJogos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, jogos)
}

func getJogosByID(c *gin.Context) {
	id:= c.Param("id")

	for _, a := range jogos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "jogo não encontrado"})
}

func postJogos(c *gin.Context) {
	var novoJogo Jogo

	if err := c.BindJSON(&novoJogo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	lastID := len(jogos) + 1
	novoJogo.ID = strconv.Itoa(lastID)

	jogos = append(jogos, novoJogo)

	c.IndentedJSON(http.StatusCreated, novoJogo)
}

func updateJogos(c *gin.Context) {
	id := c.Param("id")
	var jogoAtualizado Jogo

	if err := c.BindJSON(&jogoAtualizado); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	for i, a := range jogos {
		if a.ID == id {
			jogoAtualizado.ID = id
			jogos[i] = jogoAtualizado
			c.IndentedJSON(http.StatusOK, jogoAtualizado)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "jogo não encontrado"})
}

func deleteJogos(c *gin.Context) {
	id := c.Param("id")

	for i, a := range jogos {
		if a.ID == id {
			jogos = append(jogos[:i], jogos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "jogo deletado"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "jogo não encontrado"})
}

func main(){
	router := gin.Default()

	router.GET("/jogos", getJogos)
	router.GET("/jogos/:id", getJogosByID)
	router.POST("/jogos", postJogos)
	router.PUT("/jogos/:id", updateJogos)
	router.DELETE("/jogos/:id", deleteJogos)		

	router.Run("localhost:8080")
}


package main

import (
	// "fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerRaiz(c *gin.Context) {
	c.String(http.StatusOK, "Bienvenido al root")
}

func HandlerGophers(c *gin.Context) {

	type respuesta struct {
		Gophers []string `json:"gophers"`
	}

	r := respuesta{Gophers: []string{"gopher 1", "gopher 2"}}
	c.JSON(http.StatusOK, r)
}

func HandlerInfoGopher(c *gin.Context) {
	c.String(http.StatusOK, "Gopher es la mascota de Golang!")
}

func HandlerAbout(c *gin.Context) {
	c.String(http.StatusOK, "este endpoint nos va a devolver info sobre nuestra app")
}

func main() {
	router := gin.Default()
	//Cada vez que llamamos a GET y le pasamos una ruta, definimos un nuevo endpoint.
	router.GET("/", HandlerRaiz)
	router.GET("/gophers", HandlerGophers)
	router.GET("/gophers/info", HandlerInfoGopher)
	router.GET("/about", HandlerAbout)
	router.Run(":8080")
}

package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hola", func(c *gin.Context) {
		c.String(http.StatusOK, "Hola!")
	})
	r.POST("/hola", func(c *gin.Context) {
		//El body, header y method están contenidos en el contexto de gin.
		body := c.Request.Body
		header := c.Request.Header
		metodo := c.Request.Method

		fmt.Println("¡He recibido algo!")
		fmt.Printf("\tMetodo: %s\n", metodo)
		fmt.Printf("\tContenido del header:\n")

		for key, value := range header {
			fmt.Printf("\t\t%s -> %s\n", key, value)
		}
		fmt.Printf("\tEl body es un io.ReadCloser:(%s), y para trabajar con el vamos a tener que leerlo luego\n", body)
		c.String(200, "¡Lo recibí!") //Respondemos al cliente con 200 OK y un mensaje.
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Definimos nuestra estructura de información
// y agregamos las etiquetas correspondientes para poder realizar el unmarshalling
type Empleado struct {
	// Una etiqueta de struct se cierra con caracteres de acento grave `
	Nombre   string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
	Id       string `form:"id" json:"id"`
	Activo   string `form:"active" json:"active" binding:"required"`
}

func AutorizarEmpleado(ctxt *gin.Context) {

	var empleado Empleado
	// el metodo ShouldBindJSON de nuestro context, asociará el contenido del body
	// a los campos de la estructura que le proveamos
	if err := ctxt.ShouldBindJSON(&empleado); err != nil {
		ctxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%+v\n", empleado)
	if empleado.Nombre != "user1" || empleado.Password != "123" {
		ctxt.JSON(http.StatusUnauthorized, gin.H{"status": "no esta autorizado"})
		return
	}

	ctxt.JSON(http.StatusOK, gin.H{"status": "autorizado"})
}

func main() {
	r := gin.Default()
	r.POST("/auth", AutorizarEmpleado)
	r.Run()
}

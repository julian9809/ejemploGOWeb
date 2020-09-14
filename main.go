package main

import (
	"log"

	"github.com/julian9809/ejemploGOWeb/bd"
	"github.com/julian9809/ejemploGOWeb/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}

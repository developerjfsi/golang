package main

import (
	"log"

	"github.com/developerjfsi/golang/bd"
	"github.com/developerjfsi/golang/handlers"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}
	handlers.Manejadores()

}

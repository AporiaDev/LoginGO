package main

import (
	"log"
	"log/config"
	"log/internal/database"
)
    


func main() {
     
    cfg := config.LoadConfig()
    db, _ := database.ConectarDB(cfg)
    //Base de datos se cierra cuando el servidor se apaga
    defer db.Close()

    log.Println("Listo para recibir peticiones")


}
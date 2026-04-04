package main

import (
	
	"log"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors" 
    "time"

	"log/config"
	"log/internal/database"
	"log/internal/handlers"
	"log/internal/repository"
	"log/internal/service"
)

func main() {
	
	cfg := config.LoadConfig()

	conexionDB, err := database.ConectarDB(cfg)
	if err != nil {
		log.Fatalf("Error fatal: No se pudo conectar a PostgreSQL: %v", err)
	}
	defer conexionDB.Close() // Asegura que la base de datos se cierre al apagar el servidor

    //LA INYECCIÓN DE DEPENDENCIAS (El Cableado)
	//1
	usuarioRepo := repository.NewUserRepository(conexionDB)
	
	// Construimos el Consultorio
	authService := service.NuevoAuthRepository(usuarioRepo)
	
	// C. Construimos la Recepción
	authHandler := handlers.NuevoAuthHandler(authService)

	
	// EL ENRUTAMIENTO (El Mapa de la URL al Código)
	
	r := gin.Default()

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, // En producción, usa la URL de tu frontend
        AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: false,
        MaxAge:           12 * time.Hour,
    }))
	// Aquí está la respuesta a tu pregunta:
	// Cuando entre una petición POST a "/api/login", Gin la dirigirá.
	r.POST("/api/login", authHandler.Login)
	r.POST("api/register", authHandler.Registro)

	// 5. Encendemos el servidor en el puerto 8080
	log.Println("Servidor de Laboratorio operando en el puerto 8080...")
	r.Run(":8080")
}
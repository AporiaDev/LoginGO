package handlers

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"

	"log/internal/service"
	"log/pkg/utils"
)

type AuthHandler struct{
	svc *service.AuthService
}

func NuevoAuthHandler(svc *service.AuthService)*AuthHandler{
	return &AuthHandler{svc:svc}
}

//Estructura para recibir el JSON
type LoginInput struct{
	Correo 	 string `json:"correo" bidding:"required,email"`
	Password string `json:"password" bidding:"required"`
}

type RegisterInput struct{
	Username  string 	`json:"username" bidding:"required"`
	Correo 	 string 	`json:"correo" bidding:"required,email"`
	Password string 	`json:"password" bidding:"required"`
}

func (h *AuthHandler) Registro(c *gin.Context){
	
	var registerInput RegisterInput

	if err := c.ShouldBindJSON(&registerInput); err != nil {
    		c.JSON(400, gin.H{"error": err.Error()})
    return
	}
	log.Printf("Datos recibidos: %+v\n", registerInput)
	// Crear usuario
	err := h.svc.NuevoUsuario(
		registerInput.Username,
		registerInput.Correo,
		registerInput.Password,
	)

	if err != nil{
		utils.Fail(c,http.StatusInternalServerError, "ERROR_CREACION", "ERROR_INTERNO" )
	}
	//Repuesta exitosa
	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario creado correctamente",
	})

}
func (h  *AuthHandler) Login(c *gin.Context){
	//variable para guardar los valores del input
	var input LoginInput
	// 1. La Recepcionista lee el JSON
	if err := c.ShouldBindJSON(&input); err != nil {

		utils.Fail(c, http.StatusBadRequest, "DATOS_INVALIDOS", "Formato incorrecto")
		return
	}
	usuarioValidado,err := h.svc.Autenticar(input.Correo,input.Password)
	if err != nil{
		utils.Fail(c, http.StatusUnauthorized, "NO_AUTORIZADO", err.Error())
	}
	utils.OK(c, usuarioValidado)
}
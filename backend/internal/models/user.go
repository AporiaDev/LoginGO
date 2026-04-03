package models


type User struct {
    ID       int    `json:"id"`
    Nombre   string `json:"username"`
    Password string `json:"password"`
    Correo   string `json:"correo"`
}


package handlers

import (
	"net/http"

	"github.com/JGurus/template-initial-api/models"
	"github.com/labstack/echo/v4"
)

//Register struct
type Register struct {
	service UserHandler
}

//NewRegister .
func NewRegister(service UserHandler) Register {
	return Register{service}
}

//Signup .
func (r *Register) Signup(c echo.Context) error {
	data := models.Register{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "Estructura no válida", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if r.userExist(&data) {
		response := newResponse(Error, "El usuario ya existe", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	user := models.User{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}
	err = r.service.Create(&user)
	if err != nil {
		response := newResponse(Error, "No se pudo crear el usuario", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "Te has registrado exitosamente", nil)
	return c.JSON(http.StatusCreated, response)
}

func (r *Register) userExist(data *models.Register) bool {
	_, err := r.service.GetByUsername(data.Username)
	if err != nil {
		return false
	}
	_, err = r.service.GetByEmail(data.Email)
	if err != nil {
		return false
	}
	return true
}

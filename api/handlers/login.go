package handlers

import (
	"log"
	"net/http"

	"github.com/JGurus/template-initial-api/auth"
	"github.com/JGurus/template-initial-api/models"
	"github.com/labstack/echo/v4"
)

//Login struct
type Login struct {
	service UserHandler
}

//NewLogin .
func NewLogin(service UserHandler) Login {
	return Login{service}
}

//LogIn .
func (l *Login) LogIn(c echo.Context) error {
	data := models.Login{}
	err := c.Bind(&data)
	if err != nil {
		log.Fatal(err)
		response := newResponse(Error, "Estructura no válida", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = l.isLoginValid(&data)
	if err != nil {
		response := newResponse(Error, "Email o contraseña no válidos", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	token, err := auth.GenerateToken(&data)
	if err != nil {
		log.Fatal(err)
		response := newResponse(Error, "No se pudo generar el token", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	dataToken := map[string]string{"token": token}

	response := newResponse(Message, "OK", dataToken)
	return c.JSON(http.StatusOK, response)
}

func (l *Login) isLoginValid(data *models.Login) error {
	user, err := l.service.GetByEmail(data.Email)
	if err != nil {
		return err
	}
	err = user.ValidatedPassword(data.Password)
	if err != nil {
		return err
	}
	return nil
}

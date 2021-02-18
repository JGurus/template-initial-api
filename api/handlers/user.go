package handlers

import (
	"net/http"
	"strconv"

	"github.com/JGurus/template-initial-api/models"
	"github.com/labstack/echo/v4"
)

//User .
type User struct {
	service UserHandler
}

//NewUser .
func NewUser(service UserHandler) User {
	return User{service}
}

//Create .
func (h *User) Create(c echo.Context) error {
	data := models.User{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "Incorrect JSON structure", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	err = h.service.Create(&data)
	if err != nil {
		response := newResponse(Error, "The user could not be created", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "The user has been created successfully", nil)
	return c.JSON(http.StatusCreated, response)
}

//Update .
func (h *User) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "The id must be a positive integer", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	user := models.User{}
	err = c.Bind(&user)
	if err != nil {
		response := newResponse(Error, "Incorrect JSON structure", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	err = h.service.Update(id, &user)
	if err != nil {
		response := newResponse(Error, "Could not update user", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "The user has been updated successfully", nil)
	return c.JSON(http.StatusOK, response)
}

//Delete .
func (h *User) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "The id must be a positive integer", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = h.service.Delete(id)
	if err != nil {
		response := newResponse(Error, "Could not delete user", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "User has been removed successfully", nil)
	return c.JSON(http.StatusOK, response)
}

//GetAll .
func (h *User) GetAll(c echo.Context) error {
	users, err := h.service.GetAll()
	if err != nil {
		response := newResponse(Error, "Could not bring users", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "Successful query", users)
	return c.JSON(http.StatusOK, response)
}

//GetByID .
func (h *User) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "The id must be a positive integer", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	user, err := h.service.GetByID(id)
	if err != nil {
		response := newResponse(Error, "Could not bring user", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "Successful query", user)
	return c.JSON(http.StatusOK, response)
}

//GetByUsername .
func (h *User) GetByUsername(c echo.Context) error {
	username := c.Param("username")
	user, err := h.service.GetByUsername(username)
	if err != nil {
		response := newResponse(Error, "No se pudo traer al usuario", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "Consulta sastifactoria", user)
	return c.JSON(http.StatusOK, response)
}

//GetByEmail .
func (h *User) GetByEmail(c echo.Context) error {
	email := c.Param("email")
	user, err := h.service.GetByEmail(email)
	if err != nil {
		response := newResponse(Error, "No se pudo traer al usuario", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "Consulta sastifactoria", user)
	return c.JSON(http.StatusOK, response)
}

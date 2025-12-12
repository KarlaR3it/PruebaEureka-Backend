package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KarlaR3it/PruebaEureka-Backend/internal/models"
	"github.com/KarlaR3it/PruebaEureka-Backend/internal/service"
)

type PersonHandler struct {
	service service.PersonService
}

func NewPersonHandler(service service.PersonService) *PersonHandler {
	return &PersonHandler{service: service}
}

// POST /persons
func (h *PersonHandler) Create(c *gin.Context) {
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		response := models.NewErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if err := h.service.CreatePerson(c.Request.Context(), &person); err != nil {
		response := models.NewErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := models.NewSuccessResponse(person)
	c.JSON(http.StatusCreated, response)
}

// GET /persons
func (h *PersonHandler) GetAll(c *gin.Context) {
	persons, err := h.service.GetAllPersons(c.Request.Context())
	if err != nil {
		response := models.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := models.NewSuccessResponse(persons)
	c.JSON(http.StatusOK, response)
}

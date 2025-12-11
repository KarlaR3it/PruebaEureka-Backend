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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreatePerson(c.Request.Context(), &person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, person)
}

// GET /persons
func (h *PersonHandler) GetAll(c *gin.Context) {
	persons, err := h.service.GetAllPersons(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, persons)
}

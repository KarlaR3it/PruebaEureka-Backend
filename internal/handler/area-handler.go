package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KarlaR3it/PruebaEureka-Backend/internal/models"
	"github.com/KarlaR3it/PruebaEureka-Backend/internal/service"
)

type AreaHandler struct {
	service service.AreaService
}

func NewAreaHandler(service service.AreaService) *AreaHandler {
	return &AreaHandler{service: service}
}

// POST /areas
func (h *AreaHandler) Create(c *gin.Context) {
	var area models.Area
	if err := c.ShouldBindJSON(&area); err != nil {
		response := models.NewErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if err := h.service.CreateArea(c.Request.Context(), &area); err != nil {
		response := models.NewErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := models.NewSuccessResponse(area)
	c.JSON(http.StatusCreated, response)
}

// GET /areas
func (h *AreaHandler) GetAll(c *gin.Context) {
	areas, err := h.service.GetAllAreas(c.Request.Context())
	if err != nil {
		response := models.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := models.NewSuccessResponse(areas)
	c.JSON(http.StatusOK, response)
}

// GET /areas/count
func (h *AreaHandler) GetAreaCounts(c *gin.Context) {
	counts, err := h.service.GetAreaCounts(c.Request.Context())
	if err != nil {
		response := models.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := models.NewSuccessResponse(counts)
	c.JSON(http.StatusOK, response)
}

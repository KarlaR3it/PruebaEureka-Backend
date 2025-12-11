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
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.service.CreateArea(c.Request.Context(), &area); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, area)
}

// GET /areas
func (h *AreaHandler) GetAll(c *gin.Context) {
    areas, err := h.service.GetAllAreas(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, areas)
}

// GET /areas/count
func (h *AreaHandler) GetAreaCounts(c *gin.Context) {
    counts, err := h.service.GetAreaCounts(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, counts)
}

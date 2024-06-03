package controllers

import (
	"net/http"
	"strconv"

	"srating/bootstrap"
	"srating/domain"
	"srating/x/rest"

	"github.com/gin-gonic/gin"
)

type LocationController struct {
	LocationService domain.LocationService
	Env             *bootstrap.Env
	*rest.JSONRender
}

// CreateLocation
// @Router /locations [post]
// @Tags location
// @Query body domain.Location
// @Param payload body domain.Location true "payload"
// @Summary Create location
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *LocationController) CreateLocation(c *gin.Context) {
	var input = &domain.Location{}
	err := c.ShouldBindJSON(&input)
	rest.AssertNil(err)
	err = t.LocationService.CreateLocation(c, input)
	rest.AssertNil(err)
	t.Success(c)
}


// GetAllLocation
// @Router /locations [get]
// @Tags location
// @Summary Get all location
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param user_id query int false "user_id"
// @Param level query int false "level"
// @Param start_date query int false "start_date"
// @Param end_date query int false "end_date"
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *LocationController) GetAllLocation(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	input := domain.GetAllLocationRequest{
		PaginationRequest: domain.PaginationRequest{
			Limit: limit,
			Page:  page,
		},
	}
	total, totalCount, result, err := t.LocationService.GetAllLocation(c, input)
	rest.AssertNil(err)
	t.SendCustomData(c, map[string]interface{}{
		"status":     "success",
		"data":       result,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalCount": totalCount,
	},
	)
}

// GetLocationDetail
// @Router /locations/:id [get]
// @Tags location
// @Summary Get location by detail
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *LocationController) GetLocationDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid location ID"})
		return
	}

	result, err := t.LocationService.GetLocationDetail(c, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": result})
}

// UpdateLocation
// @Router /locations [put]
// @Tags location
// @Summary Update location
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *LocationController) UpdateLocation(c *gin.Context) {
	input := &domain.Location{}
	rest.AssertNil(c.ShouldBindJSON(&input))
	err := t.LocationService.UpdateLocation(c, input)
	rest.AssertNil(err)
	t.Success(c)
}

// DeleteLocation
// @Router /locations/:id [delete]
// @Tags location
// @Summary Delete location
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *LocationController) DeleteLocation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	rest.AssertNil(err)
	err = t.LocationService.DeleteLocation(c, uint(id))
	rest.AssertNil(err)
	t.Success(c)
}
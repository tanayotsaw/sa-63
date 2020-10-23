package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/USER/app/ent"
	"github.com/USER/app/ent/section"
	"github.com/gin-gonic/gin"
)

// SectionController defines the struct for the section controller
type SectionController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateSection handles POST requests for adding section entities
// @Summary Create section
// @Description Create section
// @ID create-section
// @Accept   json
// @Produce  json
// @Param section body ent.Section true "Section entity"
// @Success 200 {object} ent.Section
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /sections [post]
func (ctl *SectionController) CreateSection(c *gin.Context) {
	obj := ent.Section{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "section binding failed",
		})
		return
	}

	u, err := ctl.client.Section.
		Create().
		SetGroup(obj.Group).
		SetRecieve(obj.Recieve).
		SetRoom(obj.Room).
		SetSeatLeft(obj.SeatLeft).
		SetDateTime(obj.DateTime).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetSection handles GET requests to retrieve a Section entity
// @Summary Get a section entity by ID
// @Description get section by ID
// @ID get-section
// @Produce  json
// @Param id path int true "Section ID"
// @Success 200 {object} ent.Section
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /sections/{id} [get]
func (ctl *SectionController) GetSection(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Section.
		Query().
		Where(section.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListSection handles request to get a list of section entities
// @Summary List section entities
// @Description list section entities
// @ID list-section
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Section
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Sections [get]
func (ctl *SectionController) ListSection(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	sections, err := ctl.client.Section.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, sections)
}

// DeleteSection handles DELETE requests to delete a section entity
// @Summary Delete a section entity by ID
// @Description get section by ID
// @ID delete-section
// @Produce  json
// @Param id path int true "Section ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /sections/{id} [delete]
func (ctl *SectionController) DeleteSection(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Section.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateSection handles PUT requests to update a Section entity
// @Summary Update a section entity by ID
// @Description update section by ID
// @ID update-section
// @Accept   json
// @Produce  json
// @Param id path int true "Section ID"
// @Param section body ent.Section true "Section entity"
// @Success 200 {object} ent.Section
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /sections/{id} [put]
func (ctl *SectionController) UpdateSection(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Section{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "section binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Section.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewSectionController creates and registers handles for the section controller
func NewSectionController(router gin.IRouter, client *ent.Client) *SectionController {
	uc := &SectionController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitSectionController registers routes to the main engine
func (ctl *SectionController) register() {
	sections := ctl.router.Group("/sections")

	sections.GET("", ctl.ListSection)

	// CRUD
	sections.POST("", ctl.CreateSection)
	sections.GET(":id", ctl.GetSection)
	sections.PUT(":id", ctl.UpdateSection)
	sections.DELETE(":id", ctl.DeleteSection)
}

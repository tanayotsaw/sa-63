package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/USER/app/ent"
	"github.com/USER/app/ent/subject"
	"github.com/gin-gonic/gin"
)

// SubjectController defines the struct for the subject controller
type SubjectController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateSubject handles POST requests for adding subject entities
// @Summary Create subject
// @Description Create subject
// @ID create-subject
// @Accept   json
// @Produce  json
// @Param subject body ent.Subject true "Subject entity"
// @Success 200 {object} ent.Subject
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subjects [post]
func (ctl *SubjectController) CreateSubject(c *gin.Context) {
	obj := ent.Subject{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "subject binding failed",
		})
		return
	}

	u, err := ctl.client.Subject.
		Create().
		SetSubjectName(obj.SubjectName).
		SetCredit(obj.Credit).
		SetPrice(obj.Price).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetSubject handles GET requests to retrieve a subject entity
// @Summary Get a subject entity by ID
// @Description get subject by ID
// @ID get-subject
// @Produce  json
// @Param id path int true "Subject ID"
// @Success 200 {object} ent.Subject
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subjects/{id} [get]
func (ctl *SubjectController) GetSubject(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Subject.
		Query().
		Where(subject.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListSubject handles request to get a list of subject entities
// @Summary List subject entities
// @Description list subject entities
// @ID list-subject
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Subject
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subjects [get]
func (ctl *SubjectController) ListSubject(c *gin.Context) {
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

	subjects, err := ctl.client.Subject.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, subjects)
}

// DeleteSubject handles DELETE requests to delete a Subject entity
// @Summary Delete a subject entity by ID
// @Description get subject by ID
// @ID delete-subject
// @Produce  json
// @Param id path int true "subject ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subjects/{id} [delete]
func (ctl *SubjectController) DeleteSubject(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Subject.
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

// UpdateSubject handles PUT requests to update a subject entity
// @Summary Update a subject entity by ID
// @Description update subject by ID
// @ID update-subject
// @Accept   json
// @Produce  json
// @Param id path int true "Subject ID"
// @Param subject body ent.Subject true "Subject entity"
// @Success 200 {object} ent.Subject
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subjects/{id} [put]
func (ctl *SubjectController) UpdateSubject(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Subject{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Subject binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Subject.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewSubjectController creates and registers handles for the subject controller
func NewSubjectController(router gin.IRouter, client *ent.Client) *SubjectController {
	uc := &SubjectController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitSubjectController registers routes to the main engine
func (ctl *SubjectController) register() {
	subjects := ctl.router.Group("/subjects")

	subjects.GET("", ctl.ListSubject)

	// CRUD
	subjects.POST("", ctl.CreateSubject)
	subjects.GET(":id", ctl.GetSubject)
	subjects.PUT(":id", ctl.UpdateSubject)
	subjects.DELETE(":id", ctl.DeleteSubject)
}

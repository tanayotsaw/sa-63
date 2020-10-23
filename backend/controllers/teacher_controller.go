package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/USER/app/ent"
	"github.com/USER/app/ent/teacher"
	"github.com/gin-gonic/gin"
)

// TeacherController defines the struct for the teacher controller
type TeacherController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateTeacher handles POST requests for adding teacher entities
// @Summary Create teacher
// @Description Create teacher
// @ID create-teacher
// @Accept   json
// @Produce  json
// @Param teacher body ent.Teacher true "Teacher entity"
// @Success 200 {object} ent.Teacher
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /teachers [post]
func (ctl *TeacherController) CreateTeacher(c *gin.Context) {
	obj := ent.Teacher{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "teacher binding failed",
		})
		return
	}

	t, err := ctl.client.Teacher.
		Create().
		SetTeacherName(obj.TeacherName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, t)
}

// GetTeacher handles GET requests to retrieve a teacher entity
// @Summary Get a teacher entity by ID
// @Description get teacher by ID
// @ID get-teacher
// @Produce  json
// @Param id path int true "Teacher ID"
// @Success 200 {object} ent.Teacher
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /teachers/{id} [get]
func (ctl *TeacherController) GetTeacher(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	t, err := ctl.client.Teacher.
		Query().
		Where(teacher.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, t)
}

// ListTeacher handles request to get a list of teacher entities
// @Summary List teacher entities
// @Description list teacher entities
// @ID list-teacher
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Teacher
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /teachers [get]
func (ctl *TeacherController) ListTeacher(c *gin.Context) {
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

	teachers, err := ctl.client.Teacher.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, teachers)
}

// DeleteTeacher handles DELETE requests to delete a teacher entity
// @Summary Delete a teacher entity by ID
// @Description get teacher by ID
// @ID delete-teacher
// @Produce  json
// @Param id path int true "Teacher ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /teachers/{id} [delete]
func (ctl *TeacherController) DeleteTeacher(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Teacher.
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

// UpdateTeacher handles PUT requests to update a Teacher entity
// @Summary Update a teacher entity by ID
// @Description update teacher by ID
// @ID update-teacher
// @Accept   json
// @Produce  json
// @Param id path int true "Teacher ID"
// @Param teacher body ent.Teacher true "Teacher entity"
// @Success 200 {object} ent.Teacher
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /teachers/{id} [put]
func (ctl *TeacherController) UpdateTeacher(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Teacher{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "teacher binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Teacher.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewTeacherController creates and registers handles for the teacher controller
func NewTeacherController(router gin.IRouter, client *ent.Client) *TeacherController {
	uc := &TeacherController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitTeacherController registers routes to the main engine
func (ctl *TeacherController) register() {
	teachers := ctl.router.Group("/teachers")

	teachers.GET("", ctl.ListTeacher)

	// CRUD
	teachers.POST("", ctl.CreateTeacher)
	teachers.GET(":id", ctl.GetTeacher)
	teachers.PUT(":id", ctl.UpdateTeacher)
	teachers.DELETE(":id", ctl.DeleteTeacher)
}

package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/USER/app/ent"
	"github.com/USER/app/ent/lessonplan"
	"github.com/USER/app/ent/section"
	"github.com/USER/app/ent/subject"
	"github.com/USER/app/ent/teacher"
	"github.com/gin-gonic/gin"
)

// LessonplanController defines the struct for the lessonplan controller
type LessonplanController struct {
	client *ent.Client
	router gin.IRouter
}

//Lessonplanss struct
type Lessonplan struct {
	GroupssID   int
	TeacherssID int
	SubjectssID int
	Room        string
}

// CreateLessonplan handles POST requests for adding lessonplan entities
// @Summary Create lessonplan
// @Description Create lessonplan
// @ID create-lessonplan
// @Accept   json
// @Produce  json
// @Param lessonplan body ent.Lessonplan true "Lessonplan entity"
// @Success 200 {object} ent.Lessonplan
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /lessonplans [post]
func (ctl *LessonplanController) CreateLessonplan(c *gin.Context) {
	obj := Lessonplan{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "lessonplan binding failed",
		})
		return
	}

	s, err := ctl.client.Section.
		Query().
		Where(section.IDEQ(int(obj.GroupssID))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "save sec not found",
		})
		return
	}

	j, err := ctl.client.Subject.
		Query().
		Where(subject.IDEQ(int(obj.SubjectssID))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "save subject not found",
		})
		return
	}

	t, err := ctl.client.Teacher.
		Query().
		Where(teacher.IDEQ(int(obj.TeacherssID))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "save teacher not found",
		})
		return
	}

	u, err := ctl.client.Lessonplan.
		Create().
		SetGroupID(s).
		SetCourseID(j).
		SetProfessorID(t).
		SetRoom(obj.Room).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetLessonplan handles GET requests to retrieve a lessonplan entity
// @Summary Get a lessonplan entity by ID
// @Description get lessonplan by ID
// @ID get-lessonplan
// @Produce  json
// @Param id path int true "Lessonplan ID"
// @Success 200 {object} ent.Lessonplan
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /lessonplans/{id} [get]
func (ctl *LessonplanController) GetLessonplan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Lessonplan.
		Query().
		Where(lessonplan.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListLessonplan handles request to get a list of lessonplan entities
// @Summary List lessonplan entities
// @Description list lessonplan entities
// @ID list-lessonplan
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Lessonplan
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /lessonplans [get]
func (ctl *LessonplanController) ListLessonplan(c *gin.Context) {
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

	lessonplans, err := ctl.client.Lessonplan.
		Query().
		WithGroupID().
		WithCourseID().
		WithProfessorID().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, lessonplans)
}

// DeleteLessonplan handles DELETE requests to delete a lessonplan entity
// @Summary Delete a lessonplan entity by ID
// @Description get lessonplan by ID
// @ID delete-lessonplan
// @Produce  json
// @Param id path int true "Lessonplan ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /lessonplans/{id} [delete]
func (ctl *LessonplanController) DeleteLessonplan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Lessonplan.
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

// UpdateLessonplan handles PUT requests to update a lessonplan entity
// @Summary Update a lessonplan entity by ID
// @Description update lessonplan by ID
// @ID update-lessonplan
// @Accept   json
// @Produce  json
// @Param id path int true "Lessonplan ID"
// @Param lessonplan body ent.Lessonplan true "Lessonplan entity"
// @Success 200 {object} ent.Lessonplan
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /lessonplans/{id} [put]
func (ctl *LessonplanController) UpdateLessonplan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Lessonplan{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "lessonplan binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Lessonplan.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewLessonplanController creates and registers handles for the lessonplan controller
func NewLessonplanController(router gin.IRouter, client *ent.Client) *LessonplanController {
	uc := &LessonplanController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitLessonplanController registers routes to the main engine
func (ctl *LessonplanController) register() {
	lessonplans := ctl.router.Group("/lessonplans")

	lessonplans.GET("", ctl.ListLessonplan)

	// CRUD
	lessonplans.POST("", ctl.CreateLessonplan)
	lessonplans.GET(":id", ctl.GetLessonplan)
	lessonplans.PUT(":id", ctl.UpdateLessonplan)
	lessonplans.DELETE(":id", ctl.DeleteLessonplan)
}

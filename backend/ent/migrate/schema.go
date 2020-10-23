// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// LessonplansColumns holds the columns for the "lessonplans" table.
	LessonplansColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "room", Type: field.TypeString},
		{Name: "section_sec_id", Type: field.TypeInt, Nullable: true},
		{Name: "subject_subject_id", Type: field.TypeInt, Nullable: true},
		{Name: "teacher_teachers_id", Type: field.TypeInt, Nullable: true},
	}
	// LessonplansTable holds the schema information for the "lessonplans" table.
	LessonplansTable = &schema.Table{
		Name:       "lessonplans",
		Columns:    LessonplansColumns,
		PrimaryKey: []*schema.Column{LessonplansColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "lessonplans_sections_sec_id",
				Columns: []*schema.Column{LessonplansColumns[2]},

				RefColumns: []*schema.Column{SectionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "lessonplans_subjects_Subject_ID",
				Columns: []*schema.Column{LessonplansColumns[3]},

				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "lessonplans_teachers_Teachers_ID",
				Columns: []*schema.Column{LessonplansColumns[4]},

				RefColumns: []*schema.Column{TeachersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SectionsColumns holds the columns for the "sections" table.
	SectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "group", Type: field.TypeInt},
		{Name: "recieve", Type: field.TypeInt},
		{Name: "seat_left", Type: field.TypeInt},
		{Name: "room", Type: field.TypeString},
		{Name: "date_time", Type: field.TypeString},
	}
	// SectionsTable holds the schema information for the "sections" table.
	SectionsTable = &schema.Table{
		Name:        "sections",
		Columns:     SectionsColumns,
		PrimaryKey:  []*schema.Column{SectionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SubjectsColumns holds the columns for the "subjects" table.
	SubjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "subject_name", Type: field.TypeString, Unique: true},
		{Name: "credit", Type: field.TypeInt},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "teacher_teacher_id", Type: field.TypeInt, Nullable: true},
	}
	// SubjectsTable holds the schema information for the "subjects" table.
	SubjectsTable = &schema.Table{
		Name:       "subjects",
		Columns:    SubjectsColumns,
		PrimaryKey: []*schema.Column{SubjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "subjects_teachers_Teacher_ID",
				Columns: []*schema.Column{SubjectsColumns[4]},

				RefColumns: []*schema.Column{TeachersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TeachersColumns holds the columns for the "teachers" table.
	TeachersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "teacher_name", Type: field.TypeString, Unique: true},
		{Name: "user_user_id", Type: field.TypeInt, Nullable: true},
	}
	// TeachersTable holds the schema information for the "teachers" table.
	TeachersTable = &schema.Table{
		Name:       "teachers",
		Columns:    TeachersColumns,
		PrimaryKey: []*schema.Column{TeachersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "teachers_users_User_ID",
				Columns: []*schema.Column{TeachersColumns[2]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString, Unique: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LessonplansTable,
		SectionsTable,
		SubjectsTable,
		TeachersTable,
		UsersTable,
	}
)

func init() {
	LessonplansTable.ForeignKeys[0].RefTable = SectionsTable
	LessonplansTable.ForeignKeys[1].RefTable = SubjectsTable
	LessonplansTable.ForeignKeys[2].RefTable = TeachersTable
	SubjectsTable.ForeignKeys[0].RefTable = TeachersTable
	TeachersTable.ForeignKeys[0].RefTable = UsersTable
}

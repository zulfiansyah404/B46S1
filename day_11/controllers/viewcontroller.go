package controllers

import (
	"project/models"
	"github.com/labstack/echo/v4"
	"strconv"
	"net/http"
	"text/template"
)

// Fungsi mengeluarkan respond Hello World
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Worldl!")
}

// Fungsi menampilkan index.html beserta ListProject
func Home(c echo.Context) error {
	var tmp, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}

	projects := map[string]interface{} {
		"Projects": projectList,
	}

	return tmp.Execute(c.Response(), projects)
}

// Fungsi menampilkan Menu untuk menambah project
func AddProjectView(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

// Fungsi untuk menampilkan contact.html
func Contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

// Fungsi untuk menampilkan testimonial.html
func TestimonialView(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

// Fungsi untuk menampilkan project-detail.html
func ProjectDetailView(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = models.Project{}

	for i, data := range projectList {
		if id == i {
			ProjectDetail = models.Project{
				Name:    		data.Name,
				StartDate:  	data.StartDate,
				EndDate:    	data.EndDate,
				Duration:   	data.Duration,
				Description: 	data.Description,
				NodeJs:     	data.NodeJs,
				ReactJs:    	data.ReactJs,
				Golang:     	data.Golang,
				Java: 			data.Java,
				Image: 			data.Image,
			}
		}
	}
	data := map[string]interface{}{
		"Project":   ProjectDetail,
		"Id":		 id,
	}

	var tmpl, err = template.ParseFiles("views/project-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

// Fungsi untuk menampilkan menu untuk mengedit suatu project berdasarkan ID
func EditProjectView(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = models.Project{}

	for i, data := range projectList {
		if id == i {
			ProjectDetail = models.Project{
				Name:    data.Name,
				StartDate:  	data.StartDate,
				EndDate:    	data.EndDate,
				Duration:   	data.Duration,
				Description: 	data.Description,
				NodeJs:     	data.NodeJs,
				ReactJs:    	data.ReactJs,
				Golang:     	data.Golang,
				Java: 			data.Java,
				Image: 			data.Image,
			}
		}
	}

	data := map[string]interface{}{
		"Project":   ProjectDetail,
		"Id": id,
	}

	var tmpl, err = template.ParseFiles("views/edit-project.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)

}
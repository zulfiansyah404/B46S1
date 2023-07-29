package controllers

import (
	"fmt"
	"net/http"
	"project/connection"
	"project/models"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

/* KAMUS */
// var projectList = []models.Project{}

// /* SUBPROGRAM PRIVATE */
// // Prosedur untuk push (stack) project baru ke dalam Project List
// func pushProjectList(item models.Project) {
// 	projectList = append(projectList, item)
// }

// Fungsi untuk mengembalikan selisih dari tanggal awal dan tanggal akhir
func getDuration(date_start string, date_end string) string {
	date1, _ := time.Parse("2006-01-02", date_start)
	date2, _ := time.Parse("2006-01-02", date_end)

	duration := date2.Sub(date1)
	days := int(duration.Hours() / 24)
	months := days/30
	years := months/12
	centuries := years/100

	if (centuries > 0) {
		if (centuries > 1) {
			return strconv.Itoa(centuries) + " centuries"
		}
		return strconv.Itoa(centuries) + " century"
	}

	if (years > 0) {
		if (years > 1) {
			return strconv.Itoa(years) + " years"
		}
		return strconv.Itoa(years) + " years"
	}

	if (months > 0) {
		if (months > 1) {
			return strconv.Itoa(months) + " months"
		}
		return strconv.Itoa(months) + " months"
	}

	if (days > 1) {
		return strconv.Itoa(days) + " days"
	}

	return strconv.Itoa(days) + " day"
}

/* SUBPROGRAM PUBLIC */
// Fungsi untuk menambah project ke dalam projectList[]
func AddProject(c echo.Context) error {
	name := c.FormValue("name")
	startDate := c.FormValue("start-date")
	endDate := c.FormValue("end-date")
	description := c.FormValue("description")
	nodeJs := c.FormValue("node-js")
	reactJs := c.FormValue("react-js")
	golang := c.FormValue("go")
	java := c.FormValue("java")
	image := c.FormValue("image")

	// Ubah format startdate dan enddate dari string menjadi time.Time
	startDateConv, _ := time.Parse("2006-01-02", startDate)
	endDateConv, _ := time.Parse("2006-01-02", endDate)

	var newProject = models.Project{
		Name: 		name,	
		StartDate: 	startDateConv,
		EndDate:   	endDateConv,
		Duration:   getDuration(startDate, endDate),
		Description: description,
		NodeJs:     (nodeJs == "checked"),
		ReactJs:    (reactJs == "checked"),
		Golang:     (golang == "checked"),
		Java: 		(java == "checked"),
		Image: 		image,
	}

	if connection.DB.Create(&newProject).RowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to add project"})
	}

	fmt.Println(" Project added successfully")
	fmt.Println("--------------------------")
	fmt.Println("Name : " + newProject.Name)
	fmt.Println("Start Date : " + newProject.StartDate.String() + " (" + startDate + ")")
	fmt.Println("End Date : " + newProject.EndDate.String() + " (" + endDate + ")")
	fmt.Println("Duration : " + newProject.Duration)
	fmt.Println("Description : " + newProject.Description)
	fmt.Println("Node JS : " + strconv.FormatBool(newProject.NodeJs))
	fmt.Println("React JS : " + strconv.FormatBool(newProject.ReactJs))
	fmt.Println("Golang : " + strconv.FormatBool(newProject.Golang))
	fmt.Println("Java : " + strconv.FormatBool(newProject.Java))
	fmt.Println("Image : " + newProject.Image)
	fmt.Println("--------------------------\n\n")

	return c.Redirect(http.StatusMovedPermanently, "/")
}

// Fungsi untuk mengedit elemen dari listProject berdasarkan ID
func EditProject(c echo.Context) error {
	id, _:= strconv.Atoi(c.Param("id"))

	name := c.FormValue("name")
	startDate := c.FormValue("start-date")
	endDate := c.FormValue("end-date")
	description := c.FormValue("description")
	nodeJs := c.FormValue("node-js")
	reactJs := c.FormValue("react-js")
	golang := c.FormValue("go")
	java := c.FormValue("java")
	image := c.FormValue("image")

	// Ubah format startdate dan enddate dari string menjadi time.Time
	startDateConv, _ := time.Parse("2006-01-02", startDate)
	endDateConv, _ := time.Parse("2006-01-02", endDate)

	var editedProject = map[string]interface{} {
		"name": 		name,
		"start_date": 	startDateConv,
		"end_date":   	endDateConv,
		"duration":   	getDuration(startDate, endDate),
		"description": 	description,
		"node_js":     	(nodeJs == "checked"),
		"react_js":    	(reactJs == "checked"),
		"golang":     	(golang == "checked"),
		"java": 		(java == "checked"),
		"image": 		image,
	}

	// fmt.Println("NodeJs: ", editedProject.NodeJs)
	// fmt.Println("ReactJs: ", editedProject.ReactJs)
	// fmt.Println("Golang: ", editedProject.Golang)
	// fmt.Println("Java: ", editedProject.Java)

	if connection.DB.Model(&models.Project{}).Where("id = ?", id).Updates(&editedProject).RowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to edit project"})
	}
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}

// Fungsi untuk mengedit elemen dari listProject berdasarkan ID
func DeleteProject(c echo.Context) error {
	Id, _ := strconv.Atoi(c.Param("id"))

	if connection.DB.Delete(&models.Project{}, Id).RowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete project"})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}
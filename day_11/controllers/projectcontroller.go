package controllers

import (
	"project/models"
	"github.com/labstack/echo/v4"
	"time"
	"net/http"
	"strconv"
)

/* KAMUS */
var projectList = []models.Project{
	{
		Name:			"String Matching dan Regular Expression dalam Pembuatan ChatGPT Sederhana",
		StartDate:		"2021-09-01",
		EndDate:		"2021-10-01",
		Duration:		"1 month",
		Description: 	"Project Tubes III STIMA Teknik Informatika ITB merupakan project yang dikerjakan oleh 3 orang mahasiswa Teknik Informatika ITB. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma.Repository ini berisi aplikasi ChatGPT sederhana berbasis web. Aplikasi ChatGPT sederhana ini dibangun dengan mengimplementasikan algoritma string matching (Knuth-Morris-Pratt dan Boyer-Moore) dan juga regular expression. Aplikasi ini dapat menjawab pertanyaan berdasarkan masukan dari pengguna. Jenis pertanyaan yang dapat dijawab antara lain: Pertanyaan yang terdapat didalam database (menjawab menggunakan algoritma KMP atau BM) Kalkulator sederhana Menanyakan hari berdasarkan tanggal Menambahkan pertanyaan dan jawaban ke dalam database Menghapus pertanyaan dari database Selain itu, history dari pertanyaan yang dimasukkan akan disimpan ke dalam sebuah section pertanyaan untuk sebuah sesi. Pengguna juga dapat menambahkan section baru dengan mengklik add new history. Untuk pembuatan program ini, pembuatan frontend dilakukan dengan menggunakan bahasa pemrograman HTML, CSS, dan juga Java dengan menggunakan Framework React. Untuk Backend dari program dibuat dengan menggunakan bahasa pemrograman Golang. Sedangkan untuk penyimpanan database menggunakan MySQL.",
		NodeJs:     	true,
		ReactJs:    	false,
		Golang:     	true,
		Java: 			false,
		Image: 			"assets/img/project.png",
	},

}

/* SUBPROGRAM PRIVATE */
// Prosedur untuk push (stack) project baru ke dalam Project List
func pushProjectList(item models.Project) {
	projectList = append(projectList, item)
}

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

	var newProject = models.Project{
		Name: 		name,
		StartDate:  startDate,
		EndDate:    endDate,
		Duration:   getDuration(startDate, endDate),
		Description: description,
		NodeJs:     (nodeJs == "node-js"),
		ReactJs:    (reactJs == "react-js"),
		Golang:     (golang == "go"),
		Java: 		(java == "java"),
		Image: 		image,
	}

	pushProjectList(newProject)

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

	var editedProject = models.Project{
		Name: name,
		StartDate:  	startDate,
		EndDate:    	endDate,
		Duration:   	getDuration(startDate, endDate),
		Description:	description,
		NodeJs:     	(nodeJs == "node-js"),
		ReactJs:    	(reactJs == "react-js"),
		Golang:     	(golang == "go"),
		Java: 			(java == "java"),
		Image: 			image,
	}

	projectList[id] = editedProject
	return c.Redirect(http.StatusMovedPermanently, "/")
}

// Fungsi untuk mengedit elemen dari listProject berdasarkan ID
func DeleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	projectList = append(projectList[:id], projectList[id+1:]...)
	return c.Redirect(http.StatusMovedPermanently, "/")
}
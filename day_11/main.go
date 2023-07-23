package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	// Mengatur penanganan file static
	e.Static("/assets", "assets")

	// Daftar Routes GET
	e.GET("/hello", helloWorld)
	e.GET("/", home)
	e.GET("/home", home)
	e.GET("/add-project", addProject)
	e.GET("/contact", contact)
	e.GET("/testimonial", testimonial)
	e.GET("/project-detail/:id", projectDetail)

	//Daftar Routes POST
	e.POST("/", addFormProject)

	// Server
	e.Logger.Fatal(e.Start("localhost:8000"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Worldl!")
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func addProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func testimonial(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"id":      	id,
		"Title":   	"String Matching dan Regular Expression dalam Pembuatan ChatGPT Sederhana",
		"Duration": "1 month",
		"Content": 	"Project Tubes III STIMA Teknik Informatika ITB merupakan project yang dikerjakan oleh 3 orang mahasiswa Teknik Informatika ITB. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma. Project ini merupakan project yang dikerjakan untuk memenuhi tugas mata kuliah Strategi Algoritma.Repository ini berisi aplikasi ChatGPT sederhana berbasis web. Aplikasi ChatGPT sederhana ini dibangun dengan mengimplementasikan algoritma string matching (Knuth-Morris-Pratt dan Boyer-Moore) dan juga regular expression. Aplikasi ini dapat menjawab pertanyaan berdasarkan masukan dari pengguna. Jenis pertanyaan yang dapat dijawab antara lain: Pertanyaan yang terdapat didalam database (menjawab menggunakan algoritma KMP atau BM) Kalkulator sederhana Menanyakan hari berdasarkan tanggal Menambahkan pertanyaan dan jawaban ke dalam database Menghapus pertanyaan dari database Selain itu, history dari pertanyaan yang dimasukkan akan disimpan ke dalam sebuah section pertanyaan untuk sebuah sesi. Pengguna juga dapat menambahkan section baru dengan mengklik add new history. Untuk pembuatan program ini, pembuatan frontend dilakukan dengan menggunakan bahasa pemrograman HTML, CSS, dan juga Javascript dengan menggunakan Framework React. Untuk Backend dari program dibuat dengan menggunakan bahasa pemrograman Golang. Sedangkan untuk penyimpanan database menggunakan MySQL.",
		"Image": "https://user-images.githubusercontent.com/91790457/236441708-c1997cae-a9da-48e6-9d13-e9e8af6d38b8.png",
		"ListIconTech" : []string{
			"<i class='fa-brands fa-node-js'></i>",
			"<i class='fa-brands fa-golang'></i>",
		},
		"ListTech" : []string{
			"Node JS",
			"Golang",
		},
	}

	var tmpl, err = template.ParseFiles("views/project-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func addFormProject(c echo.Context) error {

	nameProject := c.FormValue("name")
	startDate := c.FormValue("start-date")
	endDate := c.FormValue("end-date")
	description := c.FormValue("description")
	nodeJs := c.FormValue("node-js")
	reactJs := c.FormValue("react-js")
	golang := c.FormValue("go")
	java := c.FormValue("java")
	image := c.FormValue("image")

	println("Name Project : " + nameProject)
	println("Start Date : " + startDate)
	println("End Date : " + endDate)
	println("Description : " + description)
	println("Technologies : " + nodeJs)
	println("Technologies : " + reactJs)
	println("Technologies : " + golang)
	println("Technologies : " + java)
	println("Image : " + image)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

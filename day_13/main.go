package main

import (	
	"github.com/labstack/echo/v4"	
	"project/controllers"
	"project/connection"
)

func main() {
	e := echo.New()
	connection.DatabaseConnect()
	
	// Mengatur penanganan file static
	e.Static("/assets", "assets")
	//File static dari directory pc
	e.Static("/path", "C:")

	// Daftar Routes GET
	// Diambil dari controllers/viewcontroller.go
	e.GET("/hello", controllers.HelloWorld)
	e.GET("/", controllers.Home)
	e.GET("/add-project", controllers.AddProjectView)
	e.GET("/contact", controllers.Contact)
	e.GET("/testimonials", controllers.TestimonialView)
	e.GET("/project-detail/:id", controllers.ProjectDetailView)
	e.GET("/edit-project/:id", controllers.EditProjectView)

	//Daftar Routes POST
	// Diambil dari controllers/projectcontroller.go
	e.POST("/", controllers.AddProject)
	e.POST("/edit-project/:id", controllers.EditProject)
	e.POST("/delete-project/:id", controllers.DeleteProject)

	// Server
	e.Logger.Fatal(e.Start("localhost:8000"))
}


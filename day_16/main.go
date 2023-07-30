package main

import (
	"project/connection"
	"project/controllers"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	connection.DatabaseConnect()
	
	// Tambahkan middleware
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

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
	e.GET("/login", controllers.LoginView)
	e.GET("/register", controllers.RegisterView)

	//Daftar Routes POST
	// Diambil dari controllers/projectcontroller.go
	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.Register)
	e.POST("/logout", controllers.Logout)
	e.POST("/", controllers.AddProject)
	e.POST("/edit-project/:id", controllers.EditProject)
	e.POST("/delete-project/:id", controllers.DeleteProject)

	// Server
	e.Logger.Fatal(e.Start("localhost:8000"))
}


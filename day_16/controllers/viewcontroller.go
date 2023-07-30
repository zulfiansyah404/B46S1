package controllers

import (
	"net/http"
	"project/connection"
	"project/models"
	"strconv"
	"text/template"
	"fmt"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var ProjectList []models.Project


// func updateProjectList() {
// 	// connection.DB.Find(&ProjectList)
// }

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

	// updateProjectList()
	if err := connection.DB.Find(&ProjectList).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : errSess.Error()})
	}

	if sess.Values["isLogin"] == nil {
		sess.Values["isLogin"] = false
	}

	fmt.Println("Home")
	fmt.Println("------------------")
	fmt.Println("IsLogin: ", sess.Values["isLogin"])
	fmt.Println("Username: ", sess.Values["username"])
	fmt.Println("Message: ", sess.Values["message"])
	fmt.Println("Status: ", sess.Values["status"])
	fmt.Println("Name: ", sess.Values["name"])
	fmt.Println("Email: ", sess.Values["email"])
	fmt.Println("------------------")
	fmt.Println("")
	fmt.Println("")

	dataResponds := map[string]interface{} {
		"Projects": ProjectList,
		"Message": sess.Values["message"],
		"Status" : sess.Values["status"],
		"User": sess.Values["username"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())
	// // Print isi projects
	// fmt.Println(projects)

	return tmp.Execute(c.Response(), dataResponds)
}

func LoginView(c echo.Context) error {
	var tmp, err = template.ParseFiles("views/login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : errSess.Error()})
	}

	dataResponds := map[string]interface{} {
		"Message": sess.Values["message"],
		"Status" : sess.Values["status"],
		"User": sess.Values["username"],
	}

	return tmp.Execute(c.Response(), dataResponds)
}

func RegisterView(c echo.Context) error {
	var tmp, err = template.ParseFiles("views/register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : errSess.Error()})
	}

	dataResponds := map[string]interface{} {
		"Message": sess.Values["message"],
		"Status" : sess.Values["status"],
		"User": sess.Values["username"],
	}

	return tmp.Execute(c.Response(), dataResponds)
}

// Fungsi menampilkan Menu untuk menambah project
func AddProjectView(c echo.Context) error {
	fmt.Println("Add Project View")
	var tmp, err = template.ParseFiles("views/project.html")

	if err != nil {
		fmt.Println("Break 1")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	sess, errSess := session.Get("session", c)
	if errSess != nil {
		fmt.Println("Break 2")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : errSess.Error()})
	}

	dataResponds := map[string]interface{} {
		"Message": sess.Values["message"],
		"Status" : sess.Values["status"],
		"User": sess.Values["username"],
		"TechError" : sess.Values["techError"],
		"DateError" : sess.Values["dateError"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	delete(sess.Values, "techError")
	delete(sess.Values, "dateError")

	sess.Save(c.Request(), c.Response())
	
	return tmp.Execute(c.Response(), dataResponds)
}

// Fungsi untuk menampilkan contact.html
func Contact(c echo.Context) error {
	var tmp, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : errSess.Error()})
	}

	dataResponds := map[string]interface{} {
		"Message": sess.Values["message"],
		"Status" : sess.Values["status"],
		"User": sess.Values["username"],
	}

	return tmp.Execute(c.Response(), dataResponds)
}

// Fungsi untuk menampilkan testimonial.html
func TestimonialView(c echo.Context) error {
	var tmp, err = template.ParseFiles("views/testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : errSess.Error()})
	}

	dataResponds := map[string]interface{} {
		"Message": sess.Values["message"],
		"Status" : sess.Values["status"],
		"User": sess.Values["username"],
	}

	return tmp.Execute(c.Response(), dataResponds)
}

// Fungsi untuk menampilkan project-detail.html
func ProjectDetailView(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = models.Project{}

	for _, data := range ProjectList {
		if id == data.ID {
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

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : errSess.Error()})
	}


	data := map[string]interface{}{
		"Project":   ProjectDetail,
		"Id":		 id,
		"StartDate": ProjectDetail.StartDate.Format("2 January 2006"),
		"EndDate":   ProjectDetail.EndDate.Format("2 January 2006"),
		"Message": sess.Values["message"],
		"Status" : sess.Values["status"],
		"User": sess.Values["username"],
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

	for _, data := range ProjectList {
		if id == data.ID {
			ProjectDetail = models.Project{
				Name:    	data.Name,
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

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : errSess.Error()})
	}

	data := map[string]interface{}{
		"Project":   ProjectDetail,
		"Id": id,
		"StartDate": ProjectDetail.StartDate.Format("2006-01-02"),
		"EndDate":   ProjectDetail.EndDate.Format("2006-01-02"),
		"Message": sess.Values["message"],
		"Status" : sess.Values["status"],
		"User": sess.Values["username"],
		"TechError" : sess.Values["techError"],
		"DateError" : sess.Values["dateError"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	delete(sess.Values, "techError")
	delete(sess.Values, "dateError")

	var tmpl, err = template.ParseFiles("views/edit-project.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)

}
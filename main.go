package main

import (
	"echo-framework/config"
	"echo-framework/controller"
	_ "echo-framework/docs"
	"echo-framework/middlewares"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	// load env first
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// initialize connect to db
	// config.Connect()
	config.StartDB()
}

// @title Echo-Framework API
// @version 1.0
// @description this is a sample service rest echo-framework
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email irfan.email@email.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9000
// @BasePath /
func main() {
	r := echo.New()

	r.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:" + config.CSRFTokenHeader,
		ContextKey:  config.CSRFKey,
	}))

	r.POST("/sayhello", controller.SayHello)
	r.GET("/index", controller.Index)

	r.GET("/", controller.HelloWorld)

	r.GET("/json", controller.JsonMap)

	r.GET("/page1", controller.Page1)

	r.GET("/user", controller.GetUser)
	r.POST("/user", controller.CreateUser)

	// employee login
	r.POST("/login", controller.EmployeeLogin)
	employeeRouter := r.Group("/employee")
	{
		employeeRouter.Use(middlewares.Authentication)
		employeeRouter.POST("", controller.CreateEmployee)
		employeeRouter.PUT("", controller.UpdateEmployee)
		employeeRouter.DELETE("", controller.DeleteEmployee)
	}

	// Route for swagger
	r.GET("/swagger/*", echoSwagger.WrapHandler)

	var PORT = os.Getenv("PORT")
	r.Logger.Fatal(r.Start(":" + PORT))

}

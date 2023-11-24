package controller

import (
	"echo-framework/config"
	"echo-framework/helpers"
	"echo-framework/models"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func HelloWorld(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello World!")
}

func JsonMap(ctx echo.Context) error {
	data := models.M{"message": "hello", "counter": "2", "statusCode": http.StatusOK}
	return ctx.JSON(http.StatusOK, data)
}

func Page1(ctx echo.Context) error {
	name := ctx.QueryParam("name")
	data := "Hello " + name
	result := fmt.Sprintf("%s", data)
	fmt.Println(result)
	return ctx.String(http.StatusOK, result)
}

func GetUser(ctx echo.Context) error {
	id := ctx.QueryParam("id")

	getUserById(id)
	user := models.User{}
	if err := ctx.Bind(&user); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

// @Summary Create Employee
// @Descriptio Create a new Employee with the input payload
// @Tags employee
// @Accept json
// @Produce json
// @Param Employee body models.Employee true "Create Employee"
// @Success 200 {object} models.Employee
// @Router /employee [post]
func CreateEmployee(ctx echo.Context) error {
	db := config.GetDB()

	employee := models.Employee{}

	if err := ctx.Bind(&employee); err != nil {
		return err
	}

	db.Debug().Create(&employee)

	fmt.Println("create employee")
	return ctx.JSON(http.StatusOK, employee)
}

func UpdateEmployee(ctx echo.Context) error {
	db := config.GetDB()

	employee := models.Employee{}

	if err := ctx.Bind(&employee); err != nil {
		return err
	}

	// db.Save(&employee)
	err := db.Debug().Model(&employee).Where("id = ?", employee.ID).Updates(models.Employee{
		Full_name: employee.Full_name,
		Password:  employee.Password,
		Email:     employee.Email,
		Division:  employee.Division,
	}).Error

	if err != nil {
		return err
	}
	fmt.Println("update employee")

	return ctx.JSON(http.StatusOK, employee)
}

func CreateUser(ctx echo.Context) error {
	db := config.GetDB()

	user := models.User{}

	if err := ctx.Bind(&user); err != nil {
		return err
	}

	db.Create(&user)

	fmt.Println("create user")
	return ctx.JSON(http.StatusOK, user)
}

func getUserById(id string) {
	db := config.GetDB()

	User := &models.User{}

	err := db.First(&User, "id = ?", id).Error

	if err != nil {
		fmt.Println("error getting user by id :", err)
	} else {
		fmt.Printf("user data : %v ", User)
	}
}

func DeleteEmployee(ctx echo.Context) error {
	db := config.GetDB()

	employee := models.Employee{}

	delResponse := models.Response{
		ResponseCode: "00",
		ResponseDesc: "Delete Success",
	}

	paramId := ctx.QueryParam("id")

	if err := ctx.Bind(&employee); err != nil {
		return err
	}

	db.Model(&employee).Where("id = ?", paramId).Delete(&employee)

	fmt.Println("Delete Employee")

	return ctx.JSON(http.StatusOK, delResponse)
}

func EmployeeLogin(ctx echo.Context) error {
	data := make(map[string]interface{})
	if err := ctx.Bind(&data); err != nil {
		return err
	}

	db := config.GetDB()
	employee := models.Employee{}
	if err := ctx.Bind(&employee); err != nil {
		return err
	}

	password := employee.Password

	err := db.Debug().Where("email = ?", employee.Email).Take(&employee).Error

	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			ResponseCode: strconv.Itoa(http.StatusUnauthorized),
			ResponseDesc: "invalid email/password",
		})
	}

	comparePass := helpers.ComparePass([]byte(employee.Password), []byte(password))

	if !comparePass {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			ResponseCode: strconv.Itoa(http.StatusUnauthorized),
			ResponseDesc: "invalid email/password",
		})
	}

	token := helpers.GenerateToken(uint(employee.ID), employee.Email)
	// return ctx.JSON(http.StatusOK, map[string]interface{}{
	// 	"ResponseCode": strconv.Itoa(http.StatusOK),
	// 	"ResponseDesc": "OK",
	// 	"Token":        token,
	// })
	return ctx.JSON(http.StatusOK, models.Response{
		ResponseCode: strconv.Itoa(http.StatusOK),
		ResponseDesc: "OK",
		Data: map[string]interface{}{
			"token": token,
		},
	})
}

func Index(c echo.Context) error {
	tmpl := template.Must(template.ParseGlob("./*.html"))
	data := make(map[string]interface{})
	data[config.CSRFKey] = c.Get(config.CSRFKey)
	return tmpl.Execute(c.Response(), data)
}

func SayHello(c echo.Context) error {
	data := make(map[string]interface{})
	if err := c.Bind(&data); err != nil {
		return err
	}

	message := fmt.Sprintf("hello %s", data["name"])
	return c.JSON(http.StatusOK, message)
}

package middlewares

import (
	"echo-framework/config"
	"echo-framework/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func ProductAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		db := config.GetDB()
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				ResponseCode: strconv.Itoa(http.StatusBadRequest),
				ResponseDesc: "invalid parameter",
			})
		}

		userData := c.Get("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))

		product := models.Product{}

		err = db.Select("user_id").First(&product, uint(productId)).Error

		if err != nil {
			return c.JSON(http.StatusNotFound, models.Response{
				ResponseCode: strconv.Itoa(http.StatusNotFound),
				ResponseDesc: "Data doesn't exist",
			})
		}

		if product.UserID != userId {
			return c.JSON(http.StatusUnauthorized, models.Response{
				ResponseCode: strconv.Itoa(http.StatusUnauthorized),
				ResponseDesc: "You are not allowed to access this data",
			})
		}

		return next(c)
	}
}

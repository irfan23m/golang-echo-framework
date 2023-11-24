package middlewares

import (
	"echo-framework/helpers"
	"echo-framework/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("Authenticating")
	return func(c echo.Context) error {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		fmt.Println("verif token", verifyToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, models.Response{
				ResponseCode: strconv.Itoa(http.StatusUnauthorized),
				ResponseDesc: err.Error(),
			})
		}

		c.Set("userCacheData", verifyToken)
		return next(c)
	}
}

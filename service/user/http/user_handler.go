package http

import (
	"net/http"
	"time"

	"github.com/go-pg/pg"
	"github.com/labstack/echo/v4"
	"gitlab.com/km/go-kafka-playground/middleware"
	"gitlab.com/km/go-kafka-playground/models"
	"gitlab.com/km/go-kafka-playground/service/user"
)

type userHandler struct {
	userUs user.UserUsecaseInf
}

func NewUserHandler(e *echo.Echo, middL *middleware.GoMiddleware, us user.UserUsecaseInf) {
	handler := &userHandler{
		userUs: us,
	}
	e.GET("/users", handler.Create)
	e.POST("/users", handler.Create)
}

func responseMessage(str string) map[string]interface{} {
	return map[string]interface{}{
		"error": str,
	}
}

func (u *userHandler) Create(c echo.Context) error {
	now := time.Now()
	nullTime := pg.NullTime{now}
	user := &models.User{
		Email:     "test@gmail.com",
		FirstName: "testFname",
		LastName:  "testLname",
		Age:       15,
		CreatedAt: nullTime,
		UpdatedAt: nullTime,
	}

	if err := u.userUs.Create(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responseMessage(err.Error()))
	}
	responseData := map[string]interface{}{
		"user": user,
	}
	return c.JSON(http.StatusOK, responseData)
}

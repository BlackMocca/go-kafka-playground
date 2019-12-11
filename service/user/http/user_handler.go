package http

import (
	"net/http"
	"strconv"
	"time"

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
	/* จาก GET ต้องเป็น POST */
	e.GET("/users", handler.Create)
	e.GET("/users/:id", handler.FetchOne)
}

func responseMessage(str string) map[string]interface{} {
	return map[string]interface{}{
		"error": str,
	}
}

func (u *userHandler) FetchOne(c echo.Context) error {
	var user *models.User
	var err error
	var id, _ = strconv.Atoi(c.Param("id"))

	user, err = u.userUs.FetchOne(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, responseMessage(err.Error()))
	}

	responseData := map[string]interface{}{
		"user": user,
	}
	return c.JSON(http.StatusOK, responseData)
}

func (u *userHandler) Create(c echo.Context) error {
	now := time.Now()
	user := &models.User{
		Email:     "test@gmail.com",
		FirstName: "testFname",
		LastName:  "testLname",
		Age:       15,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	/* create user in postgres */

	if err := u.userUs.Create(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responseMessage(err.Error()))
	}

	/* send message to create user in mongodb */
	partition, offset, err := u.userUs.InvokeCreateEvent(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responseMessage(err.Error()))
	}

	responseData := map[string]interface{}{
		"user":      user,
		"partition": partition,
		"offset":    offset,
	}
	return c.JSON(http.StatusOK, responseData)
}

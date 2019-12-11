package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gitlab.com/km/go-kafka-playground/middleware"
	"gitlab.com/km/go-kafka-playground/service/user"
)

type kafkaHandler struct {
	userUs user.UserUsecaseInf
}

func NewKafkaHandler(e *echo.Echo, middL *middleware.GoMiddleware, us user.UserUsecaseInf) {
	handler := &kafkaHandler{
		userUs: us,
	}
	/** จาก GET ต้องเป็น post */
	e.GET("/kafka/users/:id", handler.CreateUserInMongo)
}

func responseMessage(str string) map[string]interface{} {
	return map[string]interface{}{
		"error": str,
	}
}

func (k *kafkaHandler) CreateUserInMongo(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))

	user, err := k.userUs.FetchOne(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, responseMessage(err.Error()))
	}

	if err = k.userUs.CreateIntoMongoDB(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responseMessage(err.Error()))
	}

	responseData := map[string]interface{}{
		"user": user,
	}
	return c.JSON(http.StatusOK, responseData)
}

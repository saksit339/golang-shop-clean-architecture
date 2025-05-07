package custom

import "github.com/labstack/echo/v4"

type errorMessage struct {
	Message string
}

func Error(ctx echo.Context, statusCode int, message string) error {
	return ctx.JSON(statusCode,
		&errorMessage{
			Message: message,
		})
}

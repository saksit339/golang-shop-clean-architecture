package custom

import "github.com/labstack/echo/v4"

type ErrorMessage struct {
	Message string
}

func CustomError(ctx echo.Context, statusCode int, message string) error {
	return ctx.JSON(statusCode,
		&ErrorMessage{
			Message: message,
		})
}

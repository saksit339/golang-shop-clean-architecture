package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type EchoRequest interface {
	Bind(obj any) error
}

type customEchoRequest struct {
	ctx       echo.Context
	validator *validator.Validate
}

func NewCustomEchoRequest(ctx echo.Context) EchoRequest {
	return &customEchoRequest{
		ctx:       ctx,
		validator: validator.New(),
	}
}

func (c *customEchoRequest) Bind(obj any) error {
	if err := c.ctx.Bind(obj); err != nil {
		return err
	}
	if err := c.ctx.Validate(obj); err != nil {
		return err
	}
	return nil
}

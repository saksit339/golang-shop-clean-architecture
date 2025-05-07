package custom

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var once sync.Once
var validate *validator.Validate

type EchoRequest interface {
	Bind(obj any) error
}

type customEchoRequest struct {
	ctx       echo.Context
	validator *validator.Validate
}

func NewCustomEchoRequest(ctx echo.Context) EchoRequest {
	once.Do(func() {
		validate = validator.New()
	})
	return &customEchoRequest{
		ctx:       ctx,
		validator: validate,
	}
}

func (c *customEchoRequest) Bind(obj any) error {
	if err := c.ctx.Bind(obj); err != nil {
		return err
	}
	if err := c.validator.Struct(obj); err != nil {
		return err
	}
	return nil
}

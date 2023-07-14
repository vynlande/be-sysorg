package util

import (
	"database/sql"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

type ErrorInterface interface {
	Panic(interface{})
	Handler(*fiber.Ctx, error) error
}

type errorStruct struct{}

func NewError() ErrorInterface {
	return &errorStruct{}
}

func (es *errorStruct) Panic(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func (es *errorStruct) Handler(c *fiber.Ctx, err error) error {
	var e *fiber.Error
	code := fiber.StatusInternalServerError
	if errors.As(err, &e) {
		code = e.Code
		c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
		return c.Status(code).SendString(err.Error())
	}
	// logger
	log.Println(err)

	// validation
	_, ok := err.(ValidationError)
	if ok {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Data:    err.Error(),
			Message: MESSAGE_BAD_REQUEST,
		})
	}
	// no data
	none := err == sql.ErrNoRows
	if none {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Data:    err.Error(),
			Message: MESSAGE_NOT_FOUND,
		})
	}

	return c.Status(fiber.StatusBadGateway).JSON(Response{
		Data:    err.Error(),
		Message: MESSAGE_BAD_SYSTEM,
	})
}

type ValidationError struct {
	Message string
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}

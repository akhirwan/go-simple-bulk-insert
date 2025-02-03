package counter

import (
	"context"
	"go-simple-bulk-insert/domain/counter/feature"
	"go-simple-bulk-insert/domain/counter/model"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CounterHandler interface {
	CreateHandler(c *fiber.Ctx) error
}

type counterHandler struct {
	feature feature.CounterFeature
}

func NewCounterHandler(feature feature.CounterFeature) CounterHandler {
	return &counterHandler{
		feature: feature,
	}
}

func (h counterHandler) CreateHandler(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, "fiberCtx", c)

	/* payload vaildation*/
	req := new(model.CreateRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(model.ErrCreateResponse{
			Code:  c.Response().StatusCode(),
			Error: "invalid request payload",
		})
	}

	if req.Type == "" || req.Total <= 0 || req.Action == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(model.ErrCreateResponse{
			Code:  c.Response().StatusCode(),
			Error: "type, total & action are required and total must be a number",
		})
	}
	/* eof payload vaildation*/

	code, message := h.feature.CreateFeature(ctx, req) // logic
	if strings.Contains(strconv.Itoa(code), "20") {    // to handle not OK response
		return c.Status(code).JSON(model.ErrCreateResponse{
			Code:  c.Response().StatusCode(),
			Error: message,
		})
	}

	return c.Status(code).JSON(model.CreateResponse{
		Code:    c.Response().StatusCode(),
		Message: message,
	})
}

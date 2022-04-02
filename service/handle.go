package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handle struct {
	service Servicer
}

func NewHandle(service Servicer) *Handle{
	return &Handle{service : service}
}

func(h *Handle) CallLogin(c echo.Context) error{
	var req Request
	data := h.service.Login(req)
	return c.JSON(http.StatusOK, data)
}
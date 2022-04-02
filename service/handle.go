package service

import (
	"fmt"
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
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid req")
	}

	fmt.Println("my is := ", req)
	data := h.service.Login(req)
	fmt.Println("data in callLogin is := ", data)
	return c.JSON(http.StatusOK, data)
}
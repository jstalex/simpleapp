package handler

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"simpleapp/internal/domain"
	"simpleapp/internal/service"
)

func NewHandler(s service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

type Handler struct {
	service service.Service
}

func (h *Handler) GetUser(reqCtx echo.Context) error {
	email := reqCtx.QueryParam("email")
	if email == "" {
		return reqCtx.JSON(http.StatusBadRequest, ResponseErr{Message: "empty email"})
	}

	user, err := h.service.GetUserByEmail(reqCtx.Request().Context(), email)
	if err != nil {
		log.Errorf("can't find user: %v", err.Error())
		if errors.Is(err, pgx.ErrNoRows) {
			return reqCtx.JSON(http.StatusNotFound, ResponseErr{Message: "user not found"})
		}
		return reqCtx.JSON(http.StatusInternalServerError, ResponseErr{Message: "internal error"})
	}

	return reqCtx.JSON(http.StatusOK, User{user})
}

func (h *Handler) PostUser(reqCtx echo.Context) error {
	user := domain.User{}

	if err := reqCtx.Bind(&user); err != nil {
		return reqCtx.JSON(http.StatusBadRequest, ResponseErr{Message: "invalid request body"})
	}

	if err := h.service.AddUser(reqCtx.Request().Context(), user); err != nil {
		return reqCtx.JSON(http.StatusInternalServerError, ResponseErr{Message: err.Error()})
	}

	return reqCtx.String(http.StatusOK, "user successfully added")
}

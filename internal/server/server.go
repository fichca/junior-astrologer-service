package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

const (
	groupAPODURL     = "/api/apod"
	allAPODURL       = "/"
	allAPODByDateURL = "/:date"
)

type service interface {
}

type handler struct {
	l *logrus.Logger
	r fiber.Router
	s service
}

func NewHandler(
	logger *logrus.Logger,
	router fiber.Router,
	service service) *handler {

	return &handler{
		l: logger,
		r: router,
		s: service,
	}
}

func (h handler) RegisterRoutes() {
	APODGroup := h.r.Group(groupAPODURL)

	APODGroup.Get(allAPODURL, h.All)
	APODGroup.Get(allAPODByDateURL, h.AllByDate)
}

func (h handler) All(ctx *fiber.Ctx) error {
	return nil
}

func (h handler) AllByDate(ctx *fiber.Ctx) error {
	return nil
}

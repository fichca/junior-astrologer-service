package server

import (
	"context"
	"github.com/fichca/junior-astrologer-service/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	groupAPODURL     = "/api/apod"
	allAPODURL       = "/"
	allAPODByDateURL = "/:date"
)

type service interface {
	Get(ctx context.Context) ([]*model.APODResponse, error)
	GetByDate(ctx context.Context, date *time.Time) (*model.APODResponse, error)
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

func (h *handler) RegisterRoutes() {
	APODGroup := h.r.Group(groupAPODURL)

	APODGroup.Get(allAPODURL, h.All)
	APODGroup.Get(allAPODByDateURL, h.AllByDate)
}

// @Summary Get all APOD entries
// @Description Retrieves a list of all APOD entries
// @Tags APOD
// @Produce json
// @Success 200 {array} model.APODResponse
// @Failure 500 {string} string "Internal server error"
// @Router /apod/ [get]
func (h *handler) All(ctx *fiber.Ctx) error {
	ar, err := h.s.Get(ctx.Context())
	if err != nil {
		h.l.Errorf("failed to get all: %v", err)
		return ctx.Status(http.StatusInternalServerError).SendString("Ooops something wrong!")
	}

	return ctx.Status(fiber.StatusOK).JSON(ar)
}

// @Summary Get APOD entry by date
// @Description Retrieves a single APOD entry for the specified date
// @Tags APOD
// @Produce json
// @Param date path string true "Date in format yyyy-mm-dd"
// @Success 200 {object} model.APODResponse
// @Failure 400 {string} string "Bad request, invalid date format"
// @Failure 500 {string} string "Internal server error"
// @Router /apod/{date} [get]
func (h *handler) AllByDate(ctx *fiber.Ctx) error {
	dateStr := ctx.Params("date")

	date, err := time.Parse(model.DateFormat, dateStr)
	if err != nil {
		h.l.Error("Invalid date format: ", err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	//make processing and sending an error when no user is present
	apodResponse, err := h.s.GetByDate(ctx.Context(), &date)
	if err != nil {
		h.l.Error("failed to get APOD by date: ", err)
		return ctx.Status(http.StatusInternalServerError).SendString("Ooops something wrong!")
	}

	return ctx.Status(fiber.StatusOK).JSON(apodResponse)
}

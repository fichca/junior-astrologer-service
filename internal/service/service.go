package service

import (
	"github.com/fichca/junior-astrologer-service/internal/model"
	"github.com/sirupsen/logrus"
)

type apodClient interface {
	GetAPOD() (*model.APODResponse, error)
}

type service struct {
	l *logrus.Logger
	c apodClient
}

func NewService(
	logger *logrus.Logger,
	apodClient apodClient) *service {
	return &service{
		l: logger,
		c: apodClient,
	}
}

func (s *service) ProcessAPOD() error {
	panic("implement me")
}

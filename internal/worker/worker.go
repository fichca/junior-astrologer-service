package worker

import (
	"context"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type apodService interface {
	ProcessAPOD(ctx context.Context) error
}

type worker struct {
	l  *logrus.Logger
	as apodService
	cr *cron.Cron
}

func NewWorker(apodService apodService, logger *logrus.Logger) *worker {
	return &worker{
		l:  logger,
		as: apodService,
		cr: cron.New(),
	}
}

func (w *worker) Start() error {
	_, err := w.cr.AddFunc("@daily", w.executeAPODJob)

	if err != nil {
		return err
	}
	w.cr.Start()
	w.l.Info("Worker started. Waiting for the daily job.")
	return nil
}

func (w *worker) Stop() {
	w.cr.Stop()
	w.l.Info("Worker stopped.")
}

func (w *worker) executeAPODJob() {
	w.l.Info("Starting APOD daily job...")
	err := w.as.ProcessAPOD(context.Background())
	if err != nil {
		w.l.Errorf("failed to process APOD: %v", err)
	}
	w.l.Info("APOD data successfully executed.")
}

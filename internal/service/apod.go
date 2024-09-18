package service

import (
	"context"
	"fmt"
	"github.com/fichca/junior-astrologer-service/internal/model"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)

//go:generate mockgen -source=apod.go -destination=mocks/mock.go
type apodClient interface {
	GetAPOD() (*model.APODClientResponse, error)
	DownloadImage(imageURL string) (io.Reader, error)
}

type imageRepo interface {
	PutObject(image string, data io.Reader) error
	GetImageUrl(image string) (string, error)
	GetImageUrls(images []string) ([]string, error)
}

type repo interface {
	Save(ctx context.Context, apod *model.APOD) error
	GetAll(ctx context.Context) ([]*model.APOD, error)
	GetByDate(ctx context.Context, date *time.Time) (*model.APOD, error)
}

type apodService struct {
	l  *logrus.Logger
	ac apodClient
	is imageRepo
	r  repo
}

func NewAPODService(logger *logrus.Logger, apodClient apodClient, imageRepo imageRepo, apodRepo repo) *apodService {
	return &apodService{
		l:  logger,
		ac: apodClient,
		is: imageRepo,
		r:  apodRepo,
	}
}

func (s *apodService) Get(ctx context.Context) ([]*model.APODResponse, error) {
	apods, err := s.r.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get APODs from repository: %w", err)
	}
	apodResponses := make([]*model.APODResponse, 0)
	if len(apods) == 0 {
		return apodResponses, nil
	}

	var images []string
	for _, apod := range apods {
		if apod.Id != "" {
			images = append(images, apod.Id)
		}
	}
	imageUrls, err := s.is.GetImageUrls(images)
	if err != nil {
		return nil, fmt.Errorf("failed to get image URLs: %w", err)
	}

	// Implement id to url ratio
	for i, apod := range apods {
		apodResponse := &model.APODResponse{
			Title:       apod.Title,
			Explanation: apod.Explanation,
			Copyright:   apod.Copyright,
			Date:        apod.Date,
			Url:         imageUrls[i],
		}
		apodResponses = append(apodResponses, apodResponse)
	}

	return apodResponses, nil
}

func (s *apodService) GetByDate(ctx context.Context, date *time.Time) (*model.APODResponse, error) {
	//check exist
	apod, err := s.r.GetByDate(ctx, date)
	if err != nil {
		return nil, fmt.Errorf("failed to get APOD for the date: %w", err)
	}

	imageUrl, err := s.is.GetImageUrl(apod.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get image URL for APOD: %w", err)
	}

	apodResponse := &model.APODResponse{
		Title:       apod.Title,
		Explanation: apod.Explanation,
		Copyright:   apod.Copyright,
		Date:        apod.Date,
		Url:         imageUrl,
	}

	return apodResponse, nil
}

func (s *apodService) Save(ctx context.Context, title, explanation, copyright string, date *time.Time, img io.Reader) error {
	apod := model.NewAPOD(uuid.NewString(), title, explanation, copyright, date)

	exist := s.exist(apod)
	if exist {
		return fmt.Errorf("apod already exist")
	}
	err := s.r.Save(ctx, apod)
	if err != nil {
		return err
	}
	return s.is.PutObject(apod.Id, img)
}

func (s *apodService) ProcessAPOD(ctx context.Context) error {
	apod, err := s.ac.GetAPOD()
	if err != nil {
		return err
	}
	img, err := s.ac.DownloadImage(apod.Url)
	if err != nil {
		return err
	}
	date := time.Time(apod.Date)
	return s.Save(ctx, apod.Title, apod.Explanation, apod.Copyright, &date, img)
}

func (s *apodService) exist(apod *model.APOD) bool {
	//impl
	return false
}

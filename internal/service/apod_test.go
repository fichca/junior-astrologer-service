package service

import (
	"bytes"
	"context"
	"errors"
	"github.com/fichca/junior-astrologer-service/internal/model"
	mocks "github.com/fichca/junior-astrologer-service/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAPODService_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepo(ctrl)
	mockImageRepo := mocks.NewMockimageRepo(ctrl)

	service := NewAPODService(logrus.New(), nil, mockImageRepo, mockRepo)

	title := "The Mermaid Nebula Supernova Remnant"
	explanation := "New stars are born from the remnants of dead stars. The gaseous remnant of the gravitational collapse and subsequent death of a very massive star in our Milky Way created the G296.5+10.0 supernova remnant, of which the featured Mermaid Nebula is part. Also known as the Betta Fish Nebula, the Mermaid Nebula makes up part of an unusual subclass of supernova remnants that are two-sided and nearly circular. Originally discovered in X-rays, the filamentary nebula is a frequently studied source also in radio and gamma-ray light.  The blue color visible here originates from doubly ionized oxygen (OIII), while the deep red is emitted by hydrogen gas. The nebula's mermaid-like shape has proven to be useful for measurements of the interstellar magnetic field."
	copyright := "Neil Corke; Text: Natalia Lewandowska (SUNY Oswego)"
	date := time.Now()
	img := bytes.NewReader([]byte("image-data"))

	mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
	mockImageRepo.EXPECT().PutObject(gomock.Any(), gomock.Any()).Return(nil)

	err := service.Save(context.Background(), title, explanation, copyright, &date, img)
	assert.NoError(t, err)
}

func TestAPODService_GetByDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepo(ctrl)
	mockImageRepo := mocks.NewMockimageRepo(ctrl)

	service := NewAPODService(nil, nil, mockImageRepo, mockRepo)

	date := time.Now()

	apod := &model.APOD{
		Id:          "eba86cb3-1136-40c4-a273-e158fe7b1670",
		Title:       "The Mermaid Nebula Supernova Remnant",
		Explanation: "New stars are born from the remnants of dead stars. The gaseous remnant of the gravitational collapse and subsequent death of a very massive star in our Milky Way created the G296.5+10.0 supernova remnant, of which the featured Mermaid Nebula is part. Also known as the Betta Fish Nebula, the Mermaid Nebula makes up part of an unusual subclass of supernova remnants that are two-sided and nearly circular. Originally discovered in X-rays, the filamentary nebula is a frequently studied source also in radio and gamma-ray light.  The blue color visible here originates from doubly ionized oxygen (OIII), while the deep red is emitted by hydrogen gas. The nebula's mermaid-like shape has proven to be useful for measurements of the interstellar magnetic field.",
		Copyright:   "Neil Corke; Text: Natalia Lewandowska (SUNY Oswego)",
		Date:        &date,
	}

	mockRepo.EXPECT().GetByDate(gomock.Any(), &date).Return(apod, nil)

	mockImageRepo.EXPECT().GetImageUrl(apod.Id).Return("http://image-url", nil)

	apodResponse, err := service.GetByDate(context.Background(), &date)
	assert.NoError(t, err)
	assert.NotNil(t, apodResponse)
	assert.Equal(t, apod.Title, apodResponse.Title)
	assert.Equal(t, "http://image-url", apodResponse.Url)
}

func TestAPODService_GetByDate_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepo(ctrl)
	mockImageRepo := mocks.NewMockimageRepo(ctrl)

	service := NewAPODService(nil, nil, mockImageRepo, mockRepo)

	date := time.Now()

	mockRepo.EXPECT().GetByDate(gomock.Any(), &date).Return(nil, errors.New("repository error"))

	apodResponse, err := service.GetByDate(context.Background(), &date)
	assert.Nil(t, apodResponse)
	assert.EqualError(t, err, "failed to get APOD for the date: repository error")
}

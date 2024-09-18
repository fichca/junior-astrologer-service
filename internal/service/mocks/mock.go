// Code generated by MockGen. DO NOT EDIT.
// Source: apod.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	model "github.com/fichca/junior-astrologer-service/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockapodClient is a mock of apodClient interface.
type MockapodClient struct {
	ctrl     *gomock.Controller
	recorder *MockapodClientMockRecorder
}

// MockapodClientMockRecorder is the mock recorder for MockapodClient.
type MockapodClientMockRecorder struct {
	mock *MockapodClient
}

// NewMockapodClient creates a new mock instance.
func NewMockapodClient(ctrl *gomock.Controller) *MockapodClient {
	mock := &MockapodClient{ctrl: ctrl}
	mock.recorder = &MockapodClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockapodClient) EXPECT() *MockapodClientMockRecorder {
	return m.recorder
}

// DownloadImage mocks base method.
func (m *MockapodClient) DownloadImage(imageURL string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadImage", imageURL)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadImage indicates an expected call of DownloadImage.
func (mr *MockapodClientMockRecorder) DownloadImage(imageURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadImage", reflect.TypeOf((*MockapodClient)(nil).DownloadImage), imageURL)
}

// GetAPOD mocks base method.
func (m *MockapodClient) GetAPOD() (*model.APODClientResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPOD")
	ret0, _ := ret[0].(*model.APODClientResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPOD indicates an expected call of GetAPOD.
func (mr *MockapodClientMockRecorder) GetAPOD() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPOD", reflect.TypeOf((*MockapodClient)(nil).GetAPOD))
}

// MockimageRepo is a mock of imageRepo interface.
type MockimageRepo struct {
	ctrl     *gomock.Controller
	recorder *MockimageRepoMockRecorder
}

// MockimageRepoMockRecorder is the mock recorder for MockimageRepo.
type MockimageRepoMockRecorder struct {
	mock *MockimageRepo
}

// NewMockimageRepo creates a new mock instance.
func NewMockimageRepo(ctrl *gomock.Controller) *MockimageRepo {
	mock := &MockimageRepo{ctrl: ctrl}
	mock.recorder = &MockimageRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockimageRepo) EXPECT() *MockimageRepoMockRecorder {
	return m.recorder
}

// GetImageUrl mocks base method.
func (m *MockimageRepo) GetImageUrl(image string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageUrl", image)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageUrl indicates an expected call of GetImageUrl.
func (mr *MockimageRepoMockRecorder) GetImageUrl(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageUrl", reflect.TypeOf((*MockimageRepo)(nil).GetImageUrl), image)
}

// GetImageUrls mocks base method.
func (m *MockimageRepo) GetImageUrls(images []string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageUrls", images)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageUrls indicates an expected call of GetImageUrls.
func (mr *MockimageRepoMockRecorder) GetImageUrls(images interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageUrls", reflect.TypeOf((*MockimageRepo)(nil).GetImageUrls), images)
}

// PutObject mocks base method.
func (m *MockimageRepo) PutObject(image string, data io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObject", image, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutObject indicates an expected call of PutObject.
func (mr *MockimageRepoMockRecorder) PutObject(image, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockimageRepo)(nil).PutObject), image, data)
}

// Mockrepo is a mock of repo interface.
type Mockrepo struct {
	ctrl     *gomock.Controller
	recorder *MockrepoMockRecorder
}

// MockrepoMockRecorder is the mock recorder for Mockrepo.
type MockrepoMockRecorder struct {
	mock *Mockrepo
}

// NewMockrepo creates a new mock instance.
func NewMockrepo(ctrl *gomock.Controller) *Mockrepo {
	mock := &Mockrepo{ctrl: ctrl}
	mock.recorder = &MockrepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockrepo) EXPECT() *MockrepoMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *Mockrepo) GetAll(ctx context.Context) ([]*model.APOD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*model.APOD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockrepoMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*Mockrepo)(nil).GetAll), ctx)
}

// GetByDate mocks base method.
func (m *Mockrepo) GetByDate(ctx context.Context, date *time.Time) (*model.APOD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDate", ctx, date)
	ret0, _ := ret[0].(*model.APOD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDate indicates an expected call of GetByDate.
func (mr *MockrepoMockRecorder) GetByDate(ctx, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDate", reflect.TypeOf((*Mockrepo)(nil).GetByDate), ctx, date)
}

// Save mocks base method.
func (m *Mockrepo) Save(ctx context.Context, apod *model.APOD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, apod)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockrepoMockRecorder) Save(ctx, apod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*Mockrepo)(nil).Save), ctx, apod)
}

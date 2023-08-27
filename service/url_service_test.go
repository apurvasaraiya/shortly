package service

import (
	"errors"
	"reflect"
	"shortly/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestURLServiceEncodeURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks.NewMockRepository(ctrl)
	service := NewURLService(repositoryMock)
	url := "https://google.com"
	expectedId := "1234567891234560"
	GenerateNewUUIDStringFunc = func() string { return expectedId }
	repositoryMock.EXPECT().FetchIDFromURL(url).Return("", nil)
	repositoryMock.EXPECT().SaveURLAndId(url, expectedId).Return(nil)

	id, err := service.EncodeURL(url)
	if err != nil {
		t.Fatalf("error encoding url %s %v", url, err)
	}

	if id != expectedId {
		t.Fatalf("expected: %s, got: %s", expectedId, id)
	}
}

func TestURLServiceEncodeURLShouldFailWhenSaveURLAndIDFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks.NewMockRepository(ctrl)
	service := NewURLService(repositoryMock)
	url := "https://google.com"
	expectedId := "1234567891234560"
	GenerateNewUUIDStringFunc = func() string { return expectedId }
	repositoryMock.EXPECT().FetchIDFromURL(url).Return("", nil)
	repositoryMock.EXPECT().SaveURLAndId(url, expectedId).Return(errors.New("some error"))

	_, err := service.EncodeURL(url)
	if err == nil {
		t.Fatalf("error encoding url %s %v", url, err)
	}
}

func TestURLServiceEncodeURLShouldReturnExistingIdForURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks.NewMockRepository(ctrl)
	service := NewURLService(repositoryMock)
	url := "https://google.com"
	expectedId := "1234567891234560"
	GenerateNewUUIDStringFunc = func() string { return expectedId }
	repositoryMock.EXPECT().FetchIDFromURL(url).Return(expectedId, nil)

	id, err := service.EncodeURL(url)
	if err != nil {
		t.Fatalf("error encoding url %s %v", url, err)
	}

	if id != expectedId {
		t.Fatalf("expected: %s, got: %s", expectedId, id)
	}
}

func TestURLServiceEncodeURLShouldFailIfFetchIDFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks.NewMockRepository(ctrl)
	service := NewURLService(repositoryMock)
	url := "https://google.com"
	expectedId := "1234567891234560"
	GenerateNewUUIDStringFunc = func() string { return expectedId }
	repositoryMock.EXPECT().FetchIDFromURL(url).Return("", errors.New("some error"))

	_, err := service.EncodeURL(url)
	if err == nil {
		t.Fatalf("error encoding url %s %v", url, err)
	}
}

func TestURLServiceFetchURLFromID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks.NewMockRepository(ctrl)
	service := NewURLService(repositoryMock)
	id := "test-id"
	expectedURL := "https://google.com"

	repositoryMock.EXPECT().FetchURLFromID(id).Return(expectedURL, nil)

	url, err := service.FetchURLFromID(id)
	if err != nil {
		t.Fatalf("error fetching url from id %s %v", id, err)
	}

	if url != expectedURL {
		t.Fatalf("expected: %s, got: %s", expectedURL, url)
	}
}

func TestURLServiceFetchURLFromIDShouldReturnErrorWhenFetchUrlFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks.NewMockRepository(ctrl)
	service := NewURLService(repositoryMock)
	id := "test-id"
	expectedURL := "https://google.com"

	repositoryMock.EXPECT().FetchURLFromID(id).Return(expectedURL, errors.New("some error"))

	_, err := service.FetchURLFromID(id)
	if err == nil {
		t.Fatalf("error fetching url from id %s %v", id, err)
	}
}

func TestURLServiceFetchIncrementVisitCountForHostname(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks.NewMockRepository(ctrl)
	service := NewURLService(repositoryMock)
	url := "https://google.com"

	repositoryMock.EXPECT().IncrementCountForHostname("google.com").Return(nil)
	err := service.IncrementVisitCountForHostname(url)
	if err != nil {
		t.Fatalf("error incrementing visit count for url %s %v", url, err)
	}
}

func TestURLServiceFetchIncrementVisitCountForHostnameShouldFailForInvalidURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks.NewMockRepository(ctrl)
	service := NewURLService(repositoryMock)
	url := "{}"

	err := service.IncrementVisitCountForHostname(url)
	if err == nil {
		t.Fatalf("error incrementing visit count for url %s %v", url, err)
	}
}

func TestMetrics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks.NewMockRepository(ctrl)
	service := NewURLService(repositoryMock)
	expectedMetrics := map[string]uint{"google.com": 3, "gmail.com": 2, "youtube.com": 1}

	repositoryMock.EXPECT().FetchMetrics(3).Return(expectedMetrics, nil)
	metrics, err := service.Metrics()
	if err != nil {
		t.Fatalf("error fetching metrics %v", err)
	}

	if !reflect.DeepEqual(expectedMetrics, metrics) {
		t.Fatalf("expected: %v,\ngot: %v", expectedMetrics, metrics)
	}
}

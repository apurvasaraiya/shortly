package handler

import (
	"net/http"
	"net/http/httptest"
	"shortly/mocks"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestEncodeURL(t *testing.T) {
	expectedStatus := http.StatusCreated
	expectedBody := `{"url":"https://google.com","id":"id1"}
`
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := httptest.NewRequest(http.MethodPost, "/encode", strings.NewReader(`{"url":"https://google.com"}`))
	res := httptest.NewRecorder()

	mockURLService := mocks.NewMockURLService(ctrl)
	handler := NewHandler(mockURLService)

	mockURLService.EXPECT().EncodeURL("https://google.com").Return("id1", nil)
	handler.EncodeURL(res, req)
	if res.Code != expectedStatus {
		t.Fatalf("status code; expected: %d, got: %d", expectedStatus, res.Code)
	}

	if res.Body.String() != expectedBody {
		t.Fatalf("body; expected: %s, got: %s", res.Body.String(), expectedBody)
	}
}

func TestRedirect(t *testing.T) {
	expectedStatus := http.StatusMovedPermanently
	expectedRedirectLocation := "https://google.com"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := httptest.NewRequest(http.MethodGet, "/test-id", nil)
	res := httptest.NewRecorder()

	mockURLService := mocks.NewMockURLService(ctrl)
	handler := NewHandler(mockURLService)

	mockURLService.EXPECT().FetchURLFromID("test-id").Return(expectedRedirectLocation, nil)
	mockURLService.EXPECT().IncrementVisitCountForHostname(expectedRedirectLocation).Return(nil)

	handler.Redirect(res, req)
	if res.Code != expectedStatus {
		t.Fatalf("status code; expected: %d, got: %d", expectedStatus, res.Code)
	}

	loc := res.Header().Get("location")
	if loc != expectedRedirectLocation {
		t.Fatalf("redirect location; expected: %s, got: %s", expectedRedirectLocation, loc)
	}
}

func TestMetrics(t *testing.T) {
	expectedStatus := http.StatusOK
	expectedMetricsMap := map[string]uint{"google.com": 1, "youtube.com": 2}
	expectedMetricsJson := `{"google.com":1,"youtube.com":2}
`

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	res := httptest.NewRecorder()

	mockURLService := mocks.NewMockURLService(ctrl)
	handler := NewHandler(mockURLService)

	mockURLService.EXPECT().Metrics().Return(expectedMetricsMap, nil)

	handler.Metrics(res, req)
	if res.Code != expectedStatus {
		t.Fatalf("status code; expected: %d, got: %d", expectedStatus, res.Code)
	}

	if res.Body.String() != expectedMetricsJson {
		t.Fatalf("body; expected: %s, got: %s", expectedMetricsJson, res.Body.String())
	}
}

//TODO: add pending negative scenarios test cases

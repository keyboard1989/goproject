package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestApi(t *testing.T) {
	recoder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api", nil)

	Convey("Test Api", t, func() {
		expectBody := "not found url\n"

		service.Engine().ServeHTTP(recoder, req)

		So(recoder.Code, ShouldEqual, http.StatusNotFound)
		So(recoder.Body.String(), ShouldEqual, expectBody)
	})
}

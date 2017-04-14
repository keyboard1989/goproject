package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCmd(t *testing.T) {
	recoder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/cmd", nil)

	Convey("Test Cmd", t, func() {
		expectBody := "cmd\n"

		service.Engine().ServeHTTP(recoder, req)

		So(recoder.Code, ShouldEqual, http.StatusOK)
		So(recoder.Body.String(), ShouldEqual, expectBody)
	})
}

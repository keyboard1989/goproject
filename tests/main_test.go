package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {

	initConfig()
	initApp()
	ginEngine := engine()

	recoder := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/cmd", nil)
	ginEngine.ServeHTTP(recoder, req)

	Convey("Test Main", t, func() {

	})
}

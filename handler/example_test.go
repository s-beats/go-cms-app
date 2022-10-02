package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestExample(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/example", Example)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/example", nil)
	r.ServeHTTP(w, req)

	b, err := io.ReadAll(w.Body)
	if err != nil {
		t.Error(err)
	}

	strBody := string(b)

	if strBody != "\"example\"" {
		t.Error(err, "want %v,got %v", "\"example\"", strBody)
	}
}
func TestExampleWithTestServer(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/example", Example)
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(fmt.Sprintf("%s/example", ts.URL))
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Error(err, "want %v,got %v", http.StatusOK, res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}

	strBody := string(b)

	if strBody != "\"example\"" {
		t.Error(err, "want %v,got %v", "\"example\"", strBody)
	}
}

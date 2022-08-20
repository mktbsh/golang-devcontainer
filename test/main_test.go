package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mktbsh/golang-devcontainer/router"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	rt := router.Init()

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	rt.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "success", rec.Body.String())
}

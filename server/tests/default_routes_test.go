package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LaMoldy/groupme/server/pkg/router"
	"github.com/stretchr/testify/assert"
)

func TestWelcomeRoute(t *testing.T) {
    router := router.SetupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/api", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
}


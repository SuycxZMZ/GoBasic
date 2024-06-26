package httpServer

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpserve(t *testing.T) {
	t.Run("return peoples score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "players/Pepper", nil)
		response := httptest.NewRecorder()

		PalyerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}

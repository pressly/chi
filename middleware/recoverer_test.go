package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRecoverer(t *testing.T) {
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Use(Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		panic("panic")
	})

	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)
	if w.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Panic not recovered")
	}
}

func TestRecovererAbortHandler(t *testing.T) {
	defer func() {
		rcv := recover()
		if rcv != http.ErrAbortHandler {
			t.Fatalf("http.ErrAbortHandler should not be recovered")
		}
	}()

	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Use(Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		panic(http.ErrAbortHandler)
	})

	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)
}

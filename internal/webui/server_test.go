package webui

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/keyscome/blossom-router/internal/config"
)

func TestAutoDryRunDoesNotCallProvider(t *testing.T) {
	s := Server{Config: config.Config{Providers: map[string]config.Provider{}}}
	body := bytes.NewBufferString(`{"route":"auto","prompt":"design a migration plan","dry_run":true}`)
	req := httptest.NewRequest(http.MethodPost, "/api/run", body)
	w := httptest.NewRecorder()
	s.Handler().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", w.Code, w.Body.String())
	}
	var got runResponse
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got.Route != "strong" || got.Result != "" {
		t.Fatalf("response=%+v", got)
	}
}

func TestRejectsNonLocalAddress(t *testing.T) {
	err := Serve(t.Context(), config.Config{}, "0.0.0.0:7331", false, &bytes.Buffer{})
	if err == nil {
		t.Fatal("expected non-local address to be rejected")
	}
}

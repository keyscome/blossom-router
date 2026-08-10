package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestComplete(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/chat/completions" {
			t.Errorf("path=%s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"choices":[{"message":{"content":"ok"}}]}`))
	}))
	defer s.Close()
	p := OpenAICompatible{BaseURL: s.URL + "/v1", Model: "test"}
	got, err := p.Complete(context.Background(), Request{Prompt: "hello"})
	if err != nil || got != "ok" {
		t.Fatalf("got %q, err %v", got, err)
	}
}

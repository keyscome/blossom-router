package webui

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os/exec"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/keyscome/blossom-router/internal/config"
	"github.com/keyscome/blossom-router/internal/provider"
	"github.com/keyscome/blossom-router/internal/router"
)

//go:embed static/*
var staticFiles embed.FS

type Server struct {
	Config config.Config
	Client *http.Client
}

type runRequest struct {
	Route  string `json:"route"`
	Prompt string `json:"prompt"`
	DryRun bool   `json:"dry_run"`
}

type runResponse struct {
	Route  string `json:"route"`
	Reason string `json:"reason,omitempty"`
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func (s Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/status", s.status)
	mux.HandleFunc("POST /api/run", s.run)
	static, err := fs.Sub(staticFiles, "static")
	if err != nil {
		panic(err)
	}
	mux.Handle("/", http.FileServer(http.FS(static)))
	return securityHeaders(mux)
}

func (s Server) status(w http.ResponseWriter, _ *http.Request) {
	routes := make([]string, 0, len(s.Config.Providers))
	for name := range s.Config.Providers {
		routes = append(routes, name)
	}
	sort.Strings(routes)
	writeJSON(w, http.StatusOK, map[string]any{"routes": routes})
}

func (s Server) run(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var in runRequest
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeJSON(w, http.StatusBadRequest, runResponse{Error: "invalid request"})
		return
	}
	in.Prompt = strings.TrimSpace(in.Prompt)
	if in.Prompt == "" {
		writeJSON(w, http.StatusBadRequest, runResponse{Error: "prompt is empty"})
		return
	}
	route, reason := in.Route, ""
	if route == "auto" {
		decision := router.Choose(in.Prompt)
		route, reason = decision.Route, decision.Reason
	}
	if in.DryRun {
		if in.Route != "auto" {
			writeJSON(w, http.StatusBadRequest, runResponse{Error: "dry run requires the auto route"})
			return
		}
		writeJSON(w, http.StatusOK, runResponse{Route: route, Reason: reason})
		return
	}
	if !validRoute(route) {
		writeJSON(w, http.StatusBadRequest, runResponse{Error: "unknown route"})
		return
	}
	cfg, err := s.Config.Provider(route)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, runResponse{Route: route, Reason: reason, Error: err.Error()})
		return
	}
	client := provider.OpenAICompatible{BaseURL: cfg.BaseURL, APIKey: cfg.APIKey, Model: cfg.Model, Client: s.Client}
	system := ""
	if route == "code" {
		system = "You are a concise coding assistant. Return practical, correct output and avoid unrelated work."
	}
	result, err := client.Complete(r.Context(), provider.Request{Prompt: in.Prompt, System: system})
	if err != nil {
		writeJSON(w, http.StatusBadGateway, runResponse{Route: route, Reason: reason, Error: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, runResponse{Route: route, Reason: reason, Result: result})
}

func validRoute(route string) bool {
	switch route {
	case "local", "cheap", "normal", "strong", "code":
		return true
	default:
		return false
	}
}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self'; script-src 'self'; img-src 'self' data:; connect-src 'self'; base-uri 'none'; frame-ancestors 'none'")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "no-referrer")
		next.ServeHTTP(w, r)
	})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func Serve(ctx context.Context, cfg config.Config, addr string, openBrowser bool, out io.Writer) error {
	if !strings.HasPrefix(addr, "127.0.0.1:") && !strings.HasPrefix(addr, "localhost:") && !strings.HasPrefix(addr, "[::1]:") {
		return fmt.Errorf("refusing non-local address %q; the UI may access configured providers", addr)
	}
	server := &http.Server{Addr: addr, Handler: Server{Config: cfg}.Handler(), ReadHeaderTimeout: 5 * time.Second}
	url := "http://" + addr
	fmt.Fprintf(out, "Blossom Router UI: %s\nPress Ctrl+C to stop.\n", url)
	if openBrowser {
		go func() {
			time.Sleep(150 * time.Millisecond)
			_ = openURL(url)
		}()
	}
	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_ = server.Shutdown(shutdownCtx)
	}()
	err := server.ListenAndServe()
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}

func openURL(url string) error {
	var command string
	var args []string
	switch runtime.GOOS {
	case "darwin":
		command, args = "open", []string{url}
	case "linux":
		command, args = "xdg-open", []string{url}
	default:
		return nil
	}
	return exec.Command(command, args...).Start()
}

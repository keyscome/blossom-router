package app

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/keyscome/blossom-router/internal/config"
	"github.com/keyscome/blossom-router/internal/provider"
	"github.com/keyscome/blossom-router/internal/router"
	"github.com/keyscome/blossom-router/internal/webui"
)

const usage = `Usage: bloom <command> [flags] [prompt]

Commands:
  local   Use the local Ollama-compatible provider
  ask     Use the normal cloud provider
  code    Use the configurable code provider
  strong  Use the strongest configured provider
  auto    Select local, cheap, normal, or strong automatically
  serve   Open the local browser UI

Flags:
  --config PATH   Config file (default ~/.config/blossom/router.yaml)
  --dry-run       Show routing only (auto command)

The prompt may be passed as arguments or via stdin.`

func Run(ctx context.Context, args []string, stdin io.Reader, stdout, stderr io.Writer) error {
	if len(args) == 0 || args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		fmt.Fprintln(stdout, usage)
		return nil
	}
	command := args[0]
	if command == "serve" {
		fs := flag.NewFlagSet(command, flag.ContinueOnError)
		fs.SetOutput(stderr)
		configPath := fs.String("config", "", "config file")
		addr := fs.String("addr", "127.0.0.1:7331", "listen address")
		noOpen := fs.Bool("no-open", false, "do not open the browser")
		if err := fs.Parse(args[1:]); err != nil {
			return err
		}
		cfg, err := config.Load(*configPath)
		if err != nil {
			return err
		}
		return webui.Serve(ctx, cfg, *addr, !*noOpen, stdout)
	}
	valid := map[string]string{"local": "local", "ask": "normal", "code": "code", "strong": "strong", "auto": ""}
	route, ok := valid[command]
	if !ok {
		return fmt.Errorf("unknown command %q\n\n%s", command, usage)
	}
	fs := flag.NewFlagSet(command, flag.ContinueOnError)
	fs.SetOutput(stderr)
	configPath := fs.String("config", "", "config file")
	dryRun := fs.Bool("dry-run", false, "show route only")
	if err := fs.Parse(args[1:]); err != nil {
		return err
	}
	prompt := strings.TrimSpace(strings.Join(fs.Args(), " "))
	if prompt == "" {
		b, err := io.ReadAll(io.LimitReader(stdin, 10<<20))
		if err != nil {
			return err
		}
		prompt = strings.TrimSpace(string(b))
	}
	if prompt == "" {
		return errors.New("prompt is empty")
	}
	if command == "auto" {
		d := router.Choose(prompt)
		route = d.Route
		if *dryRun {
			fmt.Fprintf(stdout, "route: %s\nreason: %s\n", d.Route, d.Reason)
			return nil
		}
		fmt.Fprintf(stderr, "route: %s (%s)\n", d.Route, d.Reason)
	} else if *dryRun {
		return errors.New("--dry-run is only supported by auto")
	}
	cfg, err := config.Load(*configPath)
	if err != nil {
		return err
	}
	p, err := cfg.Provider(route)
	if err != nil {
		return err
	}
	client := provider.OpenAICompatible{BaseURL: p.BaseURL, APIKey: p.APIKey, Model: p.Model}
	system := ""
	if command == "code" {
		system = "You are a concise coding assistant. Return practical, correct output and avoid unrelated work."
	}
	result, err := client.Complete(ctx, provider.Request{Prompt: prompt, System: system})
	if err != nil {
		return err
	}
	fmt.Fprintln(stdout, result)
	return nil
}

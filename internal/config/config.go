package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Provider struct {
	BaseURL   string `yaml:"base_url"`
	APIKey    string `yaml:"api_key"`
	APIKeyEnv string `yaml:"api_key_env"`
	Model     string `yaml:"model"`
}

type Config struct {
	DefaultRoute string              `yaml:"default_route"`
	Providers    map[string]Provider `yaml:"providers"`
}

func Default() Config {
	return Config{DefaultRoute: "normal", Providers: map[string]Provider{
		"local": {BaseURL: "http://localhost:11434/v1", Model: "qwen3:8b"},
	}}
}

func DefaultPath() (string, error) {
	d, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(d, "blossom", "router.yaml"), nil
}

func Load(path string) (Config, error) {
	cfg := Default()
	if path == "" {
		var err error
		path, err = DefaultPath()
		if err != nil {
			return cfg, err
		}
	}
	b, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return applyEnv(cfg), nil
	}
	if err != nil {
		return cfg, fmt.Errorf("read config: %w", err)
	}
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return cfg, fmt.Errorf("parse config: %w", err)
	}
	if cfg.Providers == nil {
		cfg.Providers = map[string]Provider{}
	}
	return applyEnv(cfg), nil
}

func applyEnv(cfg Config) Config {
	for name, p := range cfg.Providers {
		prefix := "BLOSSOM_" + strings.ToUpper(strings.ReplaceAll(name, "-", "_")) + "_"
		if v := os.Getenv(prefix + "BASE_URL"); v != "" {
			p.BaseURL = v
		}
		if v := os.Getenv(prefix + "MODEL"); v != "" {
			p.Model = v
		}
		if v := os.Getenv(prefix + "API_KEY"); v != "" {
			p.APIKey = v
		}
		if p.APIKey == "" && p.APIKeyEnv != "" {
			p.APIKey = os.Getenv(p.APIKeyEnv)
		}
		cfg.Providers[name] = p
	}
	return cfg
}

func (c Config) Provider(name string) (Provider, error) {
	p, ok := c.Providers[name]
	if !ok {
		return p, fmt.Errorf("provider %q is not configured", name)
	}
	if p.BaseURL == "" || p.Model == "" {
		return p, fmt.Errorf("provider %q requires base_url and model", name)
	}
	return p, nil
}

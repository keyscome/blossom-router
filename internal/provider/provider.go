package provider

import "context"

type Request struct {
	Prompt string
	System string
}

type Provider interface {
	Complete(context.Context, Request) (string, error)
}

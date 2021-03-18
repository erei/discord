package discord

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

// Package discord partially implements Discord's HTTP API to streamline the OAuth2 flow for consumers.

const (
	Version = "v0.1.0"
	repoURL = "https://github.com/erei/discord"
)

var (
	Endpoint = oauth2.Endpoint{
		AuthURL:  "https://discord.com/api/oauth2/authorize",
		TokenURL: "https://discord.com/api/oauth2/token",
	}

	ErrMissingRequiredArgument = errors.New("discord: missing a required argument")

	defaultScopes    = []string{"identify"}
	defaultUserAgent = fmt.Sprintf("DiscordBot (%s, %s)", repoURL, Version)
)

// Discord is an API client.
type Discord struct {
	client     *http.Client
	userConfig *oauth2.Config
}

// New creates a new Discord client.
func New(opts ...Option) (*Discord, error) {
	d := new(Discord)

	for _, opt := range opts {
		opt(d)
	}

	if err := d.defaults(); err != nil {
		return nil, err
	}

	return d, nil
}

// Exchange performs an OAuth2 exchange.
func (d *Discord) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	ctx = context.WithValue(ctx, oauth2.HTTPClient, d.client)
	return d.userConfig.Exchange(ctx, code)
}

// AuthCodeURL returns a URL to authenticate a user with.
func (d *Discord) AuthCodeURL(state string) string {
	return d.userConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (d *Discord) defaults() error {
	if d.client == nil {
		d.client = http.DefaultClient
	}

	if d.userConfig.Endpoint.AuthURL == "" ||
		d.userConfig.Endpoint.TokenURL == "" {
		d.userConfig.Endpoint = Endpoint
	}

	if d.userConfig.ClientID == "" ||
		d.userConfig.ClientSecret == "" ||
		d.userConfig.RedirectURL == "" {
		return ErrMissingRequiredArgument
	}

	if d.userConfig.Scopes == nil {
		d.userConfig.Scopes = defaultScopes
	}

	return nil
}

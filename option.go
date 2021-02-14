package discord

import (
	"net/http"

	"golang.org/x/oauth2"
)

// Option configures a Discord.
type Option func(*Discord)

// WithHTTPClient specifies the *http.Client to use.
func WithHTTPClient(cli *http.Client) Option {
	return func(d *Discord) {
		d.client = cli
	}
}

// WithOAuth2Config specifies the *oauth2.Config to use.
func WithOAuth2Config(conf *oauth2.Config) Option {
	return func(d *Discord) {
		d.userConfig = conf
	}
}

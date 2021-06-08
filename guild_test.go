package discord

import (
	"context"
	"testing"

	"golang.org/x/oauth2"
)

func TestGetGuilds(t *testing.T) {
	oauth2Conf := &oauth2.Config{
		ClientID:     "xD",
		ClientSecret: "xD",
		RedirectURL:  "xD",
		Scopes:       []string{"guilds"},
	}

	d, err := New(WithOAuth2Config(oauth2Conf))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	guilds, er := d.GuildsFromToken(ctx, testToken)
	if er != nil {
		err, ok := er.(*Error)
		if ok {
			t.Fatalf("API error: %s\n", err.Error())
		}

		t.Fatal(err)
	}

	if len(guilds) < 1 {
		t.Fatal("no guilds were returned")
	}
}

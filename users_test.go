package discord

import (
	"context"
	"os"
	"testing"

	"golang.org/x/oauth2"
)

const testID = "67803413605253120"

var testToken = os.Getenv("DISCORD_TEST_OAUTH_TOKEN")

func TestCanGetUser(t *testing.T) {
	oauth2Conf := &oauth2.Config{
		ClientID:     "xD",
		ClientSecret: "xD",
		RedirectURL:  "xD",
	}

	d, err := New(WithOAuth2Config(oauth2Conf))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	user, er := d.UserFromToken(ctx, testToken)
	if er != nil {
		err, ok := er.(*Error)
		if ok {
			t.Fatalf("API error: %s\n", err.Error())
		}

		t.Fatal(err)
	}

	if user.ID != testID {
		t.Fatal("ids don't match")
	}
}

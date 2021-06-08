package discord

import (
	"context"
	"encoding/json"
	"net/http"
)

// User is a partial definition of a Discord User defined by the API.
type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	AvatarHash    string `json:"avatar"`
	Bot           bool   `json:"bot,omitempty"`
	System        bool   `json:"system,omitempty"`
	MFAEnabled    bool   `json:"mfa_enabled,omitempty"`
	Locale        string `json:"locale,omitempty"`
	Flags         int    `json:"flags,omitempty"`
	PremiumType   int    `json:"premium_type,omitempty"`
	PublicFlags   int    `json:"public_flags,omitempty"`
}

// Avatar returns a URL to the User's avatar, in PNG format.
func (u *User) Avatar() string {
	return rootCDN + "avatars/" + u.ID + "/" + u.AvatarHash + ".png"
}

// UserFromToken gets a Discord User from an OAuth token.
func (d *Discord) UserFromToken(ctx context.Context, token string) (*User, error) {
	url := root + "/users/@me"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", defaultUserAgent)
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := d.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, newError(res)
	}

	var u *User
	if err := json.NewDecoder(res.Body).Decode(&u); err != nil {
		return nil, err
	}

	return u, nil
}

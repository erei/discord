package discord

import (
	"context"
	"encoding/json"
	"net/http"
)

// Guild is a partial guild that the authenticated user belongs to.
type Guild struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	IconHash    string   `json:"icon"`
	Owner       bool     `json:"owner"`
	Permissions string   `json:"permissions"`
	Features    []string `json:"features"`
}

// Icon returns the URL to a Guild's icon.
func (g *Guild) Icon() string {
	return rootCDN + "icons/" + g.ID + "/" + g.IconHash + ".png"
}

// GuildsFromToken fetches the list of Guilds a Discord user belongs to.
// Requires the `guilds` OAuth2 scope.
func (d *Discord) GuildsFromToken(ctx context.Context, token string) ([]*Guild, error) {
	url := root + "/users/@me/guilds"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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

	var g []*Guild
	if err := json.NewDecoder(res.Body).Decode(&g); err != nil {
		return nil, err
	}

	return g, nil
}

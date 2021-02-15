package discord

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Error is an error returned from the Discord API.
// It implements the Error interface.
type Error struct {
	HTTPStatus int    `json:"-"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
}

func (e *Error) Error() string {
	var sb strings.Builder

	if e.HTTPStatus != 0 {
		sb.WriteString(fmt.Sprintf("HTTP %d: ", e.HTTPStatus))
	}

	if e.Code != 0 {
		sb.WriteString(fmt.Sprintf("%d: ", e.Code))
	}

	if e.Message != "" {
		sb.WriteString(e.Message)
	}

	return sb.String()
}

// helper function that makes an Error from an http.Response
func newError(res *http.Response) error {
	var e *Error

	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
		return err
	}

	e.HTTPStatus = res.StatusCode
	return e
}

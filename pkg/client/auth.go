package client

import (
	"net/http"

	"github.com/puppetlabs/relay/pkg/errors"
	"github.com/puppetlabs/relay/pkg/model"
)

type CreateTokenParameters struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateTokenResponse struct {
	Token *model.Token `json:"token"`
}

func (c *Client) CreateToken(email string, password string) errors.Error {

	params := CreateTokenParameters{
		Email:    email,
		Password: password,
	}

	response := &CreateTokenResponse{}
	if err := c.Request(
		WithMethod(http.MethodPost),
		WithPath("/auth/sessions"),
		WithBody(params),
		WithResponseInto(response),
	); err != nil {
		return err
	}

	if err := c.storeToken(response.Token); err != nil {
		return errors.NewClientInternalError().WithCause(err)
	}

	return nil
}

func (c *Client) InvalidateToken() errors.Error {
	type deleteResponse struct {
		Success bool `json:"success"`
	}

	dr := &deleteResponse{}

	// Dont propagate error: if existing token is invalid endpoint will 401. Not sure this is
	// good behavior but it's true nonetheless
	c.Request(
		WithMethod(http.MethodDelete),
		WithPath("/auth/sessions"),
		WithResponseInto(dr),
	)

	if err := c.clearToken(); err != nil {
		return errors.NewClientInternalError().WithCause(err)
	}

	return nil
}

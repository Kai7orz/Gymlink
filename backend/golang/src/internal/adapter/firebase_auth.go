package adapter

import (
	"context"

	"firebase.google.com/go/auth"
)

type authClient struct {
	client *auth.Client
}

func NewAuthClient(client *auth.Client) *authClient {
	return &authClient{client: client}
}

func (a *authClient) VerifyUser(ctx context.Context, idToken string) (*auth.Token, error) {
	return a.client.VerifyIDToken(ctx, idToken)
}

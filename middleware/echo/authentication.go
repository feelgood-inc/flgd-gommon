package middleware

import (
	"context"
	firebase "firebase.google.com/go/v4"
)

type AuthenticationMiddleware struct {
	firebase *firebase.App
}

func NewAuthenticationMiddleware(firebase *firebase.App) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		firebase: firebase,
	}
}

func (m *AuthenticationMiddleware) ValidateJWT(ctx context.Context, IDToken string) error {
	client, err := m.firebase.Auth(ctx)
	if err != nil {
		return err
	}
	_, err = client.VerifyIDToken(ctx, IDToken)
	if err != nil {
		return err
	}

	return nil
}

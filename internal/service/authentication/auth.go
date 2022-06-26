package authentication

import (
	"context"
	"ms-hw/internal/core"
	"ms-hw/internal/core/aggregate"
)

type Credential interface {
	GetCredentials(ctx context.Context, person *aggregate.Person) (res []*core.Credentials, err error)
}

type AuthService struct {
	creds Credential
}

func NewAuthService(creds Credential) *AuthService {
	return &AuthService{creds: creds}
}

func (a *AuthService) Authenticate(ctx context.Context, person *aggregate.Person) (string, error) {
	credentials, err := a.creds.GetCredentials(ctx, person)
	if err != nil {
		return "не ок", err
	}

	if len(credentials) == 0 {
		return "не ок", err
	}

	return "access granted", nil
}

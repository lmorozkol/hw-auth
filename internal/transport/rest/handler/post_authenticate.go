package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"ms-hw/internal/core/aggregate"
	"ms-hw/internal/service/authentication"
	"ms-hw/pkg/common/wraps"
	"net/http"
)

const authenticateError = "authenticateError. Message is not valid"

type AuthHandler struct {
	authService *authentication.AuthService
}

func NewAuthHandler(authService *authentication.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (as *AuthHandler) GrantAccess(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var person *aggregate.Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		wraps.WrapErrorBadRequest(w, err)
		return
	}

	if err = as.Validate(person); err != nil {
		wraps.WrapErrorBadRequest(w, err)
		return
	}

	res, err := as.authService.Authenticate(ctx, person)
	if err != nil {
		wraps.WrapErrorServer(w, err)
		return
	}

	var m = map[string]interface{}{
		"status": res,
	}

	wraps.WrapOK(w, m)
}

func (as *AuthHandler) Validate(person *aggregate.Person) error {
	if person == nil {
		return fmt.Errorf(authenticateError)
	}

	if person.Login == "" || person.Password == "" {
		return fmt.Errorf(authenticateError)
	}
	return nil
}

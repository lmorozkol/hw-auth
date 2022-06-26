package handler

import (
	"context"
	"ms-hw/internal/core/aggregate"
	"ms-hw/internal/service/authentication"
	"ms-hw/pkg/common/err_wrapper"
	"net/http"
)

type AuthHandler struct {
	authService *authentication.AuthService
}

func NewAuthHandler(authService *authentication.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (as *AuthHandler) GrantAccess(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	res, err := as.authService.Authenticate(ctx, &aggregate.Person{})
	if err != nil {

	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   res,
	}
	err_wrapper.WrapOK(w, m)
}

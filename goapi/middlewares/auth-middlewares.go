package middlewares

import (
	"context"
	"fmt"
	"goapi/utils"
	"net/http"
)

type Email string
type Uid string

const (
	e Email = "email"
	u Uid   = "uid"
)

func Authentication(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientToken := r.Header.Get("token")
		if clientToken == "" {
			utils.HandleApiErrors(w, http.StatusUnauthorized, "")
			return
		}

		claims, err := utils.ValidateToken(clientToken)
		if err != "" {
			utils.HandleApiErrors(w, http.StatusForbidden, "")
			return
		}

		fmt.Printf("aqui: %v", claims)

		ctx := r.Context()
		ctx = context.WithValue(ctx, e, claims.Email)
		ctx = context.WithValue(ctx, u, claims.Uid)
		req := r.WithContext(ctx)
		handler.ServeHTTP(w, req)
	})
}

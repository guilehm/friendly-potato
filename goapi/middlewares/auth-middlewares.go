package middlewares

import (
	"context"
	"goapi/utils"
	"net/http"
)

type Email string
type Uid string

const (
	e Email = "email"
	u Uid   = "uid"
)

func Authentication(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
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

		ctx := r.Context()
		ctx = context.WithValue(ctx, e, claims.Email)
		ctx = context.WithValue(ctx, u, claims.Uid)
		req := r.WithContext(ctx)
		handler.ServeHTTP(w, req)
	})
}

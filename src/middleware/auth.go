package middleware

import (
	"context"
	"gawds/src/utils"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metadata, err := utils.ExtractTokenMetadata(r)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), "id", metadata.UserId)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

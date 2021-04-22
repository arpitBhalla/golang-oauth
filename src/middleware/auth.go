package middleware

import (
	"context"
	"gawds/src/utils"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenData, er := utils.ExtractTokenMetadata(r)
		print(er.Error())

		// print(tokenData.AccessUuid)

		ctx := context.WithValue(r.Context(), tokenData, tokenData)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

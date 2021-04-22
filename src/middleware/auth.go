package middleware

import (
	"gawds/src/utils"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		strKey, _ := utils.ExtractToken(r)
		print(strKey.Valid)
		// ctx := context.WithValue(r.Context(), "id", "strKey")

		next.ServeHTTP(w, r) //.WithContext(ctx))

	})
}

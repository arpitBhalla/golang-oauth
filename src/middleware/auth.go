package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secretKey"))

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")
		strKey := fmt.Sprintf("%v", session.Values["id"])

		ctx := context.WithValue(r.Context(), "id", strKey)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

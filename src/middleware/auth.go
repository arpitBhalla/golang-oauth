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
			return
		}
		userid, err := utils.FetchAuth(metadata)
		if err != nil {
			return
		}
		print(userid)
		// if er != nil {
		// 	http.Error(w, er.Error(), http.StatusInternalServerError)

		// 	next.ServeHTTP(w, r)
		// 	return
		// }
		// print(er.Error())
		// print(tokenData.AccessUuid)

		// print(tokenData.AccessUuid)

		ctx := context.WithValue(r.Context(), "id", "6081466e4f3bda19877935ee")

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

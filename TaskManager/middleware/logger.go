package middleware

import (
	"fmt"
	"net/http"
)

func Logger(next http.Handler) http.Handler{
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Method :",
						r.Method,
					"Path :",
				r.URL.Path,)
				next.ServeHTTP(w,r)
		},
		
	)
}
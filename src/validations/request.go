package validations

import "net/http"

func isRequestValid(r *http.Request) bool {
	result := true
	header := r.Header.Get("Authorization")

	result = result && len(header) == 0

	return result
}

func RequestValidator(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if isRequestValid(req) {
			next.ServeHTTP(w, req)
		} else {
			http.Error(w, "Invalid request", 400)
		}
	})
}

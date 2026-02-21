// Package util - Handle OPTIONS Request which is preflight Request
package util

import "net/http"

func GlobalRouter(mux *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}
		mux.ServeHTTP(w, r)
	}

	handler := http.HandlerFunc(handleAllReq)
	return handler
}

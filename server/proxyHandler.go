package server

import (
	"net/http"

	j "github.com/ldelossa/json"
)

// proxyHandler receives an http request and passes it to FFproxy for further
// proxying to the client stream. if the hostname in the http request does not match
// a registered grpc stream a error will be returned to the client.
func proxyHandler(f *FFproxy) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// hand http.Request to FFproxy
		resp, err := f.DoProxy(r)
		if err != nil {
			j.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// write resp to response writer
		err := resp.Write(w)
		if err != nil {
			j.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	})
}

package gateway

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func HomePage(ctx context.Context, params []string, domains []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := r.URL.Query()
		param := keys.Get("param")
		for i, p := range params {
			if param == p {
				resp, err := http.Get(domains[i])
				if err != nil {
					log.Fatalln(err)
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatalln(err)
				}
				sb := string(body)

				w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
				w.WriteHeader(resp.StatusCode)
				io.WriteString(w, sb)
			}
		}
	}
}

func New(ctx context.Context, params []string, domains []string) (http.Handler, error) {
	mux := http.NewServeMux()
	mux.Handle("/", HomePage(ctx, params, domains))
	return mux, nil
}

package gateway

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func HomePage(host string, param string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("host")
		fmt.Println(host)
		fmt.Println("host")
		fmt.Println(param)
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"success": "yes"}`)
	}
}

func New(ctx context.Context, host string, param string) (http.Handler, error) {
	mux := http.NewServeMux()
	mux.Handle("/", HomePage(host, param))
	return mux, nil
}

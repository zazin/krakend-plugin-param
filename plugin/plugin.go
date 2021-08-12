package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/zazin/krakend-plugin-param/gateway"
	"net/http"
)

func init() {
	fmt.Println("krakend-grpc-post plugin loaded!!!")
}

var ClientRegisterer = registerer("test-plugin")

type registerer string

func (r registerer) RegisterClients(f func(
	name string,
	handler func(context.Context, map[string]interface{}) (http.Handler, error),
)) {
	f(string(r), func(ctx context.Context, extra map[string]interface{}) (http.Handler, error) {
		cfg := parse(extra)
		if cfg == nil {
			return nil, errors.New("wrong config")
		}
		if cfg.name != string(r) {
			return nil, fmt.Errorf("unknown register %s", cfg.name)
		}
		return gateway.New(ctx, cfg.host, cfg.param)
	})
}

func parse(extra map[string]interface{}) *opts {
	name, ok := extra["name"].(string)
	if !ok {
		return nil
	}

	host, ok := extra["host"].(string)
	if !ok {
		return nil
	}

	param, ok := extra["param"].(string)
	if !ok {
		return nil
	}

	return &opts{
		name:  name,
		host:  host,
		param: param,
	}
}

type opts struct {
	name  string
	host  string
	param string
}

func main() {

}

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
		return gateway.New(ctx, cfg.params, cfg.domains)
	})
}

func parse(extra map[string]interface{}) *opts {
	name, ok := extra["name"].(string)
	if !ok {
		return nil
	}

	rawParams, ok := extra["params"]
	if !ok {
		return nil
	}
	pr, ok := rawParams.([]interface{})
	if !ok || len(pr) < 2 {
		return nil
	}
	params := make([]string, len(pr))
	for i, e := range pr {
		params[i] = e.(string)
	}

	rawDomain, ok := extra["domains"]
	if !ok {
		return nil
	}
	dm, ok := rawDomain.([]interface{})
	if !ok || len(dm) < 2 {
		return nil
	}
	domains := make([]string, len(dm))
	for i, e := range dm {
		domains[i] = e.(string)
	}

	return &opts{
		name:    name,
		params:  params,
		domains: domains,
	}
}

type opts struct {
	name    string
	params  []string
	domains []string
}

func main() {}

package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

func parse(v interface{}) error {
	if err := env.Parse(v); err != nil {
		return fmt.Errorf("error parsing env: %s", err.Error())
	}
	return nil
}

func parsePrefix(prefix string, v interface{}) error {
	if err := env.ParseWithOptions(v, env.Options{Prefix: prefix}); err != nil {
		return fmt.Errorf("error parsing env with prefix %s: %s", prefix, err.Error())
	}
	return nil
}

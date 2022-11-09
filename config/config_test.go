package config

import (
	"errors"
	"fmt"
	"testing"

	"github.com/caarlos0/env/v6"
)

func Test_New(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		wantPort := 3333
		t.Setenv("PORT", fmt.Sprint(wantPort))

		wantEnv := "dev"
		t.Setenv("TODO_ENV", fmt.Sprint(wantEnv))

		got, err := New()
		if err != nil {
			t.Fatalf("cannot create confg: %v", err)
		}
		if got.Port != wantPort {
			t.Errorf("want %d, but %d", wantPort, got.Port)
		}
		if got.Env != wantEnv {
			t.Errorf("want %s, but %s", wantEnv, got.Env)
		}
	})

	t.Run("異常系", func(t *testing.T) {
		old := envParse
		defer func() { envParse = old }()

		wantErr := errors.New("env.Parse error")
		envParse = func(v interface{}, opts ...env.Options) error {
			return wantErr
		}

		got, err := New()

		if err != wantErr {
			t.Errorf("want %s, but %s", wantErr, err)
		}

		if got != nil {
			t.Errorf("want nil, but %v", got)
		}
	})
}

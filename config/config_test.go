package config

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_New(t *testing.T) {
	wantPort := 3333
	t.Setenv("PORT", fmt.Sprint(wantPort))

	got, err := New()
	if err != nil {
		t.Fatalf("cannot create confg: %v", err)

	}

	r, err := json.Marshal(got)

	if err != nil {
		t.Log("[DEBUG] fail Marshal")
	}
	t.Logf("[DEBUG] got is: %s", string(r))

	if got.Port != wantPort {
		t.Errorf("want %d, but %d", wantPort, got.Port)
	}

	wantEnv := "dev"
	if got.Env != wantEnv {
		t.Errorf("want %s, but %s", wantEnv, got.Env)
	}
}

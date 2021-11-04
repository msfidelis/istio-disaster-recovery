package main

import (
	"testing"

	"cc-api/pkg/configuration"
)

func TestPkgConfigurationLoad(t *testing.T) {

	t.Run("Load JSON File", func(t *testing.T) {
		configs := configuration.Load()
		got := configs.Version
		want := "v1"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Assert Correct Env File", func(t *testing.T) {
		configs := configuration.Load()
		got := configs.Env
		want := "test"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}

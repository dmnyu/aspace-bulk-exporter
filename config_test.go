package aspace_bulk_exporter

import (
	"github.com/nyudlts/go-aspace"
	"testing"
)

func TestConfig(t *testing.T) {
	config := "/etc/go-aspace.yml"
	t.Run("Test config exists", func(t *testing.T) {
		if ConfigExists(config) != true {
			t.Errorf("Config does not exist")
		}
	})

	var creds map[string]aspace.Creds

	t.Run("Test config contains dev", func(t *testing.T) {
		var err error
		creds, err = GetAspaceCreds(config)
		if err != nil {
			t.Errorf(err.Error())
		}
		if len(creds) < 1 {
			t.Errorf("Creds file did not contain any creds")
		}
	})

	t.Run("Test get environment keys", func(t *testing.T) {
		keys := GetAspaceEnvironments(creds)
		if len(keys) < 1 {
			t.Errorf("Creds file did not contain any environement keys")
		}
		t.Logf("%v", keys)
	})
}
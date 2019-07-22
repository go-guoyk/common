package common

import (
	"os"
	"testing"
)

type TestConfig struct {
	Field1 struct {
		Field11 []string
	}
	Field2 struct {
		Field21 bool
		Field22 string `default:"$TEST_FIELD22|val22"`
		Field23 string `default:"$TEST_FIELD23|val22"`
	}
}

func TestLoadYAMLConfigFile(t *testing.T) {
	var err error
	var cfg TestConfig

	os.Setenv("TEST_FIELD23", "val23")

	if err = LoadYAMLConfigFile("testdata/config.yml", &cfg); err != nil {
		t.Fatal(err)
	}

	if len(cfg.Field1.Field11) != 2 {
		t.Fatal("field11 len not match")
	}
	if !cfg.Field2.Field21 {
		t.Fatal("field21 not true")
	}
	if cfg.Field2.Field22 != "val22" {
		t.Fatal("field22 not set")
	}
	if cfg.Field2.Field23 != "val23" {
		t.Fatal("field23 not set")
	}

	t.Log(cfg)
}

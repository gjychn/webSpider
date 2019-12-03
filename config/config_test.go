package config

import (
	"testing"
)

func TestLoadConfig(t *testing.T)  {
	_, err := LoadConfig("../config/config.yaml")
	if err != nil {
		t.Error("conf err: ", err)
	}

	_, err = LoadConfig("../g.yaml")
	if err != nil {
		t.Error("conf err: ", err)
	}
}

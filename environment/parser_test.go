package environment

import (
	"testing"
)

func TestParser(t *testing.T) {
	var conf *Config
	conf = Parse()

	if conf.DbName == "" {
		t.Error("No Database Name set in Configfile")
	}
	if conf.Port != 5432 {
		t.Log("Are you sure you want to use a different port as 5432?")
	}

}
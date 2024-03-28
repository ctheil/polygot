package config_test

import (
	"reflect"
	"testing"

	"calebtheil.com/polygot/pkg/config"
)

func getOpts(args []string) *config.Opts {
	return &config.Opts{
		Args:   args,
		Config: "",
		Pwd:    "",
	}
}

func testConfig(t *testing.T, args []string, expectedArgs []string, operation config.Operation) {
	opts := getOpts(args)
	config, err := config.NewConfig(opts)
	if err != nil {
		t.Errorf("expected to get no errors %v", err)
	}

	if !reflect.DeepEqual(expectedArgs, config.Args) {
		t.Errorf("expected args to be %+v but got %+v", expectedArgs, config.Args)
	}
	if config.Operation != operation {
		t.Errorf("operation expected was %v but got %v", operation, config.Operation)
	}
}

func TestConfigPrint(t *testing.T) {
	testConfig(t, []string{}, []string{}, config.Print)
}

func TestConfigPrintKey(t *testing.T) {
	testConfig(t, []string{"foo"}, []string{"foo"}, config.Print)
}

func TestConfigAddKeyValue(t *testing.T) {
	testConfig(t, []string{"add", "foo", "bar"}, []string{"foo", "bar"}, config.Add)
}

func TestConfigRemoveKey(t *testing.T) {
	testConfig(t, []string{"remove", "foo"}, []string{"foo"}, config.Remove)
}

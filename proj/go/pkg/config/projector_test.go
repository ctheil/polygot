package config_test

import (
	"testing"

	"calebtheil.com/polygot/pkg/config"
)

func getData() *config.Data {
	return &config.Data{
		Projector: map[string]map[string]string{
			"/": {
				"foo": "bar1",
				"fem": "is_great",
			},
			"/foo": {
				"foo": "bar2",
			},
			"/foo/bar": {
				"foo": "bar3",
			},
		},
	}
}

func getProjector(pwd string, data *config.Data) *config.Projector {
	return &config.Projector{
		Config: &config.Config{
			Args:      []string{},
			Operation: config.Print,
			Pwd:       pwd,
			Config:    "Hello fem",
		},
		Data: data,
	}
}

func test(t *testing.T, p *config.Projector, key, value string) {
	v, ok := p.GetValue(key)
	if !ok {
		t.Errorf("expected to find value \"%v\"", value)
	}
	if v != value {
		t.Errorf("expected %v to be \"%v\"", value, v)
	}
}

func TestGetValue(t *testing.T) {
	data := getData()
	p := getProjector("/foo/bar", data)
	test(t, p, "foo", "bar3")
	test(t, p, "fem", "is_great")
}

func TestSetValue(t *testing.T) {
	data := getData()
	p := getProjector("/foo/bar", data)

	p.SetValue("foo", "baz")
	p.SetValue("fem", "is_better_than_great") // set in /foo/bar then change dirs
	test(t, p, "foo", "baz")

	p = getProjector("/", data) // set dir to /
	p.SetValue("fem", "is_great")
	test(t, p, "fem", "is_great")
	p.SetValue("fem", "is_awesome")
	test(t, p, "fem", "is_awesome")
}

func TestRemoveValue(t *testing.T) {
	data := getData()
	p := getProjector("/", data)
	p.RemoveValue("foo")
	_, ok := p.GetValue("foo")
	if ok {
		t.Error("expected not to find value for \"foo\"")
	}

	// change dirs to check that foo is still accessible one level up
	p = getProjector("/foo", data)
	_, ok = p.GetValue("foo")
	if !ok {
		t.Error("expected to find value of \"foo\"")
	}
}

package config

import (
	"encoding/json"
	"os"
	"path"
)

// map returns 2 values- value && ok
type Data struct {
	Projector map[string]map[string]string `json:"Projector"`
}

type Projector struct {
	Config *Config
	Data   *Data
}

func getDefaultData() Data {
	return Data{}
}

func (p *Projector) GetValue(key string) (string, bool) {
	curr := p.Config.Pwd
	prev := ""

	out := ""
	found := false
	for curr != prev {
		if dir, ok := p.Data.Projector[curr]; ok {
			if value, ok := dir[key]; ok {
				// found value
				out = value
				found = true
				break
			}
		}

		prev = curr
		curr = path.Dir(curr)
	}
	return out, found
}

func (p *Projector) GetValueAll() map[string]string {
	out := map[string]string{}

	paths := []string{}
	curr := p.Config.Pwd
	prev := ""

	for curr != prev {
		paths = append(paths, curr)

		prev = curr
		curr = path.Dir(curr)
	}
	for i := len(paths) - 1; i >= 0; i-- {
		if dir, ok := p.Data.Projector[paths[i]]; ok {
			for k, v := range dir {
				out[k] = v
			}
		}
	}

	return out
}

func (p *Projector) SetValue(key, value string) {
	pwd := p.Config.Pwd
	if _, ok := p.Data.Projector[pwd]; !ok {
		p.Data.Projector[p.Config.Pwd] = map[string]string{}
	}
	p.Data.Projector[pwd][key] = value
}

func (p *Projector) RemoveValue(key string) {
	if _, ok := p.Data.Projector[p.Config.Pwd]; ok {
		delete(p.Data.Projector[p.Config.Pwd], key)
	}
}

func defaultProjector(c *Config) *Projector {
	return &Projector{
		Config: c,
		Data: &Data{
			Projector: map[string]map[string]string{},
		},
	}
}

func NewProjector(c *Config) *Projector {
	if _, err := os.Stat(c.Config); err != nil {
		file, err := os.ReadFile(c.Config)
		if err != nil {
			return defaultProjector(c)
		}

		var data Data
		err = json.Unmarshal(file, &data)
		if err != nil {
			return defaultProjector(c)
		}
		return &Projector{
			Data:   &data,
			Config: c,
		}
	}
	return defaultProjector(c)
}

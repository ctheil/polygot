package config

import (
	"encoding/json"
	"fmt"
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

	fmt.Printf("Appending value %v:%v to %+v\n", key, value, p.Data)
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

func (p *Projector) Save() error {
	dir := path.Dir(p.Config.Config)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0o755)
		if err != nil {
			return err
		}
	}
	jsonString, err := json.Marshal(p.Data)
	if err != nil {
		return err
	}
	target := path.Join(dir, "projector.json")
	fmt.Printf("Saving %+v to %v", p.Data, target)
	if err := os.WriteFile(target, jsonString, 0o755); err != nil {
		return err
	}

	fmt.Println("\nSaved")
	return nil
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
	if _, err := os.Stat(c.Config); err == nil {
		file, err := os.ReadFile(c.Config)
		if err != nil {
			fmt.Printf("Error reading config file: %v", err)
			return defaultProjector(c)
		}

		var data Data
		err = json.Unmarshal(file, &data)
		if err != nil {
			fmt.Printf("Error unmatshaling data: %v", err)
			return defaultProjector(c)
		}
		return &Projector{
			Data:   &data,
			Config: c,
		}
	}
	return defaultProjector(c)
}

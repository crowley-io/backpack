package engine

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Configuration contains pack runtime instructions.
type Configuration struct {
	Backpack struct {
		Prehooks  []string `yaml:"pre-hooks"`
		Posthooks []string `yaml:"post-hooks"`
		Execute   []string
	}
}

func (c Configuration) Prehooks() []string {
	return c.Backpack.Prehooks
}

func (c Configuration) Posthooks() []string {
	return c.Backpack.Posthooks
}

func (c Configuration) Execute() []string {
	return c.Backpack.Execute
}

// NewConfiguration return an empty Configuration.
func NewConfiguration() *Configuration {
	return &Configuration{}
}

// ParseConfiguration parse a file path and inflate a new Configuration
func ParseConfiguration(path string) (*Configuration, error) {

	c := NewConfiguration()
	b, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(b, c); err != nil {
		return nil, err
	}

	return c, nil
}
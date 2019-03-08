package workflow

import (
	"github.com/puppetlabs/nebula/pkg/errors"
	"github.com/puppetlabs/nebula/pkg/workflow/runner"
	"gopkg.in/yaml.v2"
)

type Workflow struct {
	Version   string     `yaml:"version"`
	Name      string     `yaml:"name"`
	Variables []Variable `yaml:"variables"`
	Actions   []Action   `yaml:"actions"`
	Stage     Stage      `yaml:"stages"`
}

type Trigger struct {
	Action string `yaml:"action"`
	Branch string `yaml:"branch"`
}

type Stage struct {
	Steps   []string  `yaml:"steps"`
	StartOn []Trigger `yaml:"start_on"`

	Actions []Action
}

type Action struct {
	Name string                      `yaml:"name"`
	Kind string                      `yaml:"kind"`
	Spec map[interface{}]interface{} `yaml:"spec"`

	loadedRunner runner.ActionRunner
}

func (a Action) Runner() runner.ActionRunner {
	return a.loadedRunner
}

func (a *Action) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw map[string]interface{}

	if err := unmarshal(&raw); err != nil {
		return err
	}

	name, ok := raw["name"].(string)
	if !ok {
		return errors.NewWorkflowActionDecodeError("`name` was not a string")
	}
	a.Name = name

	kind, ok := raw["kind"].(string)
	if !ok {
		return errors.NewWorkflowActionDecodeError("`kind` was not a string")
	}
	a.Kind = kind

	if _, ok := raw["spec"]; ok {
		a.Spec = raw["spec"].(map[interface{}]interface{})
	}

	r, err := runner.NewRunner(runner.RunnerKind(a.Kind))
	if err != nil {
		return err
	}

	b, err := yaml.Marshal(raw)
	if err != nil {
		return err
	}

	if err := r.Decoder().Decode(b); err != nil {
		return err
	}

	a.loadedRunner = r

	return nil
}

type Variable struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}
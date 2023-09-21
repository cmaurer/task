package taskfile

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type If struct {
	Static string
	Sh     string
}

func (f *If) UnmarshalYAML(node *yaml.Node) error {
	switch node.Kind {

	case yaml.ScalarNode:
		var str string
		if err := node.Decode(&str); err != nil {
			return err
		}
		f.Static = str
		return nil

	case yaml.MappingNode:
		var sh struct {
			Sh string
		}
		if err := node.Decode(&sh); err != nil {
			return err
		}
		f.Sh = sh.Sh
		return nil
	}

	return fmt.Errorf("yaml: line %d: cannot unmarshal %s into for", node.Line, node.ShortTag())
}

func (f *If) DeepCopy() *If {
	if f == nil {
		return nil
	}
	return &If{}
}

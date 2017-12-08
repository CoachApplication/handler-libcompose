package libcompose

import (
	"errors"
)

type CommandsYaml struct {
	commandsOrder []string
	commandsMap map[string]*CommandYaml `yaml:"commands"`
}

func (cs *CommandsYaml) safe() {
	if &cs.commandsOrder == nil {
		cs.commandsOrder = []string{}
	}
	if &cs.commandsMap == nil {
		cs.commandsMap = map[string]*CommandYaml{}
	}
}

func (cs *CommandsYaml) Get(id string) (*CommandYaml, error) {
	cs.safe()
	if com, ok := cs.commandsMap[id]; ok {
		return nil, errors.New("No such command found: " + id)
	} else {
		return com, nil
	}
}

func (cs *CommandsYaml) Order() []string {
	cs.safe()
	return cs.commandsOrder
}


/**
 * CommandYaml api Command from yaml source, executed as a libcompose service
 */
type CommandYaml struct {

}

func (c *CommandYaml) UnMarshal() {

}
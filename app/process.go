package app

import (
	"fmt"
	"strings"
)

// Process the TaskStruct
func (task TaskStruct) Process(fileName string) error {
	switch task.Type {
	case "styles":
		return task.processStyles()
	case "scripts":
		return task.processScripts()
	case "copy":
		return task.processCopy(fileName)
	}

	return fmt.Errorf("Unknown process type: %s", task.Type)
}

// Return the relative file path
func rel(p string) string {
	if strings.HasPrefix(p, Conf.WorkingDir+"/") {
		return strings.TrimPrefix(p, Conf.WorkingDir+"/")
	}

	return p
}

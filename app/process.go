package app

import (
	"fmt"
	"strings"
)

// Process will process the ProcessStruct
func (p ProcessStruct) Process(fileName string) error {
	switch p.Type {
	case "styles":
		return p.processStyles()
	case "scripts":
		return p.processScripts()
	case "copy":
		return p.processCopy(fileName)
	}

	return fmt.Errorf("Unknown process type: %s", p.Type)
}

// Return the relative file path
func rel(p string) string {
	if strings.HasPrefix(p, Conf.WorkingDir+"/") {
		return strings.TrimPrefix(p, Conf.WorkingDir+"/")
	}

	return p
}

package app

import (
	"fmt"
)

var processTypes = map[string]bool{"styles": true, "scripts": true, "copy": true}

// Process will process the ProcessStruct
func (p ProcessStruct) Process() error {
	switch p.Type {
	case "styles":
		// return nil
		return p.processStyles()
	case "scripts":
		// return nil
		return p.processScripts()
	case "copy":
		return p.processCopy()
	}

	return fmt.Errorf("Unknown process type: %s", p.Type)
}

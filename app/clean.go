package app

import (
	"os"
	"path/filepath"

	"github.com/axllent/golp/utils"
	fg "github.com/goreleaser/fileglob"
)

// Clean deletes all matching files or directories defined in the yaml config
func Clean() error {
	for _, i := range Conf.CleanDirs {
		fullpth := filepath.ToSlash(i)
		if utils.IsDir(fullpth) {
			if err := os.RemoveAll(fullpth); err != nil {
				return err
			}
			Log().Debugf("deleted directory %s", rel(fullpth))
			continue
		}

		if utils.IsFile(fullpth) {
			if err := os.Remove(fullpth); err != nil {
				return err
			}
			Log().Debugf("deleted file %s", rel(fullpth))
			continue
		}
		matches, err := fg.Glob(fullpth, fg.MaybeRootFS)
		if err == nil {
			for _, f := range matches {
				if utils.IsFile(f) {
					if err := os.Remove(f); err != nil {
						return err
					}
					Log().Debugf("deleted file %s", rel(f))
				} else if utils.IsDir(f) {
					if err := os.RemoveAll(f); err != nil {
						return err
					}
					Log().Debugf("deleted directory %s", rel(f))
				}
			}
		}
	}

	return nil
}

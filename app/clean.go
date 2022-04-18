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
			Log().Debugf("Deleting directory %s", fullpth)
			if err := os.RemoveAll(fullpth); err != nil {
				return err
			}
			continue
		}

		if utils.IsFile(fullpth) {
			Log().Debugf("Deleting file %s", fullpth)
			if err := os.Remove(fullpth); err != nil {
				return err
			}
			continue
		}
		matches, err := fg.Glob(fullpth, fg.MaybeRootFS)
		if err == nil {
			for _, f := range matches {
				if utils.IsFile(f) {
					Log().Debugf("Deleting file %s", f)
					if err := os.Remove(f); err != nil {
						return err
					}
				} else if utils.IsDir(f) {
					Log().Debugf("Deleting directory %s", f)
					if err := os.RemoveAll(f); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

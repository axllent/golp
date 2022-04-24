package app

import (
	"os"
	"path"
	"path/filepath"

	"github.com/axllent/golp/utils"
)

// ProcessCopy will copy files from the source directory to the destination directory.
// If fileName is specified, just a single file will be copied.
func (task TaskStruct) processCopy(fileName string) error {
	sw := utils.StartTimer()

	files := task.Files()

	for _, f := range files {
		if fileName != "" && f.InFile != fileName {
			// not the same file, ignore
			continue
		}

		filename := filepath.Base(f.InFile)
		d := path.Join(task.Dist, f.OutPath)
		if !utils.IsDir(d) {
			/* #nosec G301 */
			if err := os.MkdirAll(d, 0755); err != nil {
				return err
			}
		}
		out := path.Join(d, filename)

		if err := utils.Copy(f.InFile, out); err != nil {
			return err
		}

		optimiseIfImage(task, out)

		srcStat, err := os.Stat(f.InFile)
		if err == nil {
			// get the original modification time for later
			mtime := srcStat.ModTime()
			atime := mtime // use mtime as we cannot get atime

			if err := os.Chtimes(out, atime, mtime); err != nil {
				Log().Debugf("error setting file timestamp: %v\n", err)
			}
		}

		Log().Debugf("copied %s to %s", rel(f.InFile), rel(out))
	}

	if fileName != "" {
		Log().Infof("'%s' updated in %v", task.Name, sw.Elapsed())
	} else {
		Log().Infof("'%s' copied in %v", task.Name, sw.Elapsed())
	}

	return nil
}

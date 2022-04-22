package app

import (
	"os"
	"path"
	"path/filepath"

	"github.com/axllent/golp/utils"
)

func (p ProcessStruct) processCopy() error {
	sw := utils.StartTimer()

	files := p.Files()

	for _, f := range files {
		filename := filepath.Base(f.InFile)
		d := path.Join(p.Dist, f.OutPath)
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

		optimiseIfImage(p, out)

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

	Log().Infof("'%s' copied in %v", p.Name, sw.Elapsed())

	return nil
}

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

		Log().Debugf("copied %s to %s", f.InFile, out)
	}

	Log().Infof("'%s' copied in %v", p.Name, sw.Elapsed())

	return nil
}

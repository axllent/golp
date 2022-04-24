package app

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/axllent/golp/utils"
)

func (task TaskStruct) processStyles() error {
	sw := utils.StartTimer()

	files := task.Files()

	if task.DistFile != "" {
		imports := []string{}
		for _, f := range files {
			extension := strings.ToLower(filepath.Ext(f.InFile))

			if extension == ".css" {
				c, err := utils.FileGetContents(f.InFile)
				if err != nil {
					return err
				}

				imports = append(imports, string(c))
			} else {
				imports = append(imports, fmt.Sprintf(`@import "%s";`, f.InFile))
			}
		}

		if !utils.IsDir(task.Dist) {
			/* #nosec G301 */
			if err := os.MkdirAll(task.Dist, 0755); err != nil {
				return err
			}
		}

		sassImport := strings.Join(imports, "\n")

		out := path.Join(task.Dist, task.DistFile)

		if err := compileStyles(sassImport, out, ""); err != nil {
			return err
		}

		Log().Debugf("processed %d SASS files to %s", len(files), rel(out))
		Log().Infof("'%s' compiled in %v", task.Name, sw.Elapsed())

		return nil
	}

	for _, f := range files {
		filename := filepath.Base(f.InFile)
		d := path.Join(task.Dist, f.OutPath)
		if !utils.IsDir(d) {
			/* #nosec G301 */
			if err := os.MkdirAll(d, 0755); err != nil {
				return err
			}
		}

		out := path.Join(d, filename)
		extension := strings.ToLower(filepath.Ext(filename))

		if extension == ".scss" || extension == ".sass" || extension == ".css" {
			out = out[0:len(out)-len(extension)] + ".css"

			content := fmt.Sprintf(`@import "%s";`, f.InFile)

			if extension == ".css" {
				c, err := utils.FileGetContents(f.InFile)
				if err != nil {
					return err
				}

				content = c
			}

			if err := compileStyles(string(content), out, f.InFile); err != nil {
				return err
			}

			Log().Debugf("compiled %s to %s", rel(f.InFile), rel(out))

		} else {
			Log().Warningf("unsupported stylesheet file extension: %s", f)
		}
	}

	Log().Infof("'%s' compiled in %v", task.Name, sw.Elapsed())

	return nil
}

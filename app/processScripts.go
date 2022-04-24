package app

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/axllent/golp/utils"
	"github.com/evanw/esbuild/pkg/api"
)

func (task TaskStruct) processScripts() error {
	sw := utils.StartTimer()

	files := task.Files()

	if task.DistFile != "" {
		imports := []string{}
		for _, f := range files {
			imports = append(imports, f.InFile)
		}

		if !utils.IsDir(task.Dist) {
			/* #nosec G301 */
			if err := os.MkdirAll(task.Dist, 0755); err != nil {
				return err
			}
		}

		out := path.Join(task.Dist, task.DistFile)

		options := api.BuildOptions{
			Stdin: &api.StdinOptions{
				Contents: "",
			},
			Inject:         imports,
			Outfile:        out,
			Write:          true,
			AllowOverwrite: true,
			Format:         api.FormatCommonJS,
			SourcesContent: api.SourcesContentExclude,
		}

		if task.JSBundle {
			options.Bundle = true
		}

		if Minify {
			options.MinifyWhitespace = true
			options.MinifyIdentifiers = true
			options.MinifySyntax = true
		} else {
			options.Sourcemap = api.SourceMapLinked
		}

		result := api.Build(options)

		if len(result.Errors) > 0 {
			errorMsg := fmt.Sprintf("> Error %s:%d\n%s",
				result.Errors[0].Location.File,
				result.Errors[0].Location.Line,
				result.Errors[0].Text,
			)

			return fmt.Errorf("%s", errorMsg)
		}

		Log().Debugf("compiled %d JS files to %s", len(files), rel(out))
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

		options := api.BuildOptions{
			EntryPoints:    []string{f.InFile},
			Outfile:        out,
			Write:          true,
			AllowOverwrite: true,
			SourcesContent: api.SourcesContentExclude,
		}

		if task.JSBundle {
			options.Bundle = true
		}

		if Minify {
			options.MinifyWhitespace = true
			options.MinifyIdentifiers = true
			options.MinifySyntax = true
		} else {
			options.Sourcemap = api.SourceMapLinked
		}

		result := api.Build(options)

		if len(result.Errors) > 0 {
			errorMsg := fmt.Sprintf("> Error %s:%d\n%s",
				result.Errors[0].Location.File,
				result.Errors[0].Location.Line,
				result.Errors[0].Text,
			)

			return fmt.Errorf("%s", errorMsg)
		}

		Log().Debugf("compiled %s to %s", rel(f.InFile), rel(out))
	}

	Log().Infof("'%s' compiled in %v", task.Name, sw.Elapsed())

	return nil
}

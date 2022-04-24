package app

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/axllent/golp/utils"
	"gopkg.in/yaml.v3"

	fg "github.com/goreleaser/fileglob"
)

var (
	// Conf struct
	Conf struct {
		ConfigFile string   // build tasks are relative to this config
		WorkingDir string   // working directory is the base directory of the config file
		CleanDirs  []string // is set, this directory will be deleted with clean
		Tasks      []TaskStruct
	}

	// Minify determines whether to minify the styles and scripts
	Minify bool
)

// ParseConfig reads a yaml file and returns a Conf struct
func ParseConfig() error {
	var yml = yamlConf{}
	Conf.Tasks = []TaskStruct{}

	if !utils.IsFile(Conf.ConfigFile) {
		return fmt.Errorf("Config %s does not exist", Conf.ConfigFile)
	}

	buf, err := ioutil.ReadFile(Conf.ConfigFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, &yml)
	if err != nil {
		return err
	}

	Conf.WorkingDir, err = filepath.Abs(filepath.Dir(Conf.ConfigFile))
	if err != nil {
		return err
	}

	for _, c := range yml.Clean {
		Conf.CleanDirs = append(Conf.CleanDirs, filepath.Join(Conf.WorkingDir, c))
	}

	for _, task := range yml.Styles {
		c := TaskStruct{}
		c.Type = "styles"
		c.Name = task.Name
		if c.Name == "" {
			c.Name = c.Type
		}

		c.Src = task.Src

		if task.Dist == "" {
			return fmt.Errorf("Dist is not set for %s", c.Name)
		}

		if strings.HasSuffix(task.Dist, ".css") {
			c.DistFile = filepath.Base(task.Dist)
			task.Dist = filepath.Dir(task.Dist)
		}

		c.Dist = filepath.Join(Conf.WorkingDir, task.Dist)

		if len(c.Src) > 0 {
			Conf.Tasks = append(Conf.Tasks, c)
		}
	}

	for _, task := range yml.Scripts {
		c := TaskStruct{}
		c.Type = "scripts"
		c.Name = task.Name
		if c.Name == "" {
			c.Name = c.Type
		}

		c.Src = task.Src

		if task.Dist == "" {
			return fmt.Errorf("Dist is not set for %s", c.Name)
		}

		if strings.HasSuffix(task.Dist, ".js") {
			c.DistFile = filepath.Base(task.Dist)
			task.Dist = filepath.Dir(task.Dist)
		}

		c.Dist = filepath.Join(Conf.WorkingDir, task.Dist)
		c.JSBundle = task.Bundle

		if len(c.Src) > 0 {
			Conf.Tasks = append(Conf.Tasks, c)
		}
	}

	for _, task := range yml.Copy {
		c := TaskStruct{}
		c.Type = "copy"
		c.Name = task.Name
		if c.Name == "" {
			c.Name = c.Type
		}

		c.Src = task.Src

		if task.Dist == "" {
			return fmt.Errorf("Dist is not set for %s", c.Name)
		}

		c.Dist = filepath.Join(Conf.WorkingDir, task.Dist)

		c.OptimiseImages = task.OptimiseImages

		c.SVGPrecision = task.SVGPrecision
		if c.SVGPrecision < 1 || task.SVGPrecision > 25 {
			c.SVGPrecision = 5
		}

		if len(c.Src) > 0 {
			Conf.Tasks = append(Conf.Tasks, c)
		}
	}

	if len(Conf.Tasks) == 0 {
		return fmt.Errorf("No tasks defined")
	}

	initOptimiserConfig()

	return nil
}

// Files returns all files matching the glob pattern
func (t TaskStruct) Files() []FileMap {

	fm := []FileMap{}
	exists := map[string]bool{}

	for _, pth := range t.Src {
		fullpth := filepath.ToSlash(filepath.Join(Conf.WorkingDir, pth))
		Log().Debugf("finding files in %s", rel(fullpth))
		matches, err := fg.Glob(fullpth, fg.MaybeRootFS)
		if err == nil {
			subDirFrom := ""

			if strings.Contains(fullpth, "*") {
				parts := strings.Split(fullpth, "*")
				subDirFrom = parts[0]
			}

			for _, f := range matches {
				if utils.IsFile(f) {
					subDir := ""
					// only add each file once
					if _, ok := exists[f]; ok {
						continue
					}

					if subDirFrom != "" {
						if strings.HasPrefix(f, subDirFrom) {
							if len(filepath.Dir(f)) > len(subDirFrom) {
								subDir = filepath.Dir(f)[len(subDirFrom):]
							}
						}
					}

					fm = append(
						fm,
						FileMap{
							InFile:  filepath.ToSlash(f),
							OutPath: filepath.ToSlash(subDir),
						},
					)

					exists[f] = true
				}
			}
		} else {
			Log().Error(err)
		}
	}

	return fm
}

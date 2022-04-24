package app

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/axllent/golp/utils"
)

// TestGeoIPLookup uses test data from
// https://github.com/maxmind/MaxMind-DB/tree/master/test-data

var (
	cssFiles     = []string{"bootstrap.css", "test.css"}
	jsFiles      = []string{"bootstrap.js", "merged.js"}
	copyFiles    = []string{"placeholder1.png", "placeholder2.png", "subdir/placeholder3.png"}
	skippedFiles = []string{"subdir/ignore.txt"}
)

func TestProcessConfig(t *testing.T) {
	QuietLogging = true
	Conf.ConfigFile = filepath.Join("..", "test-data", "golp.yaml")

	if err := ParseConfig(); err != nil {
		t.Error(err)
	}
}

func TestProcessClean(t *testing.T) {
	QuietLogging = true
	if err := Clean(); err != nil {
		t.Error(err)
	}
}
func TestTaskProcess(t *testing.T) {
	QuietLogging = true
	for _, task := range Conf.Tasks {
		if err := task.Process(""); err != nil {
			t.Error(err)
		}
	}
}

func TestProcessSourceMaps(t *testing.T) {
	QuietLogging = true
	cssDir := filepath.Join(Conf.WorkingDir, "dist", "css")

	for _, c := range cssFiles {
		// test CSS files
		f := filepath.Join(cssDir, c)
		if !utils.IsFile(f) {
			t.Errorf("%s not found", f)
		}

		// test sourcemaps exist
		if !utils.IsFile(f + ".map") {
			t.Errorf("%s not found", f+".map")
		}

		srcMapURL := fmt.Sprintf("/*# sourceMappingURL=%s.map */", c)

		hasSM, err := utils.FileContains(f, srcMapURL)
		if err != nil {
			t.Error(err)
		}
		if !hasSM {
			t.Errorf("%s does not contain %s", f, srcMapURL)
		}
	}

	jsDir := filepath.Join(Conf.WorkingDir, "dist", "js")

	for _, c := range jsFiles {
		// test JS files
		f := filepath.Join(jsDir, c)
		if !utils.IsFile(f) {
			t.Errorf("%s not found", f)
		}

		// test sourcemaps exist
		if !utils.IsFile(f + ".map") {
			t.Errorf("%s not found", f+".map")
		}

		srcMapURL := fmt.Sprintf("//# sourceMappingURL=%s.map", c)

		hasSM, err := utils.FileContains(f, srcMapURL)
		if err != nil {
			t.Error(err)
		}
		if !hasSM {
			t.Errorf("%s does not contain %s", f, srcMapURL)
		}
	}

	if err := Clean(); err != nil {
		t.Error(err)
	}
}

func TestProcessCompressed(t *testing.T) {
	QuietLogging = true
	Minify = true

	for _, task := range Conf.Tasks {
		if err := task.Process(""); err != nil {
			t.Error(err)
		}
	}
}

func TestProcessCompressedFiles(t *testing.T) {
	QuietLogging = true
	cssDir := filepath.Join(Conf.WorkingDir, "dist", "css")

	for _, c := range cssFiles {
		f := filepath.Join(cssDir, c)
		if !utils.IsFile(f) {
			t.Errorf("%s not found", f)
		}

		// test maps exist
		if utils.IsFile(f + ".map") {
			t.Errorf("%s should not exist", f+".map")
		}

		srcMapURL := fmt.Sprintf("/*# sourceMappingURL=%s.map */", c)

		hasSM, err := utils.FileContains(f, srcMapURL)
		if err != nil {
			t.Error(err)
		}
		if hasSM {
			t.Errorf("%s should not contain %s", f, srcMapURL)
		}
	}

	jsDir := filepath.Join(Conf.WorkingDir, "dist", "js")

	for _, c := range jsFiles {
		// test files
		f := filepath.Join(jsDir, c)
		if !utils.IsFile(f) {
			t.Errorf("%s not found", f)
		}

		// test maps exist
		if utils.IsFile(f + ".map") {
			t.Errorf("%s should not exist", f+".map")
		}

		srcMapURL := fmt.Sprintf("//# sourceMappingURL=%s.map", c)

		hasSM, err := utils.FileContains(f, srcMapURL)
		if err != nil {
			t.Error(err)
		}
		if hasSM {
			t.Errorf("%s should not contain %s", f, srcMapURL)
		}
	}
}

func TestProcessCopyFiles(t *testing.T) {
	QuietLogging = true
	copyDir := filepath.Join(Conf.WorkingDir, "dist", "images")

	for _, c := range copyFiles {
		f := filepath.Join(copyDir, c)
		if !utils.IsFile(f) {
			t.Errorf("copied file %s missing", f)
		}
	}
	for _, c := range skippedFiles {
		f := filepath.Join(copyDir, c)
		if utils.IsFile(f) {
			t.Errorf("skipped file %s exists", f)
		}
	}
}

func TestProcessCleanFinal(t *testing.T) {
	QuietLogging = true
	if err := Clean(); err != nil {
		t.Error(err)
	}
}

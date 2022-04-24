package app

import (
	"bytes"
	"fmt"
	"image/gif"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/svg"
)

// OptimiserConfig is the configuration for the optimiser
var OptimiserConfig struct {
	optimise  bool
	jpegtran  string
	jpegoptim string
	optipng   string
	pngquant  string
	gifsicle  string
}

// InitOptimiserConfig will set the optimiser config
func initOptimiserConfig() {
	OptimiserConfig.optimise = true
	OptimiserConfig.jpegtran = which("jpegtran")
	if OptimiserConfig.jpegtran == "" {
		OptimiserConfig.jpegoptim = which("jpegoptim")
	}
	OptimiserConfig.optipng = which("optipng")
	OptimiserConfig.pngquant = which("pngquant")
	OptimiserConfig.gifsicle = which("gifsicle")
}

// Which will try locate a binary in the $PATH
func which(bin string) string {
	exe, err := exec.LookPath(bin)
	if err == nil {
		return exe
	}

	return ""
}

// OptimiseImage will optimise an image if it matches a specific
// extension
func optimiseIfImage(p TaskStruct, imgPath string) {
	if !p.OptimiseImages {
		return
	}
	ext := strings.ToLower(filepath.Ext(imgPath))
	if OptimiserConfig.jpegoptim != "" && (ext == ".jpg" || ext == ".jpeg") {
		runOptimiser(OptimiserConfig.jpegtran, "-optimize", imgPath)
	}
	if OptimiserConfig.jpegoptim != "" && (ext == ".jpg" || ext == ".jpeg") {
		runOptimiser(OptimiserConfig.jpegoptim, "-f", "-s", "-o", imgPath)
	}
	if OptimiserConfig.pngquant != "" && ext == ".png" {
		runOptimiser(OptimiserConfig.pngquant, "-f", "-ext", ".png", imgPath)
	}
	if OptimiserConfig.optipng != "" && ext == ".png" {
		runOptimiser(OptimiserConfig.optipng, imgPath)
	}
	if OptimiserConfig.gifsicle != "" && ext == ".gif" {
		if isGIFAnimated(imgPath) == nil {
			runOptimiser(OptimiserConfig.gifsicle, imgPath, "-o", imgPath)
		} else {
			Log().Debugf("skipping optimisation of animated gif %s", rel(imgPath))
		}
	}

	if ext == ".svg" {
		optimiseSVG(p, imgPath)
	}
}

// RunOptimiser will run the specified command on a copy of the temporary file,
// and overwrite it if the output is smaller than the original
func runOptimiser(bin string, args ...string) {
	cmd := exec.Command(bin, args...)
	Log().Debugf("optimising %s %s", bin, strings.Join(args, " "))
	if err := cmd.Run(); err != nil {
		Log().Errorf("%s: %v\n", args[0], err)
	}
}

// IsGIFAnimated will return an error if the GIF file has more than 1 frame
func isGIFAnimated(gifFile string) error {
	file, _ := os.Open(filepath.Clean(gifFile))
	/* #nosec G307 */
	defer file.Close()

	g, err := gif.DecodeAll(file)
	if err != nil {
		return err
	}

	// Single frame = OK
	if len(g.Image) == 1 {
		return nil
	}

	return fmt.Errorf("Animated gif")
}

func optimiseSVG(p TaskStruct, imgPath string) {
	min := minify.New()
	m := svg.Minifier{Precision: p.SVGPrecision}

	file, err := os.Open(filepath.Clean(imgPath))
	if err != nil {
		Log().Error(err)
		return
	}

	/* #nosec G307 */
	defer file.Close()

	buf := &bytes.Buffer{}

	if err := m.Minify(min, buf, file, nil); err != nil {
		Log().Error(err)
		return
	}

	/* #nosec G104 */
	file.Close()

	f, err := os.Create(filepath.Clean(imgPath))
	if err != nil {
		Log().Errorf(err.Error())
		return
	}

	/* #nosec G307 */
	defer f.Close()

	Log().Debugf("optimising %s", rel(imgPath))

	if _, err := buf.WriteTo(f); err != nil {
		Log().Errorf(err.Error())
	}
}

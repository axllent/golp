package app

import (
	"fmt"
	"image/gif"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
func initOptimiserConfig(optimise bool) {
	if !optimise {
		return
	}

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
func optimiseIfImage(imgPath string) {
	if !OptimiserConfig.optimise {
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
	if OptimiserConfig.gifsicle != "" && ext == ".gif" && isGIFAnimated(imgPath) == nil {
		runOptimiser(OptimiserConfig.gifsicle)
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

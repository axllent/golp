package app

// TaskStruct for config
type TaskStruct struct {
	Type           string   // styles, scripts, copy
	Name           string   // name of the task
	Src            []string // source files
	Dist           string   // dist directory
	DistFile       string   // optional merge filename
	NoSourceMaps   bool     // optional no SourceMaps (SAS/JS)
	JSBundle       bool     // optional bundle JS files (JS)
	OptimiseImages bool     `yaml:"optimise_images"`
	SVGPrecision   int      `yaml:"svg_precision"`
}

// YamlConf is the yaml struct
type yamlConf struct {
	Clean  []string `yaml:"clean"`
	Styles []struct {
		Name         string   `yaml:"name"`
		Src          []string `yaml:"src"`
		Dist         string   `yaml:"dist"`
		NoSourceMaps bool     `yaml:"no_sourcemaps"`
	} `yaml:"styles"`
	Scripts []struct {
		Name         string   `yaml:"name"`
		Src          []string `yaml:"src"`
		Dist         string   `yaml:"dist"`
		Bundle       bool     `yaml:"bundle"`
		NoSourceMaps bool     `yaml:"no_sourcemaps"`
	} `yaml:"scripts"`
	Copy []struct {
		Name           string   `yaml:"name"`
		Src            []string `yaml:"src"`
		Dist           string   `yaml:"dist"`
		OptimiseImages bool     `yaml:"optimise_images"`
		SVGPrecision   int      `yaml:"svg_precision"`
	} `yaml:"copy"`
}

type watchMap struct {
	Path       string
	TaskStruct TaskStruct
}

// FileMap struct maps the file to the respective dist directory
type FileMap struct {
	InFile  string
	OutPath string
}

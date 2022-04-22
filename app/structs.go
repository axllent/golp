package app

// ProcessStruct for config
type ProcessStruct struct {
	Type           string   // styles, scripts, copy
	Name           string   // name of the process
	Src            []string // source files
	Dist           string   // dist directory
	DistFile       string   // optional merge filename
	JSBundle       bool     // optional bundle JS files (JS-only)
	OptimiseImages bool     `yaml:"optimise_images"`
	SVGPrecision   int      `yaml:"svg_precision"`
}

// YamlConf is the yaml struct
type yamlConf struct {
	Clean  []string `yaml:"clean"`
	Styles []struct {
		Name string   `yaml:"name"`
		Src  []string `yaml:"src"`
		Dist string   `yaml:"dist"`
	} `yaml:"styles"`
	Scripts []struct {
		Name   string   `yaml:"name"`
		Src    []string `yaml:"src"`
		Dist   string   `yaml:"dist"`
		Bundle bool     `yaml:"bundle"`
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
	Path          string
	ProcessStruct ProcessStruct
}

// FileMap struct maps the file to the respective dist directory
type FileMap struct {
	InFile  string
	OutPath string
}

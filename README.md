# Golp - automate build workflows

Golp automates your website build workflows, compiling SASS and JavaScript into configurable "dist" directories, and copying of static assets with optional image optimisation.

Golp is not a Gulp drop-in replacement, but aims to provide an easy-to-use alternative for typical websites using SASS & JavaScript. It is fast, simple, and runs from a single binary.

Internally it uses [esbuild](https://github.com/evanw/esbuild) for SASS/CSS compilation, and [golibsass](https://github.com/bep/golibsass) for JavaScript compilation - both supporting SourceMaps or minification.


## Features

- Single binary for all build tasks with simple yaml file configuration
- Ability to "watch" configured files for changes (ie: building during development)
- SourceMaps for debugging SASS & JS (disabled with minification)
- Process/compile SASS & CSS ([golibsass](https://github.com/bep/golibsass)), and JavaScript ([esbuild](https://github.com/evanw/esbuild))
- Copy static assets, including optional image optimisation (jpg, png, gif, svg)


## Motivation

Having used [Gulp](https://gulpjs.com/) for several years to build and package website assets (compiling SASS, JavaScript and copy some src assets such as images and fonts to a dist folder), I wanted to reduce the build build overheads and reliance on a whole node ecosystem just to process website assets. The development of Gulp has practically stopped, which has lead to much frustration due to outdated package requirements, many of which have known vulnerabilities (CVE issues).

Using just a single binary (Golp), I was able to reduce the build/compile time by ~50% (not even including the additional `npm install` time), as well as drastically reduce the number of node packages (required for gulp) by about 1,250 packages (9,062 files).


## Usage

Golp requires a `golp.yaml` file in your project root (see [configuration](#configuration))

```
Usage:
  golp [command]

Available Commands:
  build       Compile & copy your assets (single)
  clean       Clean (delete) your configured files/directories
  config      View an example config file
  version     Display the current version & update information
  watch       Build & watch src directories for changes

Flags:
  -c  --config string   config file (default "./golp.yaml")
  -h, --help            help for golp
  -q, --quiet           no output except for errors
  -v, --verbose         verbose logging
```


### Usage examples

```
golp build
```
This will process your files, outputting them to their respective dist directories. JavaScript and SASS files will include a SourceMap for debugging.

```
golp build -m
```
This will process your files, outputting them to their respective dist directories. JavaScript and SASS files will be minified (compressed).

```
golp watch
```
This will process your files, outputting them to their respective dist directories. Golp will then continue to watch those source directories for changes, and rebuild/recompile as necessary.


## Installation


Linux & Mac users can install it directly to `/usr/local/bin/golp` with:
```
sudo bash < <(curl -sL https://raw.githubusercontent.com/axllent/golp/develop/install.sh)
```

There are pre-built binaries available for Linux, Windows and MacOS available in the [releases](https://github.com/axllent/golp/releases/latest).

You can also install it from source: `go install github.com/axllent/golp@latest` (go, gcc & g++ required).

See [Installation](https://github.com/axllent/golp/wiki/Installation) for more information.


## Configuration

Typically your configuration file will be found in your project root folder, and named `golp.yaml`. An alternative config can be specified using the `-c` flag.

Golp has four types of task types: `clean`, `styles`, `scripts` and `copy`. Please see the [wiki](https://github.com/axllent/golp/wiki) for a full list of configuration options.

Run `golp config` to view an example config file.

See [Configuration](https://github.com/axllent/golp/wiki/Configuration) for more information.


### Example config file

Golp is really easy to configure in your `golp.yaml` file. Please refer to the [wiki](https://github.com/axllent/golp/wiki/Configuration) for more explanation.

```yaml
clean: 
  - themes/site/dist

styles:
  - src:
      - themes/site/src/sass/*.scss
      - themes/site/src/css/**.css 
      - node_modules/@dashboardcode/bsmultiselect/dist/css/BsMultiSelect.css
    dist: themes/site/dist/css/

scripts:
  - src:
      - node_modules/@popperjs/core/dist/umd/popper.min.js
      - node_modules/bootstrap/dist/js/bootstrap.min.js
      - node_modules/axios/dist/axios.min.js
      - themes/site/src/js/**.js
    dist: themes/site/dist/js/libs.js 

copy:
  - src:
      - themes/site/src/images/**
    dist: themes/site/dist/images
    optimise_images: true
    svg_precision: 5
  - src:
      - themes/site/src/fonts/**
    dist: themes/site/dist/fonts  
```

# Golp

**This app is very much in beta at the moment!**

Golp automates build workflows, compiling SASS and JavaScript into configurable "dist" directories, and copying of static assets.

Golp is not a Gulp drop-in replacement, but aims to provide an easy-to-use build tool for typical websites using SASS & JavaScript. It is fast, simple, and runs from a single binary.

Internally it uses [esbuild](https://github.com/evanw/esbuild) for SASS/CSS, and [golibsass](https://github.com/bep/golibsass) for JavaScript compilation.


## Features

- Single binary for all build tasks with simple yaml file configuration
- Ability to "watch" configured files for changes (ie: building during development)
- SourceMaps for debugging SASS & JS (disabled with minification)
- Process/compile SASS & CSS ([golibsass](https://github.com/bep/golibsass)), and JavaScript ([esbuild](https://github.com/evanw/esbuild))


## Motivation

Having used [Gulp](https://gulpjs.com/) for several years to build and package website assets (compiling SASS, JavaScript and copy some src assets such as images and fonts to a dist folder), I wanted to reduce the build build overheads and reliance on node packages. The development of Gulp has practically stopped, which has lead to much frustration due to outdated package requirements, many of which have known CVE issues.

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

There are pre-built binaries available for Linux, Windows and MacOS available in the [releases](https://github.com/axllent/golp/releases/latest).

You can also install it from source: `go install github.com/axllent/golp@latest` (go, gcc & g++ required)


## Configuration

Typically your configuration file will be found in your project root folder, and named `golp.yaml`. An alternative config can be specified using the `-c` flag.

Please note that all `styles`, `scripts` and `copy` source files are relative to your config file.

Run `golp config` to view an example config file.


### Example config file

```yaml
## Optionally specify directories and/or files to automatically delete on every build, 
## and on `golp clean`. Files can be matched with wildcards.
clean: 
  ## optional directories to delete
  - themes/site/dist

## SASS & CSS files
styles:
  - src:
      ## process all *.scss files in this folder
      - themes/site/src/sass/*.scss
      ## process all *.css files in this folder and child folders
      - themes/site/src/sass/**.css 
      ## add a specific file
      - node_modules/@dashboardcode/bsmultiselect/dist/css/BsMultiSelect.css
    ## output directory for all src files
    dist: themes/site/dist/css/

## JS scripts are processed with esbuild, and can be optionally merged into a single file,
## and optionally bundled (see https://esbuild.github.io/api/#bundle)
scripts:
  ## compile and merge these into a single libs.js file
  - src:
      - node_modules/@popperjs/core/dist/umd/popper.min.js
      - node_modules/bootstrap/dist/js/bootstrap.min.js
      - node_modules/axios/dist/axios.min.js
      - node_modules/@dashboardcode/bsmultiselect/dist/js/BsMultiSelect.min.js
      - node_modules/vuedraggable/dist/vuedraggable.umd.min.js
      - node_modules/sortablejs/Sortable.min.js
      - node_modules/vue/dist/vue.global.prod.js
    ## merge all files into a single file (only scripts & styles supported)
    dist: themes/site/dist/js/libs.js 
    ## optional name for the console output
    name: libs
    ## optionally bundle your JavaScript, @see https://esbuild.github.io/api/#bundle 
    # bundle: true

  ## compile all *.js files in this folder and child folders
  - src:
      - themes/site/src/js/**.js
    dist: themes/site/dist/js
    name: site scripts

## Copy all matching files from the src directories to the dist directories.
## This does not support merging or compressing of files do dist should be a directory name.
copy:
  - src:
      - themes/site/src/images/**
    dist: themes/site/dist/images
    name: images

  - src: 
      - themes/site/src/fonts/**
    dist: themes/site/dist/fonts/
    name: fonts
```

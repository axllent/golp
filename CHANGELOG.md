# Changelog

Notable changes to golp will be documented in this file.


## 0.1.35

### Libs
- Update modules


## 0.1.34

### Libs
- Update modules


## 0.1.33

### Feature
- Add no_sourcemaps option for JS & SASS


## 0.1.32

### Fix
- Watch directories should be relative to golp.yaml


## 0.1.31

### Docs
- Update documentation

### Feature
- Add installer script


## 0.1.30

### Chore
- Apply only relevant args to functions


## 0.1.29

### Feature
- Binary check on update command

### Fix
- Ensure extracted binary is executable (linux/darwin)


## 0.1.27

### Docs
- Update cli documentation


## 0.1.25

### Feature
- Copy while watching will only process modified files

### Fix
- gifsicle optimisation


## 0.1.24

### Feature
- Optional SVG optimisation on copy


## 0.1.23

### Debug
- Use relative paths in debug output

### Feature
- Optional image optimisation on copy
- Add --quiet option for build & watch
- Preserve timestamps of copied files

### Fix
- Fix subdirectory patterns in recursive copy


## 0.1.22

### Fix
- Close file handle after copying contents


## 0.1.21

### Feature
- Add tests and test data


## 0.1.20

### Feature
- Display compiled module versions

### Libs
- Update modules


## 0.1.19

### Config
- Dist file/folder deletion based on config file only


## 0.1.17

### Feature
- Add `golp package` alias to imply build minification
- Add changelog to release info


## 0.1.16

### Feature
- Add CHANGELOG


## 0.1.15

### Chore
- Version info to include build platform / arch
- Add tag step to build scripts
- Static Linux builds
- Switch to github.com/bep/golibsass
- Add javaScript sourcemaps for debugging

### Fix
- Windows paths
- Escape glob paths for Windows paths


## 0.1.12

### Fix
- Windows paths


## 0.1.10

### Fix
- **windows:** Escape glob paths


## 0.1.9

### Fix
- don't rebuild Windows binaries



package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/golang/dep"
	"github.com/pkg/errors"
)

func main() {
	err := do()
	if err != nil {
		panic(err)
	}
}

func do() (err error) {
	Gopath := os.Getenv("GOPATH")
	fullPath, err := filepath.Abs(os.Args[0])
	if err != nil {
		return errors.Wrap(err, "failed to make cwd absolute")
	}
	cwd := filepath.Dir(fullPath)

	ctx := &dep.Ctx{
		GOPATH:         Gopath,
		Out:            log.New(os.Stdout, "", 0),
		Err:            log.New(os.Stderr, "", 0),
		Verbose:        true,
		DisableLocking: os.Getenv("DEPNOLOCK") != "",
	}

	GOPATHS := filepath.SplitList(os.Getenv("GOPATH"))
	ctx.SetPaths(cwd, GOPATHS...)

	p, err := ctx.LoadProject()
	if err != nil {
		return errors.Wrap(err, "failed to load project")
	}

	t, err := p.ParseRootPackageTree()
	if err != nil {
		return errors.Wrap(err, "failed to parse package tree")
	}

	pkgs := make(map[string]struct{})

	for _, pkg := range t.Packages {
		if pkg.Err != nil {
			continue
		}

		for _, i := range pkg.P.Imports {
			components := strings.Split(i, "/")
			// skip import paths without at least 2 levels (github.com/1/2 is valid, but encoding/json is not)
			if len(components) < 3 {
				continue
			}
			// skip paths that aren't remote
			if !strings.ContainsAny(components[0], ".") {
				continue
			}
			// build the root import path, assuming domain.user.repo format
			rootImportPath := fmt.Sprintf("%s/%s/%s", components[0], components[1], components[2])
			// skip self
			if strings.Contains(cwd, rootImportPath) {
				continue
			}
			pkgs[rootImportPath] = struct{}{}
		}
	}

	sorted := []string{}
	for pkg := range pkgs {
		sorted = append(sorted, pkg)
	}
	sort.Strings(sorted)

	// We now have a sorted list of repositories we depend on!
	// Go searching for readmes etc

	vendor := filepath.Join(cwd, "vendor")
	if !Exists(vendor) {
		return errors.New("No vendor directory found, cannot search packages :(")
	}

	if len(sorted) == 0 {
		fmt.Println("No dependencies found.")
		return
	}

	fmt.Println("You depend on:")

	for _, pkg := range sorted {
		// pkgPath := filepath.Join(vendor, pkg)
		// if !Exists(pkgPath) {
		// 	fmt.Println("No local copy of", pkg, "found in vendor/")
		// 	continue
		// }

		// At this point I realised the Go community doesn't really do donations...
		// But at least

		fmt.Println("-", fmt.Sprintf("https://%s", pkg))
	}

	fmt.Println("Go buy em a beer!")

	return
}

// Exists simply checks if a path exists and panics on error
func Exists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		panic(err)
	}
	return true
}

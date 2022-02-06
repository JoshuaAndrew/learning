package main

import (
	"fmt"
	"os"
	"path/filepath"
)



// DefaultGlideFile is the default name for the glide.yaml file.
const DefaultGlideFile = "glide.yaml"

// GlideFile is the name of the Glide file.
//
// Setting this is not concurrency safe. For consistency, it should really
// only be set once, at startup, or not at all.
var GlideFile = DefaultGlideFile

// VendorDir is the name of the directory that holds vendored dependencies.
//
// As of Go 1.5, this is always vendor.
var VendorDir = "vendor"


// GlideWD finds the working directory of the glide.yaml file, starting at dir.
//
// If the glide file is not found in the current directory, it recurses up
// a directory.
func GlideWD(dir string) (string, error) {
	fullpath := filepath.Join(dir, GlideFile)

	if _, err := os.Stat(fullpath); err == nil {
		return dir, nil
	}

	base := filepath.Dir(dir)
	if base == dir {
		return "", fmt.Errorf("Cannot resolve parent of %s", base)
	}

	return GlideWD(base)
}


// IsLink returns true if the given FileInfo references a link.
func IsLink(fi os.FileInfo) bool {
	return fi.Mode()&os.ModeSymlink == os.ModeSymlink
}

// Vendor calculates the path to the vendor directory.
//
// Based on working directory, VendorDir and GlideFile, this attempts to
// guess the location of the vendor directory.
func Vendor() (string, error) {
	cwd, err := os.Getwd()
	fmt.Printf("GetWd:%s",cwd)
	if err != nil {
		return "", err
	}

	// Find the directory that contains glide.yaml
	yamldir, err := GlideWD(cwd)
	if err != nil {
		return cwd, err
	}

	gopath := filepath.Join(yamldir, VendorDir)

	// Resolve symlinks
	info, err := os.Lstat(gopath)
	if err != nil {
		return gopath, nil
	}
	for i := 0; IsLink(info) && i < 255; i++ {
		p, err := os.Readlink(gopath)
		if err != nil {
			return gopath, nil
		}
		if filepath.IsAbs(p) {
			gopath = p
		} else {
			gopath = filepath.Join(filepath.Dir(gopath), p)
		}
		info, err = os.Lstat(gopath)
		if err != nil {
			return gopath, nil
		}
	}

	return gopath, nil
}



func getFilelist() {
	searchPath, _ := Vendor()
	if _, err := os.Stat(searchPath); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Vendor directory does not exist.")
		}
	}

	err := filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
		if path == searchPath {
			return err
		}
		name := info.Name()
		if name == "vendor" {
			if _,err := os.Stat(path); err == nil{
				if info.IsDir() {
					fmt.Printf("Name: %s, Removing: %s \n",name,path)
					return nil
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func main() {
	//root := "D:\\tenxcloud\\workspace\\mybeego\\src\\github.com\\JoshuaAndrew\\mybeego"
	getFilelist()
}

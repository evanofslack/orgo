package orgo

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

// createDir creates a directory with with input name
func CreateDir(name string) error {
	path := filepath.Join(".", name)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// findFiles returns all paths of files matching extensions in top level dir
func FindFiles(fsys fs.FS, extensions []string) []string {
	files, err := fs.ReadDir(fsys, ".")
	if err != nil {
		log.Fatal(err)
	}
	var outFiles []string
	for _, f := range files {
		for _, e := range extensions {
			if filepath.Ext(f.Name()) == e {
				outFiles = append(outFiles, f.Name())
			}
		}
	}
	return outFiles
}

// findFilesRecursive walks a provided file system tree and returns all paths of files matching extensions
func findFilesRecursive(fsys fs.FS, extensions []string) []string {
	var files []string
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		for _, e := range extensions {
			if filepath.Ext(p) == e {
				files = append(files, p)
			}
		}
		return nil
	})
	return files
}

// moveFiles moves all input files to provided directory
func MoveFiles(inFiles []string, outDir string) error {
	for _, f := range inFiles {
		path := outDir + "/" + f
		err := os.Rename(f, path)
		if err != nil {
			return err
		}
	}
	return nil
}

package main

import (
	"os"
	"testing"
)

func TestCreateDir(t *testing.T) {

	const newDirName = "test_directory"

	// ensure directory to be created doesn't exist
	cwd, err := os.Getwd()
	if err != nil {
		t.Errorf("unable to get cwd")
	}
	dirs, err := os.ReadDir(cwd)
	if err != nil {
		t.Errorf("unable to read dirs")
	}
	for _, dir := range dirs {
		if dir.Name() == newDirName {
			t.Errorf("directory to be created already exists")
		}
	}

	// test createDir function
	err = createDir(newDirName)
	if err != nil {
		t.Errorf("error creating directory")
	}

	// ensure directory was created
	dirs, err = os.ReadDir(cwd)
	if err != nil {
		t.Errorf("unable to read dirs")
	}
	for _, dir := range dirs {
		if dir.Name() == newDirName {
			return
		}
	}
	t.Errorf("Could not find created directory")
}

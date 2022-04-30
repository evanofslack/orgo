package main

import (
	"fmt"
	"os"
	"testing"
	"testing/fstest"

	"golang.org/x/exp/slices"
)

func TestCreateDir(t *testing.T) {

	const newDirName string = "test_directory"

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
	var found bool = false
	for _, dir := range dirs {
		if dir.Name() == newDirName {
			found = true
		}
	}
	if !found {
		t.Errorf("Could not find created directory")
	}

	// Clean up created directory
	t.Cleanup(func() {
		err = os.RemoveAll(newDirName)
		if err != nil {
			t.Errorf("Failed to clean up created dir")
		}
	})
}

func TestFindFiles(t *testing.T) {
	fsys := fstest.MapFS{
		"file_1.txt":  {},
		"file_2.png":  {},
		"file_3.pdf":  {},
		"file_4.jpeg": {},
	}

	match_ext := []string{".txt", ".png", ".pdf", ".jpeg"}
	matches := findFiles(fsys, match_ext)

	want_length := 4
	if len(matches) != want_length {
		err := fmt.Sprintf("Wanted %d matches, got %d", want_length, len(matches))
		t.Errorf(err)
	}

	want_matches := []string{"file_1.txt", "file_2.png", "file_3.pdf", "file_4.jpeg"}
	if !slices.Equal(matches, want_matches) {
		err := fmt.Sprintf("Wanted %s , got %s", want_matches, matches)
		t.Errorf(err)
	}
}

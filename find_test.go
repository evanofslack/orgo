package orgo

import (
	"fmt"
	"log"
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
	err = CreateDir(newDirName)
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

		// files we want to find
		"file_1.txt":  {},
		"file_2.png":  {},
		"file_3.pdf":  {},
		"file_4.jpeg": {},

		// files we do not want to find
		"file_5.csv": {},
		"file_6.cpp": {},
		"file_7.go":  {},
		"file_8.js":  {},

		// child dir files we do not want to find
		"dir_1/file_9.txt":       {},
		"dir_1/dir_2/file_9.png": {},
	}

	t.Run("find correct files", func(t *testing.T) {
		match_ext := []string{".txt", ".png", ".pdf", ".jpeg"}
		matches := FindFiles(fsys, match_ext)
		want_length := 4
		if len(matches) != want_length {
			errMsg := fmt.Sprintf("Wanted %d matches, got %d", want_length, len(matches))
			t.Errorf(errMsg)
		}
		want_matches := []string{"file_1.txt", "file_2.png", "file_3.pdf", "file_4.jpeg"}
		if !slices.Equal(matches, want_matches) {
			errMsg := fmt.Sprintf("Wanted %s , got %s", want_matches, matches)
			t.Errorf(errMsg)
		}
	})
	t.Run("do not find incorrect files", func(t *testing.T) {
		match_ext := []string{".nope", ".fake"}
		matches := FindFiles(fsys, match_ext)
		want_length := 0
		if len(matches) != want_length {
			errMsg := fmt.Sprintf("Wanted %d matches, got %d", want_length, len(matches))
			t.Errorf(errMsg)
		}
		want_matches := []string{}
		if !slices.Equal(matches, want_matches) {
			errMsg := fmt.Sprintf("Wanted %s , got %s", want_matches, matches)
			t.Errorf(errMsg)
		}
	})
}

func TestFindFilesRecursive(t *testing.T) {
	fsys := fstest.MapFS{

		// files we want to find
		"file_1.txt":             {},
		"file_2.png":             {},
		"file_3.pdf":             {},
		"file_4.jpeg":            {},
		"dir_1/file_5.txt":       {},
		"dir_1/dir_2/file_6.png": {},

		// files we do not want to find
		"file_7.csv":              {},
		"file_8.cpp":              {},
		"file_9.go":               {},
		"file_10.js":              {},
		"dir_1/file_11.csv":       {},
		"dir_1/dir_2/file_12.cpp": {},
	}

	t.Run("find correct files", func(t *testing.T) {
		match_ext := []string{".txt", ".png", ".pdf", ".jpeg"}
		matches := findFilesRecursive(fsys, match_ext)
		want_length := 6
		if len(matches) != want_length {
			errMsg := fmt.Sprintf("Wanted %d matches, got %d", want_length, len(matches))
			t.Errorf(errMsg)
		}
		want_matches := []string{"dir_1/dir_2/file_6.png", "dir_1/file_5.txt", "file_1.txt", "file_2.png", "file_3.pdf", "file_4.jpeg"}
		if !slices.Equal(matches, want_matches) {
			errMsg := fmt.Sprintf("Wanted %s , got %s", want_matches, matches)
			t.Errorf(errMsg)
		}
	})
	t.Run("do not find incorrect files", func(t *testing.T) {
		match_ext := []string{".nope", ".fake"}
		matches := FindFiles(fsys, match_ext)
		want_length := 0
		if len(matches) != want_length {
			errMsg := fmt.Sprintf("Wanted %d matches, got %d", want_length, len(matches))
			t.Errorf(errMsg)
		}
		want_matches := []string{}
		if !slices.Equal(matches, want_matches) {
			errMsg := fmt.Sprintf("Wanted %s , got %s", want_matches, matches)
			t.Errorf(errMsg)
		}
	})
}

func TestMoveFiles(t *testing.T) {
	const newDirName string = "test_directory"

	err := CreateDir(newDirName)
	if err != nil {
		t.Errorf("failed to make output dir")
	}

	// Create temp files
	inFiles := []string{"file_1.txt", "file_2.png", "file_3.pdf", "file_4.jpeg"}
	for _, f := range inFiles {
		_, err := os.Stat(f)
		if os.IsNotExist(err) {
			file, err := os.Create(f)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
		}
	}

	err = MoveFiles(inFiles, newDirName)
	if err != nil {
		fmt.Println(err)
		t.Errorf("failed to move files")
	}
	dirs, err := os.ReadDir(newDirName)
	if err != nil {
		t.Errorf("unable to read dir")
	}

	if len(dirs) != len(inFiles) {
		errMsg := fmt.Sprintf("wanted %d files, got %d", len(inFiles), len(dirs))
		t.Errorf(errMsg)
	}
	for _, dir := range dirs {
		if !slices.Contains(inFiles, dir.Name()) {
			t.Errorf("Unexpected file in output dir")
		}
	}

	t.Cleanup(func() {
		err := os.RemoveAll(newDirName)
		if err != nil {
			t.Errorf("Failed to clean up created dir")
		}
	})
}

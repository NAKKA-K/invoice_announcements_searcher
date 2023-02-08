package main

import "testing"

func TestGetFileNames(t *testing.T) {
	files, err := GetFileNames(DataDir)
	if err != nil {
		t.Error(err)
	}

	if len(files) != 5 {
		t.Errorf("logic exception: files length %d(%v)", len(files), files)
	}
}

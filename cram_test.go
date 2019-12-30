package main_test

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"testing"
)

func runCram(cram, test string) func(t *testing.T) {
	fn := func(t *testing.T) {
		buf := new(bytes.Buffer)
		cmd := exec.Command(cram, test)
		cmd.Stdout = buf
		cmd.Stderr = buf
		if err := cmd.Run(); err != nil {
			t.Log(buf.String())
			t.Fatalf("cram: %v", err)
		}
	}
	return fn
}

func TestCram(t *testing.T) {
	cram, err := exec.LookPath("cram")
	if err != nil {
		t.Fatalf("cannot find cram test runner: %v", err)
	}
	const dir = "testdata/cram"
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatalf("cannot list cram tests: %v", err)
	}
	for _, fi := range fis {
		name := fi.Name()
		ext := filepath.Ext(name)
		if ext != ".t" {
			continue
		}
		base := name[:len(name)-len(ext)]
		t.Run(base, runCram(cram, filepath.Join(dir, name)))
	}
}

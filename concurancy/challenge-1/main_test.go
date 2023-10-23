package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go updateMessage("Testing", &wg)
	wg.Wait()
	if msg != "Testing" {
		t.Error("Expected to find Testing but did not found")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	msg = "testing"
	os.Stdout = w
	printMessage()
	w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	if !strings.Contains(output, "testing") {
		t.Error("Expected to find testing but did not found")
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	main()
	w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Expected to find Hello, universe! but did not found")
	}
	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("Expected to find Hello, cosmos! but did not found")
	}
	if !strings.Contains(output, "Hello, world!") {
		t.Error("Expected to find Hello, world! but did not found")
	}
}

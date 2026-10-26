package main

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"simple": {input: "  hello  world  ", expected: []string{"hello", "world"}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			diff := cmp.Diff(tc.expected, got)
			if diff != "" {
				t.Fatalf("%s", diff)
			}
		})
	}
}

func TestCommandExit(t *testing.T) {

	cases := map[string]struct {
		expected error
		text     string
	}{
		"print_exit": {expected: nil, text: "Closing the Pokedex... Goodbye!\n"},
	}

	for name, tc := range cases {
		//spawn os seperate process
		cmd := exec.Command(os.Args[0], "-test.run=TestCommandExitSubprocess")
		cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")

		// intercept console output from sub-process
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out

		err := cmd.Run()

		// verify the exit code
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				t.Fatalf("Process exited with code %d, want code 0", exitErr.ExitCode())
			} else {
				t.Fatalf("failed to run subprocess: %v", err)
			}
		}

		gotText := out.String()
		if !strings.Contains(gotText, tc.text) {
			t.Errorf("got console output %q, want it to contain %q", gotText, tc.text)
		}

		var mockGot error = nil
		diff := cmp.Diff(tc.expected, mockGot, cmp.Exporter(func(reflect.Type) bool { return true }))
		if diff != "" {
			t.Fatal(diff)
		}

		t.Run(name, func(t *testing.T) {
			got := commandExit()
			diff := cmp.Diff(tc.expected, got)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestCommandExitSubprocess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	_ = commandExit()
}

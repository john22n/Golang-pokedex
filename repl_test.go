package main

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"io"
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

func TestCommandHelp(t *testing.T) {
	expectedStrings := []string{
		"Welcome to the Pokedex!",
		"Usage",
	}

	for _, cmd := range getCommands() {
		expectedStrings = append(expectedStrings, cmd.name, cmd.description)
	}

	cases := map[string]struct {
		expected       error
		wantSubstrings []string
	}{
		"print_help_menu": {
			expected:       nil,
			wantSubstrings: expectedStrings,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			got := commandHelp()

			w.Close()
			os.Stdout = oldStdout
			var buf bytes.Buffer
			_, _ = io.Copy(&buf, r)
			gotText := buf.String()

			// assert function return state matches
			diff := cmp.Diff(tc.expected, got, cmp.Exporter(func(reflect.Type) bool { return true }))
			if diff != "" {
				t.Fatal(diff)
			}

			//verify header elements and dynamic cmds appear in console
			for _, subStr := range tc.wantSubstrings {
				if !strings.Contains(gotText, subStr) {
					t.Errorf("got console output %q, missing expected: %q", gotText, subStr)
				}
			}
		})
	}
}

func TestGetCommands(t *testing.T) {
	cases := map[string]struct {
		wantKeys []string
	}{
		"validate_pokedex_commands": {
			wantKeys: []string{"help", "exit"},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := getCommands()

			//assert the map contains the exit num of expected commands
			if len(got) != len(tc.wantKeys) {
				t.Fatalf("got map size %d, want %d", len(got), len((tc.wantKeys)))
			}

			for _, key := range tc.wantKeys {
				cmd, exists := got[key]
				if !exists {
					t.Errorf("expected command key %q was not found", key)
					continue
				}
				//verify internal struct fields arnt blank
				if cmd.name == "" {
					t.Errorf("comamnd %q has an empty name field", key)
				}
				if cmd.description == "" {
					t.Errorf("command %q has an emtpy description", key)
				}
				if cmd.callback == nil {
					t.Errorf("command %q has a nil callback", key)
				}
			}

			wantHelpCmd := cliCommand{
				name:        "help",
				description: "Display a help message",
				callback:    commandHelp,
			}

			diff := cmp.Diff(wantHelpCmd, got["help"], cmp.Exporter(
				func(reflect.Type) bool { return true }),
				cmp.Comparer(
					func(x, y func() error) bool {
						return reflect.ValueOf(x).Pointer() == reflect.ValueOf(y).Pointer()
					}),
			)

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
		t.Run(name, func(t *testing.T) {
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

		})
	}
}

func TestCommandExitSubprocess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	_ = commandExit()
}
